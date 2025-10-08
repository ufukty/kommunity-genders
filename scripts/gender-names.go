// This script is needed because ChatGPT constantly refused to
// provide result for the full list of names in the input file.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

type Args struct {
	Start, End, Batch int
}

type Question struct {
	MemberNames []string `json:"member-names" jsonschema:"description=Claimed to be a human name"`
}

type Answer struct {
	Genders map[string]string `json:"genders" jsonschema:"description=A dictionary from member names to either of male, female, unisex or other"`
}

type OutputFiles struct {
	Male, Female io.WriteCloser
}

func percentage(current, total int) int {
	return int(float64(current) / float64(total))
}

func Main() error {
	args := Args{}
	flag.IntVar(&args.Start, "start", 0, "start index")
	flag.IntVar(&args.End, "end", -1, "start index")
	flag.IntVar(&args.Batch, "batch", 10, "batch")
	flag.Parse()

	api := &googlegenai.GoogleAI{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	}

	g := genkit.Init(
		context.Background(),
		genkit.WithPlugins(api),
		genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
	)

	flow := genkit.DefineFlow(g, "AnswerGeneratorFlow",
		func(ctx context.Context, q *Question) (*Answer, error) {
			j, err := json.Marshal(q.MemberNames)
			if err != nil {
				return nil, fmt.Errorf("encoding question into json: %w", err)
			}
			p := fmt.Sprintf(`Create an Answer with the following requirements: Member Names: %s`, j)
			a, _, err := genkit.GenerateData[Answer](ctx, g, ai.WithPrompt(p))
			if err != nil {
				return nil, fmt.Errorf("genkit.GenerateData: %w", err)
			}
			return a, nil
		},
	)

	f, err := os.ReadFile("labels/uniq-names.txt")
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	o := OutputFiles{
		Male:   nil,
		Female: nil,
	}

	o.Male, err = os.Create("labels/male.txt")
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer o.Male.Close()

	o.Female, err = os.Create("labels/female.txt")
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer o.Female.Close()

	included, excluded := 0, 0
	pct := -1
	batch := 0

	defer func() {
		fmt.Println("total :", included+excluded)
		fmt.Println("incl. :", included)
		fmt.Println("excl. :", excluded)
		fmt.Println("pct.  :", pct)
		fmt.Println("batch :", batch)
	}()

	memberNames := strings.Split(string(f), "\n")
	if args.End == -1 {
		args.End = len(memberNames)
	}
	memberNames = memberNames[args.Start:args.End]

	for b := 0; b*args.Batch < len(memberNames); b++ {
		var (
			from = min(len(memberNames), args.Batch*(b))
			to   = min(len(memberNames), args.Batch*(b+1))
		)
		q := &Question{
			MemberNames: memberNames[from:to],
		}
		a, err := flow.Run(context.Background(), q)
		if err != nil {
			log.Fatalf("flow.Run: %v", err)
		}

		for name, gender := range a.Genders {
			switch gender {
			case "male":
				fmt.Fprintln(o.Male, name)
				included += 1
			case "female":
				fmt.Fprintln(o.Female, name)
				included += 1
			case "unisex", "other":
				excluded += 1
			default:
				fmt.Printf("WARNING: unexpected answer from LLM: %q\n", gender)
				excluded += 1
			}
		}

		if pct2 := percentage(included+excluded, len(memberNames)); pct2 > pct {
			pct = pct2
			fmt.Printf("progress: %%%d\n", pct)
		}
	}

	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
