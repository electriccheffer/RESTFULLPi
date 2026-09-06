package router 

import "net/http"
import "restfulpi/internal/handler"
import "embed"
import "io/fs"

//go:embed dist/browser/*
var applicationPath embed.FS

func NewRouter(logsPath string,sessionMangager handler.SessionManagerService) http.Handler {
	
	
	applicationFS, err := fs.Sub(applicationPath,"dist/browser")
	if err != nil{
		panic(err)
	}
		
	mux := http.NewServeMux()

	frontendHandler := handler.NewFrontendHandler(applicationFS)
	statusHandler := handler.NewStatusHandler()
	getLogsHandler := handler.NewGetLogsHandler(logsPath)
	sessionStartHandler := handler.NewSessionStartHandler(sessionMangager)

	mux.Handle("GET /",frontendHandler) 
	mux.Handle("GET /status",statusHandler)
	mux.Handle("GET /logs",getLogsHandler)
	mux.Handle("POST /logs/sessions",sessionStartHandler)
	return mux	
}
