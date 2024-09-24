package main

import (
	"fmt"
	"net/http"

	"github.com/PriyanshuSharma23/llm-rag-go/pkg/llm"
	"github.com/PriyanshuSharma23/llm-rag-go/pkg/vectorstore"
)

type uploadRequest struct {
	Documents vectorstore.Documents `json:"documents"`
}

func (app *application) upload(w http.ResponseWriter, r *http.Request) {
	var req uploadRequest
	if err := app.parseJSONBody(r, &req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err := app.vectorStore.AddDocuments(req.Documents)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.infoLogger.Println("Documents uploaded successfully")

	err = app.writeJSONResponse(w, http.StatusOK, map[string]string{"message": "Documents uploaded successfully"})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type queryRequest struct {
	QueryText string `json:"query_text"`
}

func (app *application) query(w http.ResponseWriter, r *http.Request) {
	var req queryRequest
	if err := app.parseJSONBody(r, &req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	results, err := app.vectorStore.SimilaritySearch(req.QueryText)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.infoLogger.Println("Similarity search completed successfully")

	referencesText := ""
	for _, result := range results {
		referencesText += result.Document + "\n"
	}

	llmQuery := fmt.Sprintf("%s\nRefernces:\n%s", req.QueryText, referencesText)

	resp, err := llm.GenerateCompletion(llmQuery)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSONResponse(w, http.StatusOK, envelope{"response": resp})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
