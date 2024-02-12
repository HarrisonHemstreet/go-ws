package utils

import "net/http"

// HandlerFuncWithDB type defines the expected signature for your handlers.
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// registerRoute takes a URI, a handler, and a database connection, then registers the handler for the URI.
func RegisterRoute(uri string, handlerFunc HandlerFunc) {
	wrappedHandler := func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(w, r)
	}
	http.HandleFunc(uri, wrappedHandler)
}
