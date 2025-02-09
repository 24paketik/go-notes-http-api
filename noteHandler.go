package main

import (
	"fmt"
	"net/http"
)

func noteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "метод Get")
	case "POST":

		fmt.Fprintln(w, "POST")
	case "PUT":

		fmt.Fprintln(w, "PUT")
	case "DELETE":

		fmt.Fprintln(w, "DELETE")
	default:

		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
