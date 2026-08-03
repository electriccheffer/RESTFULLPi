package router 

import "net/http"
import "restfulpi/internal/handler"
import "path/filepath"
import "embed"
import "io/fs"

func NewRouter(appPath string,logsPath string) http.Handler {
	
	//go:embed dist/browser/*
	var applicationPath embed.FS
	applicationFS, err := fs.Sub(applicationPath,"dist/browser")
	if err != nil{

	}
	
	mux := http.NewServeMux()
	statusHandler := handler.NewStatusHandler()
	getLogsHandler := handler.NewGetLogsHandler(logsPath)
 	fileServer := http.FileServer(http.FS(applicationFS))
	mux.Handle("GET /",fileServer) 
	mux.Handle("GET /status",statusHandler)
	mux.Handle("GET /logs",getLogsHandler)
	return mux	
}
