package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("5 numbers", func(t *testing.T) {
		nums := [5]int{1, 2, 3, 4, 5}

		got := sum(nums)
		want := 15

		if got != want {
			t.Errorf("got %d want %d  given, %v", got, want, nums)
		}
	})
	t.Run("Any amount of numbers", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := sumSlice(nums)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})

}
