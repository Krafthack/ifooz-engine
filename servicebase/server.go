package servicebase

import (
    "net/http"
)

type Server struct {
    mux  *http.ServeMux
}

func NewServer()  (http.Handler, *http.ServeMux) {
    mux := http.NewServeMux()
    return Server{mux}, mux
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s.mux.ServeHTTP(w, r)
}
