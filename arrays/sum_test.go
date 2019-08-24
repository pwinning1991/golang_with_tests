package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func ExampleSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers))
	//Output 15
}
