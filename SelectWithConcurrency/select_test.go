package selectwithconcurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T){
	
	t.Run("comparing the two servers and return the fasted url",func(t *testing.T) {

		slowServer := makeServer(20*time.Millisecond)
		fastServer := makeServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		fastURL := fastServer.URL
		slowURL := slowServer.URL
		want := fastURL
		got,_ := Racer(slowURL,fastURL,30*time.Millisecond)

		if want != got{
			t.Errorf("want %q got %q",want, got)
		}

	})


	t.Run("returns an error if the server repons time is greater than  specific time",func(t *testing.T) {
		slowServer := makeServer(4*time.Millisecond)	
		defer slowServer.Close()

		_,err := Racer(slowServer.URL,slowServer.URL,3*time.Millisecond)

		if err ==nil{
			t.Error("expected an error but didn't get one")
		}


	})

}


func makeServer(duration time.Duration) *httptest.Server{

	server :=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)

	}))
	return server
}
