package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("试问中国男足几多愁，恰是一群太监上青楼，无人能射!"))
	})
	http.HandleFunc("/bye", sayBye)
	http.ListenAndServe(":8080", nil)
}

func sayBye(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("再问中国男足几多愁,恰似一群几女守青楼,总是被射!"))
}
