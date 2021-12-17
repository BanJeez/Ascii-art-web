package main

import (
	"bufio"
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

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method != "POST" {
		http.Error(w, "Please use HTTP POST method", http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request in Form", http.StatusBadRequest)
		return
	}
	if _, err := os.Stat("artwork.txt"); err == nil {
		os.Remove("artwork.txt")
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

	tplPath := filepath.Join("templates", "ascii-art.html")
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
