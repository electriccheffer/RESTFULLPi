package handler_test 

import "testing"
import "path/filepath"
import "crypto/rand"
import "encoding/hex"
import "restfulpi/internal/handler"

func TestSessionManagerSuccess(t *testing.T){

	writePath := filepath.Join("..","testData","gpslogdata")
	readPath := filepath.Join("..","testData","gpsreaddata")
	sessionManager := handler.NewSessionManager(writePath,readPath)
	filePath := "09_08_2026_12_33_05.gpx"
	randomBytes := make([]byte,16)
	_,err := rand.Read(randomBytes)
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
