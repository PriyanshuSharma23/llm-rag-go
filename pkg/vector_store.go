package vector_store

type VectorStore interface {
	AddDocument()
	SimilaritySearch()
}
