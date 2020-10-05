package main

import (
	"fmt"
	"net/http"
)

// MyHandler ...
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home!")
}

// HelloHandler ...
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// WorldHandler ...
type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	handler := MyHandler{}
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/", &handler)
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServeTLS("cert.pem", "key.pem")
}
