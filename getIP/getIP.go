package main

import (
	"fmt"
	"net/http"
	"strings"
)

func getIP(r *http.Request) string {
	// Check X-Forwarded-For header first
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return strings.Split(forwarded, ",")[0]
	}
	
	// Otherwise use RemoteAddr
	return strings.Split(r.RemoteAddr, ":")[0]
}

func handler(w http.ResponseWriter, r *http.Request) {
	clientIP := getIP(r)
	fmt.Fprintf(w, "Hello World from IP: %s", clientIP)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
