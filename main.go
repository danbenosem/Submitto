package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/static/style.css", styleHandler)

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

func styleHandler(write http.ResponseWriter, read *http.Request) {

	write.Header().Set("Content-Type", "text/css")
	http.ServeFile(write, read, "static/style.css")
}
