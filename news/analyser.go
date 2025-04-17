package news

import (
	"context"
	"encoding/json"
	"fmt"

	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/tolulopejoel/newsApp/internal/database"
	"google.golang.org/api/option"
)

func Analyse(article database.Article) {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")
	model.ResponseMIMEType = "application/json"
	model.ResponseSchema = &genai.Schema{
		Type: genai.TypeArray,
		Items: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"title": {
					Type:        genai.TypeString,
					Description: "A short, emotionally charged sentence that hints at the negative side of the news to hook readers.",
				},
				"summary": {
					Type:        genai.TypeString,
					Description: "A summary highlighting the hopeful or positive side of the news.",
				},
			},
		},
	}

	prompt := `
Read this article and rewrite it with a hopeful tone. Focus on the possible good, the efforts being made, or what could be learned. End with a positive takeaway.
Article:
"""` + article.Content + `"""
`

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getResponseFields(resp))
}

func getResponseFields(resp *genai.GenerateContentResponse) (title string, summary string) {
	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		return
	}

	// Convert the text to bytes and attempt to parse it as JSON
	content := resp.Candidates[0].Content
	byteData := []byte(content.Parts[0].(genai.Text))

	var result []map[string]string
	if err := json.Unmarshal(byteData, &result); err != nil {
		log.Println("Error parsing JSON:", err)
	}

	if len(result) > 0 {
		title = result[0]["title"]
		summary = result[0]["summary"]
	}

	return
}
