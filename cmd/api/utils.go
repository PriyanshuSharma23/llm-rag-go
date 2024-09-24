package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type envelope map[string]any

func (app *application) parseJSONBody(
	r *http.Request,
	body interface{},
) error {
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

func (app *application) writeJSONResponse(
	w http.ResponseWriter,
	statusCode int,
	body interface{},
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	buf := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(w)
	return err
}
