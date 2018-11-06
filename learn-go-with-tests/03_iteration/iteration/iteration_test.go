package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"
	if repeated != expected {
		t.Errorf("expected %s, got %s\n", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	s := Repeat("a", 10)
	fmt.Println(s)
	// Output:
	// aaaaaaaaaa
}
