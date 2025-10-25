package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func GenText(str string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error:", err)
	}
	env := os.Getenv("GEMINI_API_KEY")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(env))
	if err != nil {
		log.Fatal(err)
	}
	
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash-lite")
	resp, err := model.GenerateContent(ctx, genai.Text(str))
	if err != nil {
		log.Fatal(err)
	}

	getResp := printResponse(resp)
	return getResp
}



func printResponse(resp *genai.GenerateContentResponse) string {
	var validResp string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				validResp = validResp + fmt.Sprintf("%s\n", part)
			}
		}
	}
	return validResp
}

