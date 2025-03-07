package hello

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("hello to `world` when name is not provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("hello to a name when it is provided", func(t *testing.T) {
		got := Hello("Rodik", "")
		want := "Hello, Rodik"
		assert.Equal(t, want, got)
	})

	t.Run("Hola to a name when language is Spanish", func(t *testing.T) {
		got := Hello("Mirko", "Spanish")
		expected := "Hola, Mirko"
		assert.Equal(t, expected, got)
	})

	t.Run("Bonjour to a name when language is French", func(t *testing.T) {
		got := Hello("Rodik", "French")
		expected := "Bonjour, Rodik"
		assert.Equal(t, expected, got)
	})
}

// For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy,
// so you can call helper functions from a test, or a benchmark
func assertCorrectMessage(t testing.TB, got, want string) {
	//this t.Helper() indicates it is a helper method. That means if tests fail, error line will be in
	//Caller line not inside this method
	t.Helper()
	if got != want {
		t.Errorf("Got %q but expected %q", got, want)
	}
}
