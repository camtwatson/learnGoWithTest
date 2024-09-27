package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("cam")
	want := "Hello, cam"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
