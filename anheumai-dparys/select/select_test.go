package select_pkg

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("Resolve slowest url", func(t *testing.T) {
		slowServer := buildDelayServer(20 * time.Millisecond)
		fastServer := buildDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Returns error after 10s", func(t *testing.T) {
		slowServer := buildDelayServer(12 * time.Second)
		fastServer := buildDelayServer(11 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, err := Racer(slowURL, fastURL)

		if err == nil {
			t.Error("expected error, bro")
		}
	})

	t.Run("Returns before timeout of 25ms", func(t *testing.T) {

		server := buildDelayServer(25 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("Unexpected timeout")
		}

	})

}

func buildDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}
