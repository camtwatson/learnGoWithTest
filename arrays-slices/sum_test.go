package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := sumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of some slices", func(t *testing.T) {
		got := sumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("Sum empty slices", func(t *testing.T) {
		got := sumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)

	})

}
