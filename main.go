package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name   string
	gender string
	Age    int32
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 渲染模板
	u1 := User{
		Name:   "小王子",
		gender: "男",
		Age:    18,
	}

	err = t.Execute(w, u1)
	if err != nil {
		fmt.Println(err)
		return
	}
}