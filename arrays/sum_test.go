package main

import "testing"

func TestSum(t *testing.T) {

	// 数组有一个固定的容量，可以在声明它的时候来定义它。我们有两种方式来定义：
	// case1: [N]type{value1, value2, ..., valueN}
	// case2: [...]type{value1, value2, ..., valueN}
	//numbers := [5]int{1, 2, 3, 4, 5}
	//
	//got := Sum(numbers)
	//want := 15
	//
	//if got != want {
	//	// 有时候在错误信息之中打印函数的输入也很有用。
	//	t.Errorf("got %d, want %d, %v", got, want, numbers)
	//}

	// 还有一件事，就是使用数组实在是太不方便了，指定数组大小非常麻烦！
	// 于是切片（slice）应运而生，任意的大小！
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		// replace arrays with slices
		got := Sum(numbers)
		want := 15

		checkSum(t, got, want, numbers)
	})
}

func checkSum(t *testing.T, got int, want int, numbers []int) {
	t.Helper() // 从而使其指向真正调用的地方，而不是辅助函数内部

	if got != want {
		t.Errorf("got %d, want %d, %v", got, want, numbers)
	}
}
