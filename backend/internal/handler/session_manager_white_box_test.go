package handler 

import "testing"
import "context"
import "io"
import "bytes"
import "time"
import "errors"
import "restfulpi/internal/file_operations"

func TestReadNMEACancellation(t *testing.T){

	readNMEAContext, cancel := context.WithCancel(context.Background())
	
	simulatedReader,simulatedWriter := io.Pipe()
	defer simulatedReader.Close()
	defer simulatedWriter.Close()
	
	var outputBuffer bytes.Buffer 

	opener := file_operations.NewFileOpener()
	
	sessionManager := NewSessionManager(t.TempDir(),t.TempDir(),opener)
	
	errorChannel := make(chan error,1)
	
	go func(){
		
		errorChannel <- sessionManager.readNMEA(readNMEAContext,
							simulatedReader,
							&outputBuffer)
	}()			
	
	cancel() 	
	select {
		
		case err := <-errorChannel:
			if !errors.Is(err,context.Canceled){
			
				t.Errorf("Expected: Context.Canceled Got:%v ",err)
			}
		case <- time.After(1 * time.Second):
			t.Error("Time exceeded.")
	}
}	
