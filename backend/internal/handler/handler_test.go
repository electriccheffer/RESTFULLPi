package handler_test

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
	
	router := router.NewRouter(buildPath,buildPath)
	
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
	router := router.NewRouter(buildPath,buildPath)
	
	request := httptest.NewRequest(http.MethodGet,"/",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	if response.Code != http.StatusNotFound{
			
		t.Errorf("Expected: %d, Got: %d",http.StatusNotFound,response.Code)
	}

}

func TestAngularAppDeliveryFileIsDirectory(t *testing.T){

	buildPath := filepath.Join("..","handler")
	router := router.NewRouter(buildPath,buildPath)
	
	request := httptest.NewRequest(http.MethodGet,"/",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	if response.Code != http.StatusNotAcceptable{
		t.Errorf("Expected: %d, Got: %d",http.StatusNotAcceptable,response.Code)
	}
}

func TestStatusHandler(t *testing.T){

	buildPath := filepath.Join("..","handler")
	router := router.NewRouter(buildPath,buildPath)
	
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

func TestGetLogsHandlerSuccessEmpty(t *testing.T){
	
	path := filepath.Join("..","..","testData","EmptyDirectory")
	router := router.NewRouter(path,path)
	
	request := httptest.NewRequest(http.MethodGet,"/logs",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	
	got,err := io.ReadAll(response.Body)
	if err != nil{
		t.Errorf("Error reading response: " + err.Error())
	}
	
	var gotArray []models.Logs
	err = json.Unmarshal(got,&gotArray)
	if err != nil{
		t.Errorf("Error Unmarshling response: %s",err.Error())
	}
	if len(gotArray) != 0 {
		t.Errorf("Length of array response was not zero")
	}
}

func TestGetLogsHandlerSuccessOneFile(t *testing.T){
	
	path := filepath.Join("..","..","testData","OneFileDirectory")
	router := router.NewRouter(path,path)
	
	request := httptest.NewRequest(http.MethodGet,"/logs",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
		
	got,err := io.ReadAll(response.Body)
	if err != nil{
		t.Errorf("Error reading response: " + err.Error())
	}	
	
	expectedName := "2006-01-02.gpx" 
	expectedLog :=  models.Logs{expectedName}
	expectedLogArray := []models.Logs{expectedLog}
	expectedMarshal, err := json.Marshal(expectedLogArray)
	if err != nil{
		t.Errorf("Difficulty with marshaling expected log")
	}
	if !bytes.Equal(expectedMarshal,got){
		t.Errorf("Incorrect data returned from one file")
	}
}

func TestGetLogsHandlerSuccessTwoFiles(t *testing.T){
	
	path := filepath.Join("..","..","testData","DirectoryWithTwoFiles")
	router := router.NewRouter(path,path)
	
	request := httptest.NewRequest(http.MethodGet,"/logs",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	
	got,err := io.ReadAll(response.Body)
	if err != nil{
		t.Errorf("Error reading response: " + err.Error())
	}

	expectedName := "2006-01-02.gpx" 
	expectedLog := models.Logs{expectedName}
	expectedSecondName := "2006.gpx"
	expectedSecondLog := models.Logs{expectedSecondName}
	expectedLogArray := []models.Logs{expectedLog,expectedSecondLog}
	expectedMarshal, err := json.Marshal(expectedLogArray)
	if err != nil{
		t.Errorf("Dirriculty with marshaling expected log")
	}
	if !bytes.Equal(expectedMarshal,got){
		t.Errorf("Incorrect data returned from file")
	}
}

func TestGetLogsHandlerSuccessFilesWithDirectory(t *testing.T){

	path := filepath.Join("..","..","testData","DirectoryWithTwoFilesAndDirectory")
	router := router.NewRouter(path,path)
	
	request := httptest.NewRequest(http.MethodGet,"/logs",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	
	got,err := io.ReadAll(response.Body)
	if err != nil{
		t.Errorf("Error reading response: " + err.Error())
	}

	expectedName := "2006-01-02.gpx" 
	expectedLog := models.Logs{expectedName}
	expectedSecondName := "2006.gpx"
	expectedSecondLog := models.Logs{expectedSecondName}
	expectedLogArray := []models.Logs{expectedLog,expectedSecondLog}
	expectedMarshal, err := json.Marshal(expectedLogArray)
	if err != nil{
		t.Errorf("Dirriculty with marshaling expected log")
	}
	if !bytes.Equal(expectedMarshal,got){
		t.Errorf("Incorrect data returned from file")
	}
	
}

func TestGetLogsHandlerFailNoExist(t *testing.T){

}

func TestGetLogsHandlerFailNotDirectory(t *testing.T){



}

