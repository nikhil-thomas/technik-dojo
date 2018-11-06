package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3
	if got != want {
		t.Errorf("expected %d got %d\n", want, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output:
	// 3
}
