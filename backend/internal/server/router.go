package router 

import "net/http"
import "restfulpi/internal/handler"
import "path/filepath"
import "embed"
import "io/fs"

//go:embed dist/browser/*
var applicationPath embed.FS

func NewRouter(logsPath string) http.Handler {
	
	
	applicationFS, err := fs.Sub(applicationPath,"dist/browser")
	if err != nil{
		panic(err)
	}
		
	mux := http.NewServeMux()

	frontendHandler := handler.NewFrontendHandler(applicationFS)
	statusHandler := handler.NewStatusHandler()
	getLogsHandler := handler.NewGetLogsHandler(logsPath)
	
	mux.Handle("GET /",frontendHandler) 
	mux.Handle("GET /status",statusHandler)
	mux.Handle("GET /logs",getLogsHandler)
	return mux	
}
