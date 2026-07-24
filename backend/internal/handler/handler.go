package handler
import "net/http"
import "os"

type FrontendHandler struct{
		
	filePath string	

} 

func NewFrontendHandler(path string) *FrontendHandler {
	frontHandler := &FrontendHandler{filePath:path}
	return frontHandler 
	
}

func (fh *FrontendHandler) ServeHTTP(response http.ResponseWriter,request*http.Request) {
	
	response.WriteHeader(http.StatusOK)
	application,err := os.ReadFile(fh.filePath)
	if err != nil {
		
		http.Error(response,err.Error(),http.StatusInternalServerError)
		return
	}
	response.Write(application);
}
