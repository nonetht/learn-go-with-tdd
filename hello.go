package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// 1.
//func main() {
//	fmt.Println("Hello, world")
//}

// 2. 将代码和外部世界（side-effect）分离，才是也给好的测试方法。
// 3. 如果Hello的长度太长的话，你可以将函数进行拆解
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	// 将if 重构为 switch-case语句
	return greetingPrefix(language) + name
}

// 这里的重构部分，我们使用了 named return value，就是设置返回值为 prefix，从而创建
// prefix 变量在函数之中，int: 0, string: "", 只需要调用 return 即可，无需 return prefix
// Go之中，公共函数以大写字母开头，私有函数以小写字母开头
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
