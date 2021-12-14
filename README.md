# Ascii-Art-Web

ASCII-art-web is a project that creates and maintains a server, in which it's possible for the user to use a web GUI (graphical user interface). to make text string into ASCII art.

## running the project

1. From your teminal going to the ascii-art-web dir.

2. then type `go run .` this will start the svera on port 8080

3. in a web browser go to localhost 8080

## Main Code Overview 

By useing the Package http provides us with a HTTP client and server implementations.

We uses the Get, Head, Post, and PostForm to make HTTP requests:

`func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method != "POST" {
		http.Error(w, "Please use HTTP POST method", http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request in Form", http.StatusBadRequest)
		return
	}`



## Authors 

David, Brksygmr, Banji

