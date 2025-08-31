// ===============================================================
// GO-MIPS Instruction Set Editor and Simulator
// @dev-xero on GitHub
// 2025
// ===============================================================
package main

import (
	"fmt"
	"net/http"
)

// ===============================================================
// HTTP Server for the web editor
// ===============================================================
func startServer() {
	// Starts up a HTTP server to listen for requests to the specified
	// port. It also serves static files required by the web editor
	port := 8080
	fmt.Printf("Server running at http://localhost:%d\n", port)

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/index.html")
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// ===============================================================
// Entry point
// ===============================================================
func main() {
	startServer()
}
