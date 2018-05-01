package main

import (
	"net/http"
	"fmt"
	"html/template"
	"crypto/md5"
	"io"
	"strconv"
	"time"
)

func httpHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "hello world,go web!!")

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求方式是", r.Method);
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		t, _ := template.ParseFiles("view/login.html");
		h:=md5.New();
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t.Execute(w, token);
	} else {
		r.ParseForm() //默认是不支持解析form字段的
		fmt.Fprintln(w, "Method is ", r.Method);
		fmt.Fprint(w, "username:", r.Form["username"], "password:", r.Form["password"])
		for key, value := range r.Form {
			fmt.Println("key", key)
			fmt.Println("value", value)
		}
	}

	//fmt.Fprint(w, )
}

func main() {
	http.HandleFunc("/goweb.go", httpHandler);
	http.HandleFunc("/login.go", login);
	http.ListenAndServe(":9000", nil);
}
