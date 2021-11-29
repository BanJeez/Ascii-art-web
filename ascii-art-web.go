package main

import (
	"fmt"
	"log"
	"net/http"
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

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// checkes if user is useing the right path
	if r.URL.Path != "/test" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	// check for verifying the type of the request
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello Test!")
}

func main() {
	fileServer1 := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer1)
	http.HandleFunc("/test", helloHandler) // the web path
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
