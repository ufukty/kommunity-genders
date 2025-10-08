// This script is needed because ChatGPT constantly refused to
// provide result for the full list of names in the input file.
package main

import (
	"context"
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
	Start, End int
}

type Question struct {
	MemberName string `json:"member-name" jsonschema:"description=Claimed to be a human name"`
}

type Answer struct {
	Gender string `json:"gender" jsonschema:"enum=male,female,unisex,other"`
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
			p := fmt.Sprintf(`Create an Answer with the following requirements: Member Name: %s`, q.MemberName)
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

	defer func() {
		fmt.Println("total    :", included+excluded)
		fmt.Println("included :", included)
		fmt.Println("excluded :", excluded)
	}()

	memberNames := strings.Split(string(f), "\n")
	if args.End == -1 {
		args.End = len(memberNames)
	}
	memberNames = memberNames[args.Start:args.End]
	pct := -1

	for i, memberName := range memberNames {
		q := &Question{
			MemberName: memberName,
		}
		a, err := flow.Run(context.Background(), q)
		if err != nil {
			log.Fatalf("flow.Run: %v", err)
		}

		switch a.Gender {
		case "male":
			fmt.Fprintln(o.Male, memberName)
			included += 1
		case "female":
			fmt.Fprintln(o.Female, memberName)
			included += 1
		case "unisex", "other":
			excluded += 1
		default:
			fmt.Printf("WARNING: unexpected answer from LLM: %q\n", a.Gender)
			excluded += 1
		}

		if pct2 := percentage(i, len(memberNames)); pct2 > pct {
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
