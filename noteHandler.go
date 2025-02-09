package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Note struct {
	Text string `json:"text"`
}

var (
	notes     = make(map[int]Note)
	idCounter = 1
	mu        sync.Mutex
)

func noteHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	switch r.Method {
	case "GET":
		//response := map[string]string{"message": "ноте"}
		//json.NewEncoder(w).Encode(response)

		var noteList []Note
		for _, note := range notes {
			noteList = append(noteList, note)
		}
		json.NewEncoder(w).Encode(noteList)

	case "POST":
		var note Note
		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		notes[idCounter] = note
		idCounter++
		json.NewEncoder(w).Encode(note)

	case "PUT":
		var request struct {
			ID   int    `json:"id"`
			Text string `json:"text"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if note, exists := notes[request.ID]; exists {
			// Обновляем текст заметки
			note.Text = request.Text
			notes[request.ID] = note
			json.NewEncoder(w).Encode(note)
		} else {
			http.Error(w, "Заметка не найдена", http.StatusNotFound)
		}
	case "DELETE":
		var noteID int
		if err := json.NewDecoder(r.Body).Decode(&noteID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if _, exists := notes[noteID]; exists {
			delete(notes, noteID)
			response := map[string]string{"message": "Deleted note"}
			json.NewEncoder(w).Encode(response)
		} else {
			http.Error(w, "Note ID not found", http.StatusBadRequest)
		}
	default:

		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
