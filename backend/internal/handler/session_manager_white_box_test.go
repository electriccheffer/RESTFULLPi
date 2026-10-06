package handler 

import "testing"
import "context"
import "io"
import "bytes"
import "time"
import "errors"
import "os"
import "math"
import "encoding/xml"
import "path/filepath"
import "restfulpi/internal/file_operations"
import "restfulpi/internal/models"

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

func TestReadNMEASuccess(t *testing.T){
	
	readNMEAContext, cancel := context.WithCancel(context.Background())
	simulatedReadFile,writer  := io.Pipe()
	defer simulatedReadFile.Close()
	defer writer.Close() 
	
	opener := file_operations.NewFileOpener()	
	dir := t.TempDir()
	writePath := filepath.Join(dir,"WriteFile.gpx")
	writeFile, err := opener.Create(writePath)
	if err != nil {
		t.Errorf("unexpected error: %s",err.Error())
	}
	sessionManager := NewSessionManager(dir,dir,opener)
	
	errorChannel := make(chan error,1)

	go func(){
		validSentence := "$GPRMC,123519.50,A,4807.038,"+
			 "N,01131.000,E,022.4,084.4,210926,003.1,W*40\r\n"
		byteSentence := []byte(validSentence)
		for i := 0 ; i < 5 ; i++{
			writer.Write(byteSentence)	
		}
		cancel()  
	}()
	
	go func(){
		errorChannel <- sessionManager.readNMEA(readNMEAContext,
							simulatedReadFile,
							writeFile)
	
	}()
	
	select {
		
		case err := <-errorChannel:
			if !errors.Is(err,context.Canceled){
			
				t.Errorf("Expected: Context.Canceled Got:%v ",err)
			}
			_ = writeFile.Close()
			fileContents, err := os.Open(writePath)
			if err != nil {
				t.Errorf("unexpected error: %s",err.Error())
			}
			defer fileContents.Close()
			var gpxFile models.GPSFile
			decoder := xml.NewDecoder(fileContents)	
			err = decoder.Decode(&gpxFile)
			if err != nil{
				t.Errorf("unexpected error: %s", err.Error())
			}
			expectedLatitude := 48.1173
			latitude := gpxFile.Track.Segment.Points[0].Latitude
			delta := 0.000001
			if math.Abs(expectedLatitude - latitude) > delta{
				t.Errorf("expected: %f got: %f",expectedLatitude,
								latitude)
			}
		case <- time.After(1 * time.Second):
			t.Error("Time exceeded.")
	}
}

//TODO: Test Error case for readNEMA bad serial port 
//TODO: Test Error case for readNEMA bad write case 
//TODO: Test Error cases for bad NEMA sentences bad date 
//TODO: Test Error case for bad NEMA sentences invalid 
//TODO: Test Error case for bad NEMA sentences bad checksum 
//TODO: Test Error case for bad NEMA sentences bad lat 
//TODO: Test Error case for bad NEMA sentences bad long
//TODO: Test Error case for bad NEMA sentences blank fields
