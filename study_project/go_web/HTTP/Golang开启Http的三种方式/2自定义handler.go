package main

import (
	"net/http"
	"regexp"
)

type MyHandler struct {
}

func (my MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/abc" {
		w.Write([]byte("三问中国男足几多愁，恰是阳痿患者逛青楼,预设不能"))
		return
	}

	//http://localhost:8080/index1.html
	//http://localhost:8080/index2.html
	exp1 := regexp.MustCompile(`/index[1-9].html`)
	result := exp1.FindAllStringSubmatch(r.URL.Path, -1)
	if len(result) != 0 {
		w.Write([]byte("四问中国男足几多愁,恰是一群小孩上青楼,尚射不能"))
		return
	}
	w.Write([]byte("index"))
}

type MyHandler1 struct {
}

func (my MyHandler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/abc" {
		w.Write([]byte("xdl"))
		return
	}

	//http://localhost:8080/index1.html
	//http://localhost:8080/index2.html
	exp1 := regexp.MustCompile(`/index[1-9].html`)
	result := exp1.FindAllStringSubmatch(r.URL.Path, -1)
	if len(result) != 0 {
		w.Write([]byte("区块链"))
		return
	}
	w.Write([]byte("index1"))
}

func main() {
	http.Handle("/", MyHandler{})
	http.ListenAndServe(":8080", MyHandler1{})
}
