package http

import "net/http"

func ListenAndServe(addr string, h http.Handler) error {
	return http.ListenAndServe(addr, h)
}
