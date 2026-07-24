package router

import "net/http"

func NewRouter(handler http.Handler) http.Handler {

	mux := http.NewServeMux()
	mux.Handle("GET /", handler)
	return mux	
}
