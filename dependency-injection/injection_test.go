package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	t.Run("should write to the buffer", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Gopher")

		got := buffer.String()
		want := "Hello, Gopher"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
