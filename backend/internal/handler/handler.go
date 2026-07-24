package handler
import "net/http"
import "os"
import "errors"

type FrontendHandler struct{
		
	filePath string	

} 

func NewFrontendHandler(path string) *FrontendHandler {
	frontHandler := &FrontendHandler{filePath:path}
	return frontHandler 
	
}

func (fh *FrontendHandler) ServeHTTP(response http.ResponseWriter,request*http.Request) {
	
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
	response.Write(application);
}
