package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		/*name := r.PostFormValue("name")
		fmt.Println("name = ", name)

		age := r.PostFormValue("age")
		fmt.Println("age = ", age)


		sex := r.PostFormValue("sex")
		fmt.Println("sex = ", sex)

		hobby := r.PostFormValue("hobby")
		fmt.Println("hobby = ", hobby)

		address := r.PostFormValue("address")
		fmt.Println("address = ", address)*/

		r.ParseForm()
		name := r.PostForm["name"]
		fmt.Println("name = ", name)

		age := r.PostForm["age"]
		fmt.Println("age = ", age)

		sex := r.PostForm["sex"]
		fmt.Println("sex = ", sex)

		hobby := r.PostForm["hobby"]
		fmt.Println("hobby = ", hobby)

		address := r.PostForm["address"]
		fmt.Println("address = ", address)

	})
	http.ListenAndServe(":8080", nil)
}
