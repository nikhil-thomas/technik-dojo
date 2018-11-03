package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'\n", got, want)
	}
}

func TestHello2(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'\n", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello2("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World' for empty input", func(t *testing.T) {
		got := Hello2("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
