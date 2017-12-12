package middleware

import "net/http"

type Router interface {
	Handle(pattern string, handler http.Handler)
	ServeHTTP(http.ResponseWriter, *http.Request)
}
