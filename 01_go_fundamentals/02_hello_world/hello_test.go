package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jack")
	want := "Hello, Jack"

	if got != want {
		t.Errorf("Expected %q got %q.", want, got)
	}
}
