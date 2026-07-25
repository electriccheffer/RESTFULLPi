package router

import "net/http"

func NewRouter(frontendHandler http.Handler,statusHandler http.Handler) http.Handler {

	mux := http.NewServeMux()
	mux.Handle("GET /", frontendHandler)
	mux.Handle("GET /status",statusHandler);
	return mux	
}
