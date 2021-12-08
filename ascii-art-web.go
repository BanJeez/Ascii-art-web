package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

// func formHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}
// 	fmt.Fprintf(w, "POST request successful")
// 	name := r.FormValue("name")
// 	banner := r.FormValue("banner")

// 	fmt.Fprintf(w, "Name = %s\n", name)
// 	fmt.Fprintf(w, "Banner = %s\n", banner)
// }

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

	ascii := r.FormValue("ascii")
	banner := r.FormValue("banner")

	fmt.Fprintf(w, "Ascii = %s", ascii)
	fmt.Fprintf(w, "Banner = %s", banner)

	AsciiArt(ascii, banner)

	type Asciiart struct {
		Art1 string
		Art2 string
		Art3 string
		Art4 string
		Art5 string
		Art6 string
		Art7 string
		Art8 string
	}

	art := []string{}
	art = append(art, "")

	fArt, err := os.OpenFile("artwork.txt", os.O_RDONLY, 0400)
	if err != nil {
		http.Error(w, "Error when Outputting", http.StatusInternalServerError)
		return
	}
	defer fArt.Close()

	lineScanner := bufio.NewScanner(fArt)
	lineScanner.Split(bufio.ScanLines)

	for lineScanner.Scan() {
		art = append(art, lineScanner.Text())
	}
	fmt.Println(art[1])
	fmt.Println(art[2])
	fmt.Println(art[3])
	fmt.Println(art[4])
	fmt.Println(art[5])
	fmt.Println(art[6])
	fmt.Println(art[7])
	fmt.Println(art[8])

	artwork := Asciiart{
		Art1: art[1],
		Art2: art[2],
		Art3: art[3],
		Art4: art[4],
		Art5: art[5],
		Art6: art[6],
		Art7: art[7],
		Art8: art[8],
	}
	// var input []string
	// tempmap := make(map[int][]string)
	// letters := strings.Split(ascii, "")

	// for i := 0; i < len(ascii); i++ {
	// 	input = (AsciiArt(letters[i], banner))
	// 	for j, line := range input {
	// 		tempmap[j] = append(tempmap[j], line)
	// 	}
	// }

	// for k := 0; k < 8; k++ {
	// 	for m := 0; m < len(ascii); m++ {
	// 		fmt.Print(tempmap[k][m])
	// 		// fmt.Fprintf(w, tempmap[k][m])
	// 	}
	// 	fmt.Println("")
	// 	// fmt.Fprintln(w, "")
	// }

	tplPath := filepath.Join("static", "ascii-art.gohtml")
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

// going to integrate with formHandler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// check for verifying the type of the request
	if r.Method != "GET" {
		http.Error(w, "Please use HTTP GET method", http.StatusBadRequest)
		return
	}

	tplPath := filepath.Join("static", "index.html")
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

	http.HandleFunc("/*", pathHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", http.HandlerFunc(pathHandler)); err != nil {
		log.Fatal(err)
	}
}
