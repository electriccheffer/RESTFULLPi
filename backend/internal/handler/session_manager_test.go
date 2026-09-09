package handler_test 

import "testing"
import "path/filepath"
import "crypto/rand"
import "encoding/hex"
import "os"
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


type NoPermissionsOpener struct{}

func NewNoPermissionsOpener()file_operations.Opener{

	npo := &NoPermissionsOpener{}
	return npo
}

func (npo *NoPermissionsOpener) Open(path string)(file_operations.FileHandle,error){
	
	return nil,os.ErrPermission	
}

func (npo *NoPermissionsOpener) Create(path string)(file_operations.FileHandle,error){

	return nil,os.ErrPermission	
}

