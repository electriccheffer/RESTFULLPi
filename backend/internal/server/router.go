package router 

import "net/http"
import "restfulpi/internal/handler"

func NewRouter(filePath string) http.Handler {

	mux := http.NewServeMux()
	frontendHandler := handler.NewFrontendHandler(filePath)
	statusHandler := handler.NewStatusHandler()
	getLogsHandler := handler.NewGetLogsHandler()
 	 
	mux.Handle("GET /", frontendHandler)
	mux.Handle("GET /status",statusHandler)
	mux.Handle("GET /logs",getLogsHandler)
	return mux	
}
