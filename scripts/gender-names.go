// This script is needed because ChatGPT constantly refused to
// provide result for the full list of names in the input file.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

var systemInstruction = `You are a careful name annotator. 
For the given NAME, infer the most likely gender category
from this set: ["male","female","unisex","unknown"].

Rules:
- Prefer "unisex" if the name is commonly used by multiple genders in any major locale.
- Use "unknown" for initials, handles, company names, or if you are not confident.
- Consider cultural/linguistic context broadly (e.g., Turkish, Arabic, Persian, Slavic, Western European).
- Return STRICT JSON, no extra text.

Schema:
{"label":"<male|female|unisex|unknown>"}`

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

var prompt = `
You are a careful name annotator. For the given NAME, infer the most likely gender
category from: ["male","female","unisex","unknown"].

Rules:
- Prefer "unisex" if the name is commonly used by multiple genders in any major locale.
- Use "unknown" for initials, handles, organization names, or if confidence is low.
- Consider cultural/linguistic contexts (e.g., Turkish, Arabic, Persian, Slavic, Western European).
- Return STRICT JSON and nothing else.

Schema:
{"label":"<male|female|unisex|unknown>"}

EXAMPLE
NAMES: ["mehmet","ayşe","alex","kim","özge","deniz","abc holdings"]
OUTPUT: {
"Mehmet":"male"
"Ayşe":"female",
"Alex":"unisex",
"Kim":"unisex",
"Özge":"female",
"Deniz":"unisex",
"ABC Holdings":"unknown"
}
END EXAMPLE

NAMES: {{.}}
`

func Main() error {
	args := Args{}
	flag.IntVar(&args.Start, "start", 0, "start index")
	flag.IntVar(&args.End, "end", -1, "start index")
	flag.IntVar(&args.Batch, "batch", 10, "batch")
	flag.Parse()

	t, err := template.New("").Parse(prompt)
	if err != nil {
		return fmt.Errorf("parsing prompt template: %w", err)
	}

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
			names := bytes.NewBuffer([]byte{})
			if err := json.NewEncoder(names).Encode(q.MemberNames); err != nil {
				return nil, fmt.Errorf("encoding question into json: %w", err)
			}
			prompt := bytes.NewBufferString("")
			if err = t.Execute(prompt, names); err != nil {
				return nil, fmt.Errorf("templating the prompt: %w", err)
			}
			s := prompt.String()
			fmt.Println(s)
			a, _, err := genkit.GenerateData[Answer](ctx, g, ai.WithPrompt(s))
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
			case "unisex", "unknown":
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
