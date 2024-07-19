package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAsyncRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowerServer := makeTestServer(20 * time.Millisecond)
		fastServer := makeTestServer(0 * time.Millisecond)

		// Close the servers when the test finishes
		defer slowerServer.Close()
		defer fastServer.Close()

		slowUrl := slowerServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := AsyncRacer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		// prefer def closer to instantiation
		// slowerServer.Close()
		// fastServer.Close()
	})

	t.Run("returns an error if a server doesn't respond within 3s", func(t *testing.T) {
		serverA := makeTestServer(11 * time.Millisecond)
		serverB := makeTestServer(12 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableAsyncRacer(serverA.URL, serverB.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeTestServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
