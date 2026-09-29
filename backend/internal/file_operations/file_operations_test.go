package file_operations_test

import  "path/filepath"
import "testing"
import "math"
import "errors"
import "time"
import  "restfulpi/internal/file_operations"

func TestDirectoryReadSuccess(t *testing.T){
			
	path := filepath.Join("..","..","testData","DirectoryWithFiles")
	directoryReader := file_operations.NewDirectoryRead(path)
	files,err := directoryReader.GetFiles()
	if err != nil{
	
		t.Errorf("Error while reading directory: " + err.Error())
	}
	
	expectedFiles := [2]string{"fileOne","fileTwo"}
		
	totalExpectedFiles := len(expectedFiles)

	if len(files) != totalExpectedFiles{
		t.Errorf("Improper number of files. Expected:%d Got:%d",
			  totalExpectedFiles,len(files) )
	}

	allFound := true 
	for _, file := range files {
		
		for expectedIndex, expectedFile := range expectedFiles {
			
			if file.Name() == expectedFile{
				break
			}
			if expectedIndex == totalExpectedFiles - 1 {
				allFound = false	
			}
			
		}
		if(allFound == false){
				
			t.Errorf("File not found:%s ", file.Name())
			break
		}
	}
	
}

func TestDirectoryReadFailureIsFile(t *testing.T){

	path := filepath.Join(".","file_operations.go")
	directoryReader := file_operations.NewDirectoryRead(path)
	
	_,err := directoryReader.GetFiles()
	if err == nil {
		t.Errorf("Should throw error directory is file")
	}
}

func TestDirectoryReadEmptyDirectory(t *testing.T){

	path := filepath.Join("..","..","testData","EmptyDirectory")
	directoryReader := file_operations.NewDirectoryRead(path)
	files,err := directoryReader.GetFiles()
	if err != nil{
		t.Errorf("Should not throw error. Error is:%s ", err.Error())	
	}
	if len(files) != 0{
		t.Errorf("The directory is not empty")
	}	
}

func TestDirectoryReadDirectoryDoesNotExist(t *testing.T){

	path := filepath.Join(".","DoesntExist")
	directoryReader := file_operations.NewDirectoryRead(path)
	_,err := directoryReader.GetFiles()
	if err == nil {
		t.Errorf("Directory does not exist. Should throw error")
	}
}


func TestParseSentenceSuccess(t *testing.T){
	
	validSentence := "$GPRMC,123519.50,A,4807.038,"+
			 "N,01131.000,E,022.4,084.4,210926,003.1,W*40"
	
	expectedLatitude := 48.117300
	expectedLongitude := 11.516667
	expectedTime := time.Date(2026,time.September,21,12,35,19,500000000,time.UTC)	
	
	parser := file_operations.NewGPSParser()
	trackPoint, err := parser.ParseSentence(validSentence)
	if err != nil {
	
		t.Errorf("unexpected error: %s",err.Error())
	}
	
	delta := .000001
	if math.Abs(trackPoint.Latitude - expectedLatitude) > delta{
		t.Errorf("expected: %f got: %f",expectedLatitude,trackPoint.Latitude)
	}
	if math.Abs(trackPoint.Longitude - expectedLongitude) > delta{
		t.Errorf("expected: %f got: %f",expectedLongitude,trackPoint.Longitude)
	}
	if trackPoint.Time != expectedTime {
		t.Errorf("expected:%v got:%v",expectedTime,trackPoint.Time)
	}
}

func TestParseSentenceInvalidChecksum(t *testing.T){
	
	invalidSentence := "$GPRMC,123519.50,A,4807.038," +
			 "N,01131.000,E,022.4,084.4,210926,003.1,W*44"
	
	parser := file_operations.NewGPSParser()
	trackPoint, err := parser.ParseSentence(invalidSentence)
	if trackPoint != nil {
		t.Error("ParseSentence did not return nil. Expected error.")
	}
	if err != nil{
		var parserErr *file_operations.GPSParserError
		if !errors.As(err,&parserErr){
			t.Error("ParseSentence incorrect error type")
		}
		if parserErr.Code != 74 {
			t.Errorf("Incorrect error code returned expected: %d got: %d",
				74,parserErr.Code)	
		}
	}
	
} 

