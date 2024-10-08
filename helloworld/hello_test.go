// This is the beginning of me learning Go with Learn Go With Test
package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("cam", " ")
		want := "Hello, cam"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when empty string", func(t *testing.T) {
		got := hello("", " ")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Hello in spanish", func(t *testing.T) {
		got := hello("cam", "spanish")
		want := "Hola, cam"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Hello in French", func(t *testing.T) {
		got := hello("cam", "french")
		want := "Bonjour, cam"
		assertCorrectMessage(t, got, want)
	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
