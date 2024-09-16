package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello with a name", func(t *testing.T) {
		got := Hello("Jack", "")
		want := "Hello, Jack"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is passed in", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %q got %q.", want, got)
	}
}
