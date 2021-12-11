package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ArtPiece struct {
	ArtLines []string
}

func processTemplate(tplPath string, w http.ResponseWriter) {
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

func assciArtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method != "POST" {
		http.Error(w, "Please use HTTP POST method", http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request in Form", http.StatusBadRequest)
		return
	}

	ascii := r.PostForm.Get("ascii")
	asciiSlice := strings.Split(ascii, "\r\n")
	banner := r.FormValue("banner")

	for i := 0; i < len(asciiSlice); i++ {
		AsciiArt(asciiSlice[i], banner)
	}

	fArt, err := os.OpenFile("artwork.txt", os.O_RDONLY, 0400)
	if err != nil {
		http.Error(w, "Error when Outputting", http.StatusInternalServerError)
		return
	}
	defer fArt.Close()

	lineScanner := bufio.NewScanner(fArt)
	lineScanner.Split(bufio.ScanLines)

	allArt := []string{}
	allArt = append(allArt, "")

	for lineScanner.Scan() {
		allArt = append(allArt, lineScanner.Text())
	}

	artwork := ArtPiece{
		ArtLines: allArt,
	}

	tplPath := filepath.Join("templates", "ascii-art.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Parse Error: %v", err)
		http.Error(w, "Error when Parsing", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, artwork)
	if err != nil {
		log.Printf("Execute Error: %v", err)
		http.Error(w, "Error when Executing", http.StatusInternalServerError)
		return
	}
}

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
	http.HandleFunc("/*", pathHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", http.HandlerFunc(pathHandler)); err != nil {
		log.Fatal(err)
	}
}
