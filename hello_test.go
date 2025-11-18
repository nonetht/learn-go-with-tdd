package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("in Spanish", func(t *testing.T) {
		// 更多的需求，现在我们要支持第二个参数，就是指定问候语的语言。
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})
}

// t *testing.TB 为什么会出现问题呢？是一个*testing.T, *testing.B 都满足的接口
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // 告诉测试套件，这是一个辅助方法
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
