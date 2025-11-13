package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

// t *testing.TB 为什么会出现问题呢？
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // 告诉测试套件，这是一个辅助方法
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
