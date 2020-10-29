package sumarray

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 3 numbers", func(t *testing.T) {

		numbers := [3]int{5, 5, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("Expected %d, got %d", expected, got)
		}

	})

	t.Run("collection of any size", func(t *testing.T) {

		numbers := []int{2, 2, 2, 2}

		got := SumSlice(numbers)
		want := 8

		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}

	})

	t.Run("sum all collections", func(t *testing.T) {

		got := SumAll([]int{2, 2}, []int{5, 5})
		want := []int{4, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %d, got %d", want, got)
		}

	})

}
