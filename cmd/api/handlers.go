package main

import (
	"net/http"

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

func (app *application) query(w http.ResponseWriter, r *http.Request) {
}
