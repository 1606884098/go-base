package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	Id   int
	Name string
}

func main() {
	http.HandleFunc("/", BodyHandler)
	http.ListenAndServe(":8080", nil)
}

func BodyHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read body fail:", err)
		w.WriteHeader(500)
		return
	}

	var u User
	if err = json.Unmarshal(body, &u); err != nil {
		fmt.Println("json unmarshal fail: ", err)
		w.WriteHeader(500)
		return
	}
	w.Write([]byte("user.id:" + strconv.Itoa(u.Id)))
	w.Write([]byte("user.name:" + u.Name))
}
