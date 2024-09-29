package main

import "testing"

func TestSum(t *testing.T) {
	nums := [5]int{1, 2, 3, 4, 5}

	got := sum(nums)
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
