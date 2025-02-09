package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/note", noteHandler)
	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}
