package vectorstore

type VectorStore interface {
	AddDocuments(documents Documents)
	SimilaritySearch()
}

type Document struct {
	Content  string
	Metadata map[string]interface{}
}

type Documents []Document
