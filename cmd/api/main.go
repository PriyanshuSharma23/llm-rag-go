package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PriyanshuSharma23/llm-rag-go/pkg/vectorstore"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type application struct {
	vectorStore vectorstore.VectorStore
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func main() {
	godotenv.Load()
	app := &application{}

	app.infoLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	app.errorLogger = log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)

	client, collection, err := vectorstore.NewChromaClient()
	if err != nil {
		app.errorLogger.Fatal(err)
	}

	app.vectorStore = vectorstore.NewChromaDB(client, collection)
	app.infoLogger.Println("Connected to ChromaDB")

	app.infoLogger.Println("Starting server...")

	router := chi.NewRouter()

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Post("/v1/upload", app.upload)
	router.Post("/v1/query", app.query)

	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		port = "8080"
	}

	app.infoLogger.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		app.errorLogger.Fatal(err)
	}
}