func TestParseSentenceInvalid(t *testing.T){

	invalidSentence := "$GPRMC,123519.50,V,4807.038,"+
			   "N,01131.000,E,022.4,084.4,210926,003.1,W*57"
	
	parser := file_operations.NewGPSParser()
	trackPoint, err := parser.ParseSentence(invalidSentence)

	if trackPoint != nil {
		t.Error("ParseSentence did not return nil. Expected error.")
	}
	if err != nil{
		var parserErr *file_operations.GPSParserError
		if !errors.As(err,&parserErr){
			t.Error("ParseSentence incorrect error type")
		}
		if parserErr.Code != 42 {
			t.Errorf("Incorrect error code returned expected: %d got: %d",
				42,parserErr.Code)	
		}
	}

}

func TestParseSentenceBadLatitude(t *testing.T){

	invalidSentence := "$GPRMC,123519.50,A,480A.038,"+
			 "N,01131.000,E,022.4,084.4,210926,003.1,W*36"
	
	parser := file_operations.NewGPSParser()
	trackPoint, err := parser.ParseSentence(invalidSentence)

	if trackPoint != nil {
		t.Error("ParseSentence did not return nil. Expected error.")
	}
	if err != nil{
		var parserErr *file_operations.GPSParserError
		if !errors.As(err,&parserErr){
			t.Error("ParseSentence incorrect error type")
		}
		if parserErr.Code != 22 {
			t.Errorf("Incorrect error code returned expected: %d got: %d",
				22,parserErr.Code)	
		}
		if parserErr.Error() != "GPSParserError 22: error parsing latitude" {
			t.Errorf("incorrect message expected:%s got:%s",
						"GPSParserError 22: error parsing latitude",
						parserErr)
		}
	}
	
}

func TestParseSentenceBadLongitude(t *testing.T){
	
	invalidSentence := "$GPRMC,123519.50,A,4807.038,"+
			 "N,0113A.000,E,022.4,084.4,210926,003.1,W*30"
		
	parser := file_operations.NewGPSParser()
	trackPoint, err := parser.ParseSentence(invalidSentence)

	if trackPoint != nil {
		t.Error("ParseSentence did not return nil. Expected error.")
	}
	if err != nil{
		var parserErr *file_operations.GPSParserError
		if !errors.As(err,&parserErr){
			t.Error("ParseSentence incorrect error type")
		}
		if parserErr.Code != 22 {
			t.Errorf("Incorrect error code returned expected: %d got: %d",
				22,parserErr.Code)	
		}
		if parserErr.Error() != "GPSParserError 22: error parsing longitude" {
			t.Errorf("incorrect message expected:%s got:%s",
						"GPSParserError 22: error parsing latitude",
						parserErr)
		}
	}
	
}

//TODO: Test parser bad lat hemisphere
func TestParseSentenceBadLatitudeHemisphere(t *testing.T){
	
	invalidSentence := "$GPRMC,123519.50,A,4807.038,"+
			 "E,01131.000,E,022.4,084.4,210926,003.1,W*4B"
	
	parser := file_operations.NewGPSParser()
	trackPoint, err := parser.ParseSentence(invalidSentence)

	if trackPoint != nil {
		t.Error("ParseSentence did not return nil. Expected error.")
	}
	if err != nil{
		var parserErr *file_operations.GPSParserError
		if !errors.As(err,&parserErr){
			t.Error("ParseSentence incorrect error type")
		}
		if parserErr.Code != 22 {
			t.Errorf("Incorrect error code returned expected: %d got: %d",
				22,parserErr.Code)	
		}
		if parserErr.Error() != "GPSParserError 22: invalid hemisphere" {
			t.Errorf("incorrect message expected:%s got:%s",
						"GPSParserError 22: invalid hemisphere",
						parserErr)
		}
	}	
}
//TODO: Test parser bad long hemisphere
//TODO: Test parser bad time 
//TODO: Test parser missing fields valid checksum 
