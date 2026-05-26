package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from snipppet"))
}
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("the id is %v", id)
	w.Write([]byte(msg))
}
func creatsnippetGET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("creat a snipppet"))
}
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("add a new post"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", showSnippet)
	mux.HandleFunc("GET /snippet/create", creatsnippetGET)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	log.Print("starting new server at :8080 ")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
