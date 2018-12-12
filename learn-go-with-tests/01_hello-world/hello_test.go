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

func TestHelloWithLanguage(t *testing.T) {
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello3("Spanish", "Chris")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello3("French", "Chris")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'\n", got, want)
	}
}
