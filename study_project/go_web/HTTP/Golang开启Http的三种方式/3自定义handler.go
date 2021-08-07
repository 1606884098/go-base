package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myhandler{})
	mux.HandleFunc("/bye", sayBye)
	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	server.ListenAndServe()
}

type myhandler struct {
}

func (my *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("xdl"))
}

func sayBye(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("再问中国男足几多愁,恰似一群几女守青楼,总是被射!"))
}
