package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// 来自 bytes 包的 Buffer 类型实现了 Writer 接口，
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
