package sloppy

import (
	"fmt"
	"net/http"
)

// Sloppy is a wrapper for http.Handler and act as an interceptor to
// provide hints / suggestions, when a user has a typo in the request
type Sloppy struct {
	http.Handler

	routes    *pathTree
	onSuggest OnSuggestFunc
}

// OnSuggestFunc needs to be provided from the user to customize the
// return value
type OnSuggestFunc func(string) []byte

// New creates a new Sloppy instance that can be used as interceptor
func New(handler http.Handler, routes []string, onSuggest OnSuggestFunc) Sloppy {
	tree := newPathTree()
	for _, route := range routes {
		tree.addPath(route)
	}

	return Sloppy{
		Handler:   handler,
		routes:    tree,
		onSuggest: onSuggest,
	}
}

// ServeHTTP intercepts requests before passing it on to the actual
// handler function
func (s Sloppy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	interceptor := &interceptResponseWriter{
		ResponseWriter: w,
		pathTree:       s.routes,
		uri:            req.RequestURI,
		onSuggest:      s.onSuggest,
	}
	s.Handler.ServeHTTP(interceptor, req)
}

// Print outputs sloppy's internal path tree to stdout
func (s Sloppy) Print() {
	fmt.Println(s.routes.string())
}

type interceptResponseWriter struct {
	http.ResponseWriter
	*pathTree

	uri       string
	status    int
	onSuggest OnSuggestFunc
}

func (w *interceptResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *interceptResponseWriter) Write(buf []byte) (int, error) {
	// If status code is not 404, proceed as usual
	if w.status != 404 {
		return w.ResponseWriter.Write(buf)
	}

	suggested, ok := Suggest(w.pathTree, w.uri)
	// No suggested url
	if !ok {
		return w.ResponseWriter.Write(buf)
	}

	resp := w.onSuggest(suggested)

	return w.ResponseWriter.Write(resp)
}
