package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, %q</h1>", html.EscapeString(r.URL.Path))
}

func main() {
	// read config from environment
	port := os.Getenv("PORT")

	// start server
	http.HandleFunc("/", handler)
	log.Println("listening=" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
