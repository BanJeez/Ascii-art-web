package main

import "net/http"

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/ascii-art":
		asciiArtHandler(w, r)
	default:
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}