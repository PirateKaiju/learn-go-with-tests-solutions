package countdown

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("test countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySlepper := &SpySleeper{}

		Countdown(buffer, spySlepper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		if spySlepper.Calls != 4 {
			t.Errorf("not enough calls, expected 4, got %d", spySlepper.Calls)
		}
	})

	t.Run("sleep before each print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v, got %v", want, spySleepPrinter.Calls)
		}

	})
}
