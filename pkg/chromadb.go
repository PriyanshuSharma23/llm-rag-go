package vector_store

import chroma "github.com/amikos-tech/chroma-go"

type chromaDB struct {
}

func NewChromaClient(url string) (*chroma.Client, error) {
	client, err := chroma.NewClient(url, chroma.WithDatabase("llm-rag-go"))
	return client, err
}

func NewChromaDB() VectorStore {
	return &chromaDB{}
}

func (vs *chromaDB) AddDocument() {
}

func (vs *chromaDB) SimilaritySearch() {
}
