package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// 1.
//func main() {
//	fmt.Println("Hello, world")
//}

// 2. 将代码和外部世界（side-effect）分离，才是也给好的测试方法。
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	} else if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
