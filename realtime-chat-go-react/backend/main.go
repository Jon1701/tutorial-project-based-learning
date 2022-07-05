package main

import (
	"fmt"
	"net/http"
)

func HandleSimpleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple Server")
}

func setupRoutes() {
	http.HandleFunc("/", HandleSimpleServer)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}