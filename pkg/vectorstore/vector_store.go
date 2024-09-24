package vectorstore

type VectorStore interface {
	AddDocuments(documents Documents) error
	SimilaritySearch(query string) ([]SearchResult, error)
}

type Document struct {
	Content  string                 `json:"content"`
	Metadata map[string]interface{} `json:"metadata"`
}

type Documents []Document

type SearchResult struct {
	Document string                 `json:"document"`
	Distance float32                `json:"distance"`
	Metadata map[string]interface{} `json:"metadata"`
	ID       string                 `json:"id"`
}
