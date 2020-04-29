package messages

import (
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expected := "Hello, Gopher!\n"
	if got != expected {
		t.Errorf("Did not get expected result, wanted: %q, got: %q\n", expected, got)
	}
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expected := "Goodbye, Gopher\n"
	if got != expected {
		t.Errorf("Did not get expected result, wanted: %q, got: %q\n", expected, got)
	}
}

func TestFailureTypes(t *testing.T) {
	t.Error("Error signals a failed test, but does not stop the rest from executing")
	t.Fatal("Fatal will stop the test")
	t.Error("This will never be reached")
}

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct{
		input string
		expected string
	}{
		{input: "Gopher", expected: "Hello, Gopher!\n"},
		{input: "", expected: "Hello, !\n"},
	}
	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expected {
			t.Errorf("Did not get expected result, wanted: %q, got: %q\n", s.expected, got)
		}
	}
}


