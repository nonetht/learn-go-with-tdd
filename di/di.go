package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
1. test code: 就是使用注入依赖更加便于测试。
2.
3. ...
*/

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name) // Fprintf 类型于 fmt.Printf, 但它接收一个Writer将字符串发送到目标位置。
}

func myGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(myGreeterHandler)))
}
