package main

import (
	"context"
	"fmt"
	"os"

	chroma "github.com/amikos-tech/chroma-go"
	"github.com/amikos-tech/chroma-go/openai"
	"github.com/amikos-tech/chroma-go/types"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	applicatonName := os.Getenv("APPLICATION_NAME")
	if applicatonName == "" {
		applicatonName = "llm-rag-go"
	}

	url := os.Getenv("CHROMA_URL")
	if url == "" {
		url = "http://localhost:8000"
	}

	client, err := chroma.NewClient(url)
	if err != nil {
		panic(err)
	}

	// check if database exists
	_, err = client.GetDatabase(context.Background(), applicatonName, nil)
	if err == nil {
		fmt.Println("Database already exists")
		goto collection
	}

	_, err = client.CreateDatabase(context.Background(), applicatonName, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database created successfully")

collection:
	// check for collection
	_, err = client.GetCollection(context.Background(), "documents", nil)
	if err == nil {
		fmt.Println("Collection already exists")
		return
	}

	openaiEf, err := openai.NewOpenAIEmbeddingFunction(os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		panic(err)
	}

	_, err = client.CreateCollection(context.Background(), "documents", nil, true, openaiEf, types.L2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Collection created successfully")
}
