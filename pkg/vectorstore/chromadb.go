package vectorstore

import (
	"context"
	"fmt"
	"os"

	chroma "github.com/amikos-tech/chroma-go"
)

type chromaDB struct {
	collection *chroma.Collection
	client     *chroma.Client
}

func NewChromaClient() (*chroma.Client, *chroma.Collection, error) {
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
		return nil, nil, fmt.Errorf("error creating chroma client: %s", err)
	}

	newCollection, err := client.GetCollection(context.Background(), "documents", nil)

	if err != nil {
		fmt.Println(err.Error())
		return nil, nil, fmt.Errorf("error getting chroma collection: %s", err)
	}

	return client, newCollection, nil
}

func NewChromaDB(client *chroma.Client, collection *chroma.Collection) VectorStore {
	return &chromaDB{
		collection: collection,
		client:     client,
	}
}

func (vs *chromaDB) AddDocuments(documents Documents) {
}

func (vs *chromaDB) SimilaritySearch() {
}
