package handler

import "net/http"
import "os"
import "errors"
import "encoding/json"
import "fmt"

import "restfulpi/internal/models"
import "restfulpi/internal/file_operations"


type FrontendHandler struct{
		
	filePath string	

} 

func NewFrontendHandler(path string) *FrontendHandler {
	frontHandler := &FrontendHandler{filePath:path}
	return frontHandler 
	
}

func (fh *FrontendHandler) ServeHTTP(response http.ResponseWriter,request *http.Request) {
	
	exists,err := os.Stat(fh.filePath)
	if err != nil{
		
		if errors.Is(err,os.ErrNotExist){
			response.WriteHeader(http.StatusNotFound)
			http.Error(response,err.Error(),http.StatusNotFound)
			return
		}
	
	}
	if exists.IsDir() {
		response.WriteHeader(http.StatusNotAcceptable)
		errorMessage := "file was directory"
		http.Error(response,errorMessage,http.StatusNotAcceptable)
		return	
	}
	application,err := os.ReadFile(fh.filePath)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)	
		http.Error(response,err.Error(),http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(application)
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
		//Handle error 
	}	
	
	var logs []models.Logs	
	for _, entry := range logFiles{		
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
