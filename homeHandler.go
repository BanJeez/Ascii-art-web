package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method != "GET" {
		http.Error(w, "Please use HTTP GET method", http.StatusBadRequest)
		return
	}

	tplPath := filepath.Join("templates", "index.html")
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