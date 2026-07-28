package router 

import "net/http"
import "restfulpi/internal/handler"

func NewRouter(appPath string,logsPath string) http.Handler {

	mux := http.NewServeMux()
	frontendHandler := handler.NewFrontendHandler(appPath)
	statusHandler := handler.NewStatusHandler()
	getLogsHandler := handler.NewGetLogsHandler(logsPath)
 	 
	mux.Handle("GET /", frontendHandler)
	mux.Handle("GET /status",statusHandler)
	mux.Handle("GET /logs",getLogsHandler)
	return mux	
}
