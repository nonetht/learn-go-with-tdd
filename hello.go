package main

import "fmt"

// 1.
//func main() {
//	fmt.Println("Hello, world")
//}

// 2. 将代码和外部世界（side-effect）分离，才是也给好的测试方法。
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
