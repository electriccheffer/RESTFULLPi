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
	
	//check that the file exists 
	_,err := os.Stat(fh.filePath)
	if err != nil{
		
		if errors.Is(err,os.ErrNotExist){
			response.WriteHeader(http.StatusNotFound)
			http.Error(response,err.Error(),http.StatusNotFound)
			return
		}
	
	}
	//check for read errors 
	application,err := os.ReadFile(fh.filePath)
	if err != nil {
		
		http.Error(response,err.Error(),http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(application);
}
