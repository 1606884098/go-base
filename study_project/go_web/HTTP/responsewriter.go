package main

import (
	"bytes"
	"fmt"
	"net/http"
)

type MyResponseWriter struct {
	http.ResponseWriter
	bodyBytes *bytes.Buffer
}

func (mrw MyResponseWriter) Write(body []byte) (int, error) {
	mrw.bodyBytes.Write(body) //记录下返回的内容
	return mrw.ResponseWriter.Write(body)
}

func (mrw MyResponseWriter) Body() []byte {
	return mrw.bodyBytes.Bytes()
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	m := MyResponseWriter{
		ResponseWriter: w,
		bodyBytes:      bytes.NewBuffer(nil),
	}
	m.Header().Add("Content-Type", "text/html")
	//如果需要设置状态码，需要在内容返回之前进行设置，内容返回之后设置无效，默认是200
	m.WriteHeader(404)
	m.Write([]byte("<h1>Hello World</h1>"))

	fmt.Println("body: ", string(m.Body()))
}

func main() {
	http.HandleFunc("/", HttpHandler)
	http.ListenAndServe(":8080", nil)
}
