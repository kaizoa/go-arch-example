package http

import "net/http"

type router struct {
	mux *http.ServeMux
}

func (s *router) Handle(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

func (s *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func NewRouter() *router {
	return &router{mux: http.NewServeMux()}
}
