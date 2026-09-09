package handler_test 

import "testing"
import "path/filepath"
import "crypto/rand"
import "encoding/hex"
import "os"
import "errors"
import "syscall"
import "restfulpi/internal/handler"
import "restfulpi/internal/file_operations"

func TestSessionManagerSuccess(t *testing.T){
	
	testDir := t.TempDir()
	testReadPath := filepath.Join(testDir,"read.gpx")
	testWritePath := filepath.Join(testDir,"write.gpx")
	readFile,err := os.Create(testReadPath)
	if err != nil{
		
		t.Errorf("Error creating read file:%s ",testReadPath)
	
	}
	writeFile,err := os.Create(testWritePath)
	if err != nil{
		
		t.Errorf("Error creating write file:%s ",testWritePath)
	
	}
	defer readFile.Close()
	defer writeFile.Close()
	
	fileOpener := file_operations.NewFileOpener()
	sessionManager := handler.NewSessionManager(testWritePath,testReadPath,fileOpener)
	filePath := "09_08_2026_12_33_05.gpx"
	randomBytes := make([]byte,16)
	_,err = rand.Read(randomBytes)
	if err != nil {
		t.Errorf("Error reading random bytes: %s \n",err.Error())	
	}
	id := hex.EncodeToString(randomBytes)		
	
	session, err := sessionManager.StartSession(id,filePath)
	if err != nil {
		t.Errorf("Error starting session: %s \n",err.Error())
	}
	
	if session.FileName != filePath {
		t.Errorf("Paths do not match. expected:%s got:%s \n",filePath,session.FileName)
	}

	expectedIdLength := 32
	lengthOfId := len(session.Id)
	if expectedIdLength != lengthOfId {
		t.Errorf("Id of incorrect length. Expected:%d Got: %d \n",
								expectedIdLength, lengthOfId)
	}	
}


type NoPermissionsOpenerWrite struct{}

func NewNoPermissionsOpenerWrite()file_operations.Opener{

	npo := &NoPermissionsOpenerWrite{}
	return npo
}

func (npo *NoPermissionsOpenerWrite) Open(path string)(file_operations.FileHandle,error){
	
	return os.Open(path)	
}

func (npo *NoPermissionsOpenerWrite) Create(path string)(file_operations.FileHandle,error){

	return nil,os.ErrPermission	
}

func TestSessionManagerStartSessionNoPermissions(t *testing.T){
	
	testDirectory := t.TempDir()
	readPath := filepath.Join(testDirectory,"read.gpx")
	writePath := filepath.Join(testDirectory) 
	testReadFile, err := os.Create(readPath)
	if err != nil{
		t.Errorf("Error creating read file: %s",err.Error())
	}
	defer testReadFile.Close()
	 
	noPermissionsOpenerWrite := NewNoPermissionsOpenerWrite()
	sessionManager := handler.NewSessionManager(writePath,readPath,noPermissionsOpenerWrite)
	
	randomBytes := make([]byte,16)
	_,err = rand.Read(randomBytes)
	if err != nil {
		t.Errorf("Error reading random bytes: %s \n",err.Error())	
	}
	id := hex.EncodeToString(randomBytes)
	writeFilePath := "write.gpx"
	_,err = sessionManager.StartSession(id,writeFilePath)
	if err == nil {
		t.Error("No error thrown in no permissions case")
	}
	if err != nil {
		var managerError *handler.SessionManagerError
		
		if errors.As(err,&managerError){
			if managerError.Code != int(syscall.EACCES) {
				t.Errorf("Incorrect error code expected:%d got:%d \n",
									os.ErrPermission,
									managerError.Code)
			}
		}else{t.Error("Error of incorrect type")}		
			
	}
	
	
}

