package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mindworker/sloppy"
)

func main() {
	// Create Endpoints with Gorilla's mux as usual
	r := mux.NewRouter()
	r.HandleFunc("/", handler("Hello there! Try out /v3/tracking"))
	r.HandleFunc("/v4/trackings", handler("Great! How about this /v4/courier/sall"))
	r.HandleFunc("/v4/couriers", handler("Great! How about this /v4/courier/sall"))
	r.HandleFunc("/v4/couriers/all", handler("How about dynamic urls? /v4/notifications/123/456/ad"))
	r.HandleFunc("/v4/notifications/{slug}/{tracking_number}/add", handler("Hope you liked it!"))

	// Use FromGorilla() to wrap gorilla mux router with sloppy
	// and provide a function to handle suggestions
	http.Handle("/", sloppy.FromGorilla(r, func(suggestion string) (resp []byte) {
		return []byte("Did you mean: " + suggestion + "?")
	}))

	http.ListenAndServe(":8080", nil)
}

// handler simply respond with the string message given
func handler(text string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(text))
	}
}
