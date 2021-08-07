package main

import (
	"fmt"
	"net/http"
)

/*
1.Get将表单中的数据按照key=value的形式，添加到action所指向的URL后面，并且两者使用?连接，
而各个变量之间使用&连接，Post是将表单中的数据放在请求体中进行提交
2.Get是不安全的，因为在传输过程中，数据被放在请求的URL中，敏感数据容易泄露，Post的所有操作对用户来说是不可见的。
3.Get传输的数据量小，这主要是因为受URL长度限制，而POST可以传递大量数据。

*/
//http://localhost:8080/GetAndPost?name=xdl&age=18&name=区块链
//map[string][]string
//http://localhost:8080/GetAndPost
func main() {
	http.HandleFunc("/GetAndPost", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.Form["name"]
		pwd := r.Form["pwd"]
		fmt.Println("name = ", name)
		fmt.Println("pwd = ", pwd)
		w.Write([]byte("Hello Go Http Server!"))
	})
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("start http server fail, err = ", err)
	}
}
