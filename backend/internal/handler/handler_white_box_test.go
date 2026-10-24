package handler

import "encoding/hex"
import "crypto/rand"
import "io"
import "testing"
import "time"
import "path/filepath"
import "net/http"
import "net/http/httptest"
import "encoding/json"
import "regexp"
import "os"
import "restfulpi/internal/models"
import "restfulpi/internal/file_operations"

func TestSessionHandlerIntegrationSuccessCase(t *testing.T){

	readDirectory := t.TempDir()
	writeDirectory := t.TempDir()
	readFilePath := filepath.Join(readDirectory,"device")
	fileOpener := file_operations.NewFileOpener()
	_,err := fileOpener.Create(readFilePath)
	if err != nil{

		t.Errorf("unexpected error creating read path: %s",err.Error())
	}	

	sessionManager := NewSessionManager(writeDirectory,readFilePath,fileOpener)
	
	handler := NewSessionStartHandler(sessionManager)
	
	request := httptest.NewRequest(http.MethodPost,"/logs/sessions",nil)
	response := httptest.NewRecorder()
			
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

func TestSessionStartHandlerIntegrationRedundantIdSuccess(t *testing.T){

	readDirectory := t.TempDir()
	writeDirectory := t.TempDir()
	readFilePath := filepath.Join(readDirectory,"device")
	fileOpener := file_operations.NewFileOpener()
	_,err := fileOpener.Create(readFilePath)
	if err != nil{

		t.Errorf("unexpected error creating read path: %s",err.Error())
	}	

	sessionManager := NewSessionManager(writeDirectory,readFilePath,fileOpener)
		
	handler := NewSessionStartHandler(sessionManager)
	
	nonRandomBytes := make([]byte,16) 
	for i := 0 ; i < 16 ; i++{
		nonRandomBytes[i] = byte(i)
	}
	failedCallCount := 0
	handler.randomReader = func(b []byte)(int,error){
		
		length := len(b)	
		if failedCallCount < 1{
			for i := range length{
				b[i] = byte(i)
			}
			failedCallCount++
			return length,nil
		}
		
		return rand.Read(b)
	}
	id := hex.EncodeToString(nonRandomBytes)
	sessionStandin := &models.Session{}
	sessionManager.sessions[id] = sessionStandin
	request := httptest.NewRequest(http.MethodPost,"/logs/sessions",nil)
	response := httptest.NewRecorder()
			
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

func TestSessionStartHandlerIntegrationRedundantIdFailure(t *testing.T){

	readDirectory := t.TempDir()
	writeDirectory := t.TempDir()
	readFilePath := filepath.Join(readDirectory,"device")
	fileOpener := file_operations.NewFileOpener()
	_,err := fileOpener.Create(readFilePath)
	if err != nil{

		t.Errorf("unexpected error creating read path: %s",err.Error())
	}	

	sessionManager := NewSessionManager(writeDirectory,readFilePath,fileOpener)
		
	handler := NewSessionStartHandler(sessionManager)
	
	nonRandomBytes := make([]byte,16) 
	for i := 0 ; i < 16 ; i++{
		nonRandomBytes[i] = byte(i)
	}
	handler.randomReader = func(b []byte)(int,error){
		
		length := len(b)	
		for i := range length{
			b[i] = byte(i)
		}
		return length,nil
		
	}
	id := hex.EncodeToString(nonRandomBytes)
	sessionStandin := &models.Session{}
	sessionManager.sessions[id] = sessionStandin
	request := httptest.NewRequest(http.MethodPost,"/logs/sessions",nil)
	response := httptest.NewRecorder()
			
	handler.ServeHTTP(response,request)
	
	if response.Code != http.StatusConflict {
		t.Errorf("incorrect status code expected:%d got:%d",
							http.StatusConflict,
							response.Code)
	}
	
}

func TestSessionStartHandlerIntegrationRedundantFileSucccess(t *testing.T){

	readDirectory := t.TempDir()
	writeDirectory := t.TempDir()
	readFilePath := filepath.Join(readDirectory,"device")
	
	timeStamp := time.Date(2026,10,21,16,0,0,0,time.UTC)
	badFileName := timeStamp.Format("01_02_06_15_04_05.gpx")
	badWritePath := filepath.Join(writeDirectory,badFileName)	
	fileOpener := file_operations.NewFileOpener()
	_,err := fileOpener.Create(readFilePath)
	if err != nil {
		t.Errorf("unexpected error creating read path: %s",err.Error())
	}
	
	_,err = fileOpener.Create(badWritePath)
	if err != nil {
		t.Errorf("unexpected error creating write path: %s",err.Error())
	}
	
	sessionManager := NewSessionManager(writeDirectory,readFilePath,fileOpener)
	handler := NewSessionStartHandler(sessionManager)
	handler.timeMaker = func()(time.Time){
		return timeStamp			
	}
	
	request := httptest.NewRequest(http.MethodPost,"/logs/sessions",nil)
	response := httptest.NewRecorder()
			
	handler.ServeHTTP(response,request)
	
	if response.Code != http.StatusCreated {
		t.Errorf("incorrect status code expected:%d got:%d",
							http.StatusCreated,
							response.Code)
	}
	
	timeStamp = time.Date(2026,10,21,16,0,0,1,time.UTC)
	badFileName = timeStamp.Format("01_02_06_15_04_05.gpx")
	badWritePath = filepath.Join(writeDirectory,badFileName)
	_,err = os.Stat(badWritePath)
	if err != nil {
		t.Errorf("file not created error code: %s",err.Error())
	}
}

func TestSesionStartHandlerIntegrationRedundantFileFailure(t *testing.T){

	readDirectory := t.TempDir()
	writeDirectory := t.TempDir()
	
	fileOpener := file_operations.NewFileOpener()
	readFilePath := filepath.Join(readDirectory,"device")
	
	timeStampOne := time.Date(2026,10,21,16,0,0,0,time.UTC)
	timeStampTwo := timeStampOne.Add(1 * time.Second)
	timeStampThree := timeStampTwo.Add(1 * time.Second)
	timeStampFour := timeStampThree.Add(1 * time.Second) 

	badFileNameOne := timeStampOne.Format("01_02_06_15_04_05.gpx")
	badFileNameTwo := timeStampTwo.Format("01_02_06_15_04_05.gpx")
	badFileNameThree := timeStampThree.Format("01_02_06_15_04_05.gpx")
	badFileNameFour := timeStampFour.Format("01_02_06_15_04_05.gpx")
	
	badWritePathOne := filepath.Join(writeDirectory,badFileNameOne)
	badWritePathTwo := filepath.Join(writeDirectory,badFileNameTwo)
	badWritePathThree := filepath.Join(writeDirectory,badFileNameThree)
	badWritePathFour := filepath.Join(writeDirectory,badFileNameFour)
	
	_,err := fileOpener.Create(badWritePathOne)
	if err != nil {
		t.Errorf("unexpected error creating write path: %s",err.Error())
	}
	_,err = fileOpener.Create(badWritePathTwo)
	if err != nil {
		t.Errorf("unexpected error creating write path: %s",err.Error())
	}
	_,err = fileOpener.Create(badWritePathThree)
	if err != nil {
		t.Errorf("unexpected error creating write path: %s",err.Error())
	}
	_,err = fileOpener.Create(badWritePathFour)
	if err != nil {
		t.Errorf("unexpected error creating write path: %s",err.Error())
	}
	
	_,err = fileOpener.Create(readFilePath)
	if err != nil{
		t.Errorf("unexpected error creating read path: %s",err.Error())
	}
	
	numberOfTries := 0
	sessionManager := NewSessionManager(writeDirectory,readFilePath,fileOpener)
	handler := NewSessionStartHandler(sessionManager)
	handler.timeMaker = func()(time.Time){
		switch numberOfTries {
			
			case 0:
				numberOfTries++ 
				return timeStampOne 
			case 1:
				numberOfTries++ 
				return timeStampTwo 
			case 2:
				numberOfTries++ 
				return timeStampThree
			case 3:
				numberOfTries++ 
				return timeStampFour
			
		}
		return timeStampOne
	}
	
	request := httptest.NewRequest(http.MethodPost,"/logs/sessions",nil)
	response := httptest.NewRecorder()
			
	handler.ServeHTTP(response,request)
	
	if response.Code != http.StatusConflict {
		t.Errorf("incorrect status code expected:%d got:%d",
							http.StatusConflict,
							response.Code)
	}	
	
}

func TestSessionStartHandlerIntegrationNoReadPath(t *testing.T){
	
	readDirectory := t.TempDir()
	writeDirectory := t.TempDir()
	readFilePath := filepath.Join(readDirectory,"device")
			
	fileOpener := file_operations.NewFileOpener()
	sessionManager := NewSessionManager(writeDirectory,readFilePath,fileOpener)
	handler := NewSessionStartHandler(sessionManager)
	
	request := httptest.NewRequest(http.MethodPost,"/logs/sessions",nil)
	response := httptest.NewRecorder()
			
	handler.ServeHTTP(response,request)
	if response.Code != http.StatusInternalServerError {
		t.Errorf("incorrect status returned expected:%d got:%d",
						http.StatusInternalServerError,response.Code)
	}
}
//TODO: Integration test write path doesn't exist 
//TODO: Integration test read path permissions error
//TODO: Integration test write path permissions error
