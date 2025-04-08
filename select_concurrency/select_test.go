package select_concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func createDelayedMockServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("should return the faster website", func(t *testing.T) {

		slowServer := createDelayedMockServer(10 * time.Millisecond)
		fastServer := createDelayedMockServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		fastUrl := fastServer.URL
		slowUrl := slowServer.URL

		want := fastUrl
		got, _ := Race(fastUrl, slowUrl)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should throw timeout error after the specified timeout", func(t *testing.T) {
		server := createDelayedMockServer(7 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRace(server.URL, server.URL, time.Millisecond*2)

		if err == nil {
			t.Errorf("Expecting timeout error, getting none")
		}
	})
}
