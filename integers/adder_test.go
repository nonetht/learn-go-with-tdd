package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		// 这里相较于helloworld而言，是 %d 而不是 %q，采用整数而不是字符串
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, expected)
	}
}

// ExampleAdd 示例函数，添加示例函数，是您的代码更加容易访问
func ExampleAdd() {
	sums := Add(1, 5)
	fmt.Println(sums)
}
