package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	banner := r.FormValue("banner")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Banner = %s\n", banner)
}

func assciArtHandler(w http.ResponseWriter, r *http.Request) {

}

// going to integrate with formHandler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// check for verifying the type of the request
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	tplPath := filepath.Join("static", "form.html")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Parse Error: %v", err)
		http.Error(w, "Error when Parsing", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Execute Error: %v", err)
		http.Error(w, "Error when Executing", http.StatusInternalServerError)
		return
	}
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/ascii-art":
		assciArtHandler(w, r)
	default:
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}

func main() {
	// fileServer1 := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer1)
	http.HandleFunc("/", pathHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", http.HandlerFunc(pathHandler)); err != nil {
		log.Fatal(err)
	}
}
