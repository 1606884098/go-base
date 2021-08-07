package main

import (
	"fmt"
	"net/http"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Method:", r.Method)
	fmt.Println("URL:", r.URL)
	fmt.Println("URL.Path:", r.URL.Path)
	fmt.Println("RemoveAddress:", r.RemoteAddr)
	fmt.Println("UserAgent:", r.UserAgent())
	fmt.Println("Header.Accept", r.Header.Get("Accept"))
}

//http://localhost:8080/aa/bb?name=xdl
func main() {
	http.HandleFunc("/", HttpHandler)
	http.ListenAndServe(":8080", nil)
}
