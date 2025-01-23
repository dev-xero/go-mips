package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	// Writing this by hand like a caveman
	fmt.Println("Server running at http://localhost:8000")

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/index.html")
	})

	http.ListenAndServe(":8000", nil)
}

func main() {
	startServer()
}
