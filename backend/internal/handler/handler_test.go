package handler_test

import "testing"
import "path/filepath"
import "net/http"
import "net/http/httptest"
import "os"
import "io"
import "bytes"
import "encoding/json"
import "regexp"

import "restfulpi/internal/models"
import "restfulpi/internal/server"

func TestAngularAppDelivery(t *testing.T){
	
	buildPath := filepath.Join("..","server","dist","browser","index.html")
	
	sessionManager := NewSessionStartSuccess()
	
	router := router.NewRouter(buildPath,sessionManager)
	
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

	buildPath := filepath.Join("..","server","dist","browser","index.html")
	
	sessionManager := NewSessionStartSuccess()
	router := router.NewRouter(buildPath,sessionManager)
	expected,err := os.ReadFile(buildPath)
	if err != nil {
		t.Errorf("error reading file:" + err.Error())
	}
	expectedSlice := expected[1:20] 

	request := httptest.NewRequest(http.MethodGet,"/noexist",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	if response.Code != http.StatusOK{
			
		t.Errorf("Expected: %d, Got: %d",http.StatusOK,response.Code)
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

func TestAngularAppDeliveryFileIsDirectory(t *testing.T){

	buildPath := filepath.Join("..","server","dist","browser","index.html")
	
	sessionManager := NewSessionStartSuccess()
	router := router.NewRouter(buildPath,sessionManager)
	expected,err := os.ReadFile(buildPath)
	if err != nil {
		t.Errorf("error reading file:" + err.Error())
	}
	expectedSlice := expected[1:20] 


	request := httptest.NewRequest(http.MethodGet,"/dist",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	if response.Code != http.StatusOK{
		t.Errorf("Expected: %d, Got: %d",http.StatusOK,response.Code)
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

func TestStatusHandler(t *testing.T){

	buildPath := filepath.Join("..","handler")

	sessionManager := NewSessionStartSuccess()
	router := router.NewRouter(buildPath,sessionManager)
	
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

	sessionManager := NewSessionStartSuccess()
	router := router.NewRouter(path,sessionManager)
	
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

	sessionManager := NewSessionStartSuccess()
	router := router.NewRouter(path,sessionManager)
	
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

	sessionManager := NewSessionStartSuccess()
	router := router.NewRouter(path,sessionManager)
	
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

	sessionManager := NewSessionStartSuccess()
	router := router.NewRouter(path,sessionManager)
	
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
	
	path := filepath.Join(".","NoExist")

	sessionManager := NewSessionStartSuccess()
	router := router.NewRouter(path,sessionManager)
	
	request := httptest.NewRequest(http.MethodGet,"/logs",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)

	_,err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response: " + err.Error())
	}
	
	if response.Code != http.StatusNotFound{
		t.Errorf("Expected:%d Got:%d",http.StatusNotFound,response.Code)
	}	
	
}

func TestGetLogsHandlerFailNotDirectory(t *testing.T){

	path := filepath.Join(".","handler_test.go")

	sessionManager := NewSessionStartSuccess()
	router := router.NewRouter(path,sessionManager)
	
	request := httptest.NewRequest(http.MethodGet,"/logs",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)

	_,err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response: " + err.Error())
	}
	
	if response.Code != http.StatusNotFound{
		t.Errorf("Expected:%d Got:%d",http.StatusNotFound,response.Code)
	}

}

type SessionStartSuccess struct{}

func NewSessionStartSuccess() *SessionStartSuccess{
	sss := &SessionStartSuccess{}
	return sss
}

func (sss *SessionStartSuccess) StartSession(id string, filePath string) (*models.Session,error){
	
	session := &models.Session{FileName:filePath,Id:id}
	
	return session,nil
}

func TestSessionStartHandlerSuccess(t *testing.T){

	buildPath := filepath.Join("..","server","dist","browser","index.html")

	sessionManager := NewSessionStartSuccess()
	router := router.NewRouter(buildPath,sessionManager)
	
	request := httptest.NewRequest(http.MethodPost,"/logs/sessions",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	
	got,err := io.ReadAll(response.Body)
	
		
	if err != nil {
		t.Errorf("Error reading response:%s " , err.Error())	
	}

	if response.Code != http.StatusCreated{
		t.Errorf("Expected: %d Got: %d",http.StatusCreated,response.Code)
	}

	var session models.Session
	
	err = json.Unmarshal(got,&session)

	if err != nil {

		t.Errorf("Error unmarshaling session object:%s ", err.Error())
	}

	lengthOfId := 32
	if len(session.Id) != lengthOfId {
		t.Errorf("Incorrect length of id. expected: %d got:%d",
			 lengthOfId,len(session.Id))
	}

	fileNamePattern := `^\d{2}_\d{2}_\d{2}_\d{2}_\d{2}_\d{2}\.gpx$`
	regularExpression := regexp.MustCompile(fileNamePattern)
	if !regularExpression.MatchString(session.FileName){
	
		t.Errorf("File name did not match pattern got: %s",session.FileName)

	}
}
