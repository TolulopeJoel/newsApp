package news

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Analyse(content string) {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-pro-latest")
	model.ResponseMIMEType = "application/json"
	model.ResponseSchema = &genai.Schema{
		Type: genai.TypeArray,
		Items: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"Title": {
					Type:        genai.TypeString,
					Description: "A short, emotionally charged sentence that hints at the negative side of the news to hook readers.",
				},
				"positiveSide": {
					Type:        genai.TypeString,
					Description: "A summary highlighting the hopeful or positive side of the news.",
				},
				"turnAroundScore": {
					Type:        genai.TypeInteger,
					Description: "Score from 1 to 10 representing how strong the turnaround from negative to positive is.",
				},
			},
		},
	}

	prompt := `
You are a positivity filter for news.
Article:
"""` + content + `"""
`

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	fmt.Println(resp.UsageMetadata.TotalTokenCount)
	fmt.Println(resp.PromptFeedback)
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
