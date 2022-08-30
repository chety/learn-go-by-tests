package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Rodik")
	want := "Hello Rodik"
	if got != want {
		t.Errorf("Got %q but expected %q", got, want)
		t.Fail()
	}
}
