package vectorstore

type VectorStore interface {
	AddDocuments(documents Documents) error
	SimilaritySearch()
}

type Document struct {
	Content  string                 `json:"content"`
	Metadata map[string]interface{} `json:"metadata"`
}

type Documents []Document
