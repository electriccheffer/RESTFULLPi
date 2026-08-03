package handler

import "net/http"
import "os"
import "errors"
import "encoding/json"
import "fmt"
import "io/fs"
import "restfulpi/internal/models"
import "restfulpi/internal/file_operations"


type FrontendHandler struct{

	fileServer: http.Handler,
	fileSystem: fs.FS		

} 

func NewFrontendHandler(fileSystem fs.FS) *FrontendHandler {
	
	fileServer := http.FileServer(http.FS(fileSystem))
	frontHandler := &FrontendHandler{fileServer:fileServer,fileSystemfileSystem}
	return frontHandler 
	
}

func (fh *FrontendHandler) ServeHTTP(response http.ResponseWriter,request *http.Request) {
	
	path := request.URL.Path
	if path != "/"{
		
		if file,err := fh.fileSystem.Open(path[1:]) ; err != nil{
	
			request.URL.Path = "/"
		} else{
				
			stat, err := file.Stat()
			if err != nil || stat.IsDir() {
				request.URL.Path = "/"
			}
			_ = file.Close()  
		}
	}	
	
	fh.fileServer.ServeHTTP(response,request)
}

type StatusHandler struct{}

func NewStatusHandler() *StatusHandler{
	
	sh := &StatusHandler{}
	return sh
	
}

func (sh *StatusHandler) ServeHTTP(response http.ResponseWriter,request *http.Request){
	
	status := models.Status{Device:"RESTFULPi",Status:"Online"}
	encodedStatus,err := json.Marshal(status)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		http.Error(response,err.Error(),http.StatusInternalServerError)
		return
	}		
	response.WriteHeader(http.StatusOK)
	response.Write(encodedStatus)
		
	return

}

type GetLogsHandler struct{
	path string
}

func NewGetLogsHandler(path string) *GetLogsHandler{
	lh := &GetLogsHandler{path}
	return lh 
}

func (glh *GetLogsHandler) ServeHTTP(response http.ResponseWriter,request *http.Request){
	
	directoryReader := file_operations.NewDirectoryRead(glh.path)
	logFiles, err := directoryReader.GetFiles()		
	if err != nil {
		response.Header().Set("Content-Type","application/json")
		response.Header().Set("X-Content-Type-Options","nosniff")
		response.WriteHeader(http.StatusNotFound)
		return	
	}	
	
	var logs []models.Logs	
	for _, entry := range logFiles{		
		if entry.IsDir(){
			continue
		}
		log := models.Logs{entry.Name()}
		logs = append(logs,log)
	}

	jsonLogs,err := json.Marshal(logs)
	if err != nil{
		fmt.Print("error in ServeHTTP marshaling")
	
	}
	response.WriteHeader(http.StatusOK)
	response.Write(jsonLogs)		
	return
}
