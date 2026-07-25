package handler

import "testing"
import "path/filepath"
import "net/http"
import "net/http/httptest"
import "os"
import "io"
import "bytes"
import "encoding/json"

import "restfulpi/internal/models"
import "restfulpi/internal/server"

func TestAngularAppDelivery(t *testing.T){
	
	buildPath := filepath.Join("..","..","testData","browser","index.html")
	handler := NewFrontendHandler(buildPath)
	statusHandler := NewStatusHandler() 
	
	router := router.NewRouter(handler,statusHandler)
	
	expected,err := os.ReadFile(buildPath)
	if err != nil {
		t.Errorf("error reading file:" + err.Error())
	}
	expectedSlice := expected[1:20] 
	request := httptest.NewRequest(http.MethodGet,"/",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	
	if response.Code != http.StatusOK {
				
		t.Errorf("status: %d want %d",response.Code,http.StatusOK)
	}

	got, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("error reading body " + err.Error())
	} 					
	gotSlice := got[1:20]
	if !bytes.Equal(expectedSlice,gotSlice) {

		t.Errorf("the files were not the same")
	}	
}

func TestAngularAppDeliveryFileNotFound(t *testing.T){

	buildPath := filepath.Join("..","NotReal.php")
	handler := NewFrontendHandler(buildPath)
	statusHandler := NewStatusHandler()
	router := router.NewRouter(handler,statusHandler)
	
	request := httptest.NewRequest(http.MethodGet,"/",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	if response.Code != http.StatusNotFound{
			
		t.Errorf("Expected: %d, Got: %d",http.StatusNotFound,response.Code)
	}

}

func TestAngularAppDeliveryFileIsDirectory(t *testing.T){

	buildPath := filepath.Join("..","handler")
	handler := NewFrontendHandler(buildPath)	
	statusHandler := NewStatusHandler()
	router := router.NewRouter(handler,statusHandler)
	
	request := httptest.NewRequest(http.MethodGet,"/",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	if response.Code != http.StatusNotAcceptable{
		t.Errorf("Expected: %d, Got: %d",http.StatusNotAcceptable,response.Code)
	}
}

func TestStatusHandler(t *testing.T){

	buildPath := filepath.Join("..","handler")
	frontendHandler := NewFrontendHandler(buildPath)			
	statusHandler := NewStatusHandler() 
	router := router.NewRouter(frontendHandler,statusHandler)
	
	request := httptest.NewRequest(http.MethodGet,"/status",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	
	expectedStatus := models.Status{Device:"RESTFULPi",Status:"Online"}
	expectedJson,err := json.Marshal(expectedStatus)
	if err != nil{
	
		t.Errorf("Error with encoding expected response " + err.Error())

	}	


	if response.Code != http.StatusOK {
	
		t.Errorf("Expected: %d , Got: %d",http.StatusOK,response.Code )
	}	
	
	got,err := io.ReadAll(response.Body)
	if err != nil{
		
		t.Errorf("Error reading response: " + err.Error())
	
	}
	if !bytes.Equal(expectedJson,got){

		t.Errorf("Expected json was not returned")
	}	
}
