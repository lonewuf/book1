package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("test")
	fmt.Println("test 2")
	w.Write([]byte("Hello world"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test show")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		fmt.Println(id)
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test post")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)	
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("Create snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fmt.Println("Listening on port 4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
