package main

import (
	"fmt"
	"net/http"
)

/*
curl -id "first_name=sausheong&last_name=chang" 127.0.0.1:8080/body
*/

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
