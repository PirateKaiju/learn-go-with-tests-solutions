package synccounter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	/*t.Run("increment 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)
	})*/

	t.Run("concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
