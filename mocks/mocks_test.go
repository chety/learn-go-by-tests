package mocks

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	t.Run("Should write 3,2,1, Go! the terminal", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpyCountdownOperations{}

		CountDown(buffer, sleeper)
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		calls := len(sleeper.Calls)
		if calls != 3 {
			t.Errorf("sleeper.Calls is %d, want %d", calls, 3)
		}
	})

	t.Run("Should write 3,2,1, Go! in the right order and respecting time.Sleeper", func(t *testing.T) {
		sleeper := &SpyCountdownOperations{}
		CountDown(sleeper, sleeper)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, sleeper.Calls) {
			t.Errorf("wanted calls %v, got %v", want, sleeper.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{Duration: sleepTime, Sleeper: spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("sleeper.Sleeper() got %v, want %v", spyTime.durationSlept, sleepTime)
	}

}
