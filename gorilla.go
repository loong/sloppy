package sloppy

import (
	"regexp"

	"github.com/gorilla/mux"
)

var gorillaExp = regexp.MustCompile("{.+?}")

// FromGorilla takes a gorilla.Mux to create Sloppy interceptor
func FromGorilla(handler *mux.Router, onSuggest OnSuggestFunc) Sloppy {
	routes := []string{}
	handler.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		r := gorillaExp.ReplaceAllString(t, "{}")
		routes = append(routes, r)
		return nil
	})

	return New(handler, routes, onSuggest)
}
