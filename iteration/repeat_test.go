package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("test repeat for 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("test repeat for 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t *testing.T, repeated string, expected string) {
	if repeated != expected {
		t.Errorf("Repeated was incorrect, got: %q, want: %q.", repeated, expected)
	}
}

// benchmark 基准测试
func BenchmarkRepeat(b *testing.B) { // testing.B 提供了对 loop函数的访问
	// ... setup ...
	for b.Loop() {
		Repeat("a", 100)
	}
}
