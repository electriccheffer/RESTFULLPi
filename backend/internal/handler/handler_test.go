package handler

import "testing"
import "path/filepath"
import "net/http"
import "net/http/httptest"
import "os"
import "io"
import "bytes"
import "restfulpi/internal/server"

func TestAngularAppDelivery(t *testing.T){
	
	buildPath := filepath.Join("..","..","testData","browser","index.html")
	handler := NewFrontendHandler(buildPath)
	router := router.NewRouter(handler)
	
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
	router := router.NewRouter(handler)
	
	request := httptest.NewRequest(http.MethodGet,"/",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	if response.Code != http.StatusNotFound{
			
		t.Errorf("Expected: %d, Got: %d",http.StatusNotFound,response.Code)
	}

}
