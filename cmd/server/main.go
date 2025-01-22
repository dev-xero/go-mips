package main

import (
	"fmt"
	"net/http"
	"os"
)

func startServer() {
	fmt.Println("Server running at http://localhost:8000")

	// Writing this by hand like a caveman
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		// Read HTML file
		htmlFilePath := "./public/index.html"
		htmlContent, err := os.ReadFile(htmlFilePath)

		if err != nil {
			http.Error(w, "Unable to load HTML file.", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(htmlContent)
	})

	http.ListenAndServe(":8000", nil)
}

func main() {
	startServer()
}
