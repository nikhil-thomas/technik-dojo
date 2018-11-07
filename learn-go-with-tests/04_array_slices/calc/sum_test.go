package calc

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	got := Sum(nums)
	want := 10
	if want != got {
		t.Errorf("expected %d, got %d\n", want, got)
	}

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if reflect.DeepEqual(want, got) {
		t.Errorf("expected %d, got %d\n", want, got)
	}
}

func TestSumTails(t *testing.T) {
	t.Run("calculate sum of tails", func(t *testing.T) {
		got := SumTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		if reflect.DeepEqual(want, got) {
			t.Errorf("expected %d, got %d\n", want, got)
		}
	})

	t.Run("safely return empty slices", func(t *testing.T) {
		got := SumTails([]int{}, []int{0, 1, 2, 3})
		want := []int{0, 6}
		if reflect.DeepEqual(want, got) {
			t.Errorf("expected %d, got %d\n", want, got)
		}
	})
}
