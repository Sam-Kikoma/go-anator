package main

import "testing"

func TestGreet_English(t *testing.T) {
	lang := language("en")
	want := "Hello World"

	got := greet(lang)

	if got != want {
		//test has failed
		t.Errorf("expected: %q, Got: %q", got, want)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour le Monde"

	got := greet(lang)

	if got != want {
		//test has failed
		t.Errorf("expected: %q, Got: %q", got, want)
	}
}

func TestGreet_Akkadian(t *testing.T) {
	//Akkadian is not implemented in the greet function, so we expect an empty string
	lang := language("akk")
	want := ""

	got := greet(lang)

	if got != want {
		//test has failed
		t.Errorf("expected: %q, Got: %q", got, want)
	}
}

//Stopping point was 2.3.1
