package main

import (
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World For Go Server!<h1>"))
}

func main() {
	http.Handle("/", &helloHandler{})
	http.ListenAndServe(":8080", nil)
}
