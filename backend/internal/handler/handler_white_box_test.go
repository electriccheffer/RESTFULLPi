package handler

import "io"
import "testing"
import "path/filepath"
import "net/http"
import "net/http/httptest"
import "encoding/json"
import "regexp"
import "restfulpi/internal/models"
import "restfulpi/internal/file_operations"

func TestSessionHandlerIntegrationSuccessCase(t *testing.T){

	// set up test directories
	readDirectory := t.TempDir()
	writeDirectory := t.TempDir()
	readFilePath := filepath.Join(readDirectory,"device")
	fileOpener := file_operations.NewFileOpener()
	_,err := fileOpener.Create(readFilePath)
	if err != nil{

		t.Errorf("unexpected error creating read path: %s",err.Error())
	}	

	// create session manager 
	sessionManager := NewSessionManager(writeDirectory,readFilePath,fileOpener)
	
	// create handler 
	handler := NewSessionStartHandler(sessionManager)
	
	//create a the requests and response 
	request := httptest.NewRequest(http.MethodPost,"/logs/sessions",nil)
	response := httptest.NewRecorder()
			
	//call the handler
	handler.ServeHTTP(response,request)
	
	got, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("unexpected error reading response: %s",err.Error())
	}
	
	if response.Code != http.StatusCreated {
		t.Errorf("incorrect status code expected:%d got:%d",http.StatusCreated,
								   response.Code)
	}
	
	var session models.Session 
	
	err = json.Unmarshal(got,&session)
	if err != nil {
		t.Errorf("unexpected error unmarshaling response: %s",err.Error())	
	}
	
	fileName := session.FileName
	fileNamePattern := `^\d{2}_\d{2}_\d{2}_\d{2}_\d{2}_\d{2}\.gpx$`	
	regularExpression := regexp.MustCompile(fileNamePattern)
	if !regularExpression.MatchString(fileName){
		t.Errorf("File name did not match the pattern got:%s",fileName)
	}
	
}
