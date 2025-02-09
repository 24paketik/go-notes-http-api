package main

import (
	"encoding/json"
	"net/http"
)

type Note struct {
	Text string `json:"text"`
}

func noteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		response := map[string]string{"message": "ноте"}
		json.NewEncoder(w).Encode(response)
	case "POST":
		var note Note

		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(note)

	case "PUT":
		var note Note

		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(note)
	case "DELETE":
		response := map[string]string{"message": "ноте делете"}
		json.NewEncoder(w).Encode(response)
	default:

		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
