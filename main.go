package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", basicHandler)
	mux.HandleFunc("/test", testHandler)
	server := &http.Server{
		Addr: ":3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server")
	}
	
}

func basicHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func testHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Test!"))
}