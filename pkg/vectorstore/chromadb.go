package vectorstore

import (
	"context"
	"fmt"
	"os"

	chroma "github.com/amikos-tech/chroma-go"
	"github.com/amikos-tech/chroma-go/openai"
	"github.com/amikos-tech/chroma-go/types"
)

type chromaDB struct {
	collection *chroma.Collection
	client     *chroma.Client
}

func openAIEmbeddingFunction(apiKey string) (types.EmbeddingFunction, error) {
	return openai.NewOpenAIEmbeddingFunction(apiKey)
}

func newIdGenerator() types.IDGenerator {
	return types.NewULIDGenerator()
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

func (vs *chromaDB) AddDocuments(documents Documents) error {
	embeddingFunction, err := openAIEmbeddingFunction(os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		return fmt.Errorf("error creating embedding function: %s", err)
	}

	rs, err := types.NewRecordSet(
		types.WithEmbeddingFunction(embeddingFunction),
		types.WithIDGenerator(newIdGenerator()),
	)

	if err != nil {
		return fmt.Errorf("error creating record set: %s", err)
	}

	for _, document := range documents {
		rs.WithRecord(
			types.WithDocument(document.Content),
			types.WithMetadatas(document.Metadata),
		)
	}

	if _, err := rs.BuildAndValidate(context.Background()); err != nil {
		return fmt.Errorf("error building and validating record set: %s", err)
	}

	_, err = vs.collection.AddRecords(context.Background(), rs)
	if err != nil {
		return fmt.Errorf("error adding records to collection: %s", err)
	}

	return nil
}

func (vs *chromaDB) SimilaritySearch() {
}
