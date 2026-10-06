package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe("localhost:8080", nil)
}

func homeHandler(write http.ResponseWriter, read *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")

	if err != nil {

		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}

	temp.Execute(write, nil)
}
