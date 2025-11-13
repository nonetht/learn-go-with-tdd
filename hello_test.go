package main

import "testing"

func TestHello(t *testing.T) {
	// 使用Go 这种静态语言的时候，编译器非常重要！
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
