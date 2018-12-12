package webchecker

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	url1 := "http://www.facebook.com"
	url2 := "http://www.quii.co.uk"

	want := url1
	got := Racer(url1, url2)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
func TestRacerWithHTTPtest(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL

	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRacerPing(t *testing.T) {
	t.Run("returns url of faster server", func(t *testing.T) {
		fastServer := makeDelayedServer(0)
		slowServer := makeDelayedServer(10 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		want := fastServer.URL
		got, _ := RacerPing(fastServer.URL, slowServer.URL)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("returns error if server doesn't respond withing 5s", func(t *testing.T) {
		serverA := makeDelayedServer(1 * time.Second)
		serverB := makeDelayedServer(2 * time.Second)
		defer serverA.Close()
		defer serverB.Close()

		// _, err := RacerPing(serverA.URL, serverB.URL)
		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 5*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}
