package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, nil)
}
