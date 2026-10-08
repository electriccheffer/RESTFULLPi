package handler 

import "testing"
import "context"
import "io"
import "bytes"
import "time"
import "errors"
import "os"
import "math"
import "syscall"
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

func TestMultipleLinesSuccess(t *testing.T){
	
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
		invalidSentence := "$GPVTG,054.7,T,034.4,M,005.5,N,010.2,K,A*48\r\n"	
		byteSentence := []byte(validSentence)
		invalidByteSentence := []byte(invalidSentence)
		for i := 0 ; i < 5 ; i++{
			if i == 2 {
				writer.Write(invalidByteSentence)
				continue
			}
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
			if len(gpxFile.Track.Segment.Points) != 4{
				t.Errorf("expected length:%d got: %d",4,
						len(gpxFile.Track.Segment.Points))
			}	
			delta := 0.000001
			if math.Abs(expectedLatitude - latitude) > delta{
				t.Errorf("expected: %f got: %f",expectedLatitude,
								latitude)
			}
		case <- time.After(1 * time.Second):
			t.Error("Time exceeded.")
	}

}

func TestErrorCaseReadNMEABadSerialPort(t *testing.T){
	
	readNMEAContext, _ := context.WithCancel(context.Background())
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
			if i == 2 {
				writer.CloseWithError(io.ErrUnexpectedEOF)	
				break
			}
			writer.Write(byteSentence)	
			
		}
	}()
	
	go func(){
		errorChannel <- sessionManager.readNMEA(readNMEAContext,
							simulatedReadFile,
							writeFile)
	
	}()
	
	select {
		
		case err := <-errorChannel:
			if !errors.Is(err,io.ErrUnexpectedEOF){
			
				t.Errorf("Expected: EOF error Got:%v ",err)
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
			if len(gpxFile.Track.Segment.Points) != 2{
				t.Errorf("expected length:%d got: %d",2,
						len(gpxFile.Track.Segment.Points))
			}
			if math.Abs(expectedLatitude - latitude) > delta{
				t.Errorf("expected: %f got: %f",expectedLatitude,
								latitude)
			}
		case <- time.After(1 * time.Second):
			t.Error("Time exceeded.")
	}

}

type ErrorWriter struct{
	writer io.Writer
	failAt int
	bytesWritten int
	err error
}

func NewErrorWriter(writer io.Writer ,failAt int, err error)io.Writer{
	
	ew := &ErrorWriter{writer:writer,failAt:failAt,
			   bytesWritten:0, err: err}
	return ew
}

func (ew *ErrorWriter) Write( content []byte)(int,error){
	ew.bytesWritten += len(content)	
	if ew.bytesWritten >= ew.failAt {
		return 0,ew.err 
	}else{
		return ew.writer.Write(content)
	}
}



func TestReadNMEABadWriterErrorCase(t *testing.T){

	readNMEAContext, cancel := context.WithCancel(context.Background())
	simulatedReadFile,writer  := io.Pipe()
	defer simulatedReadFile.Close()
	defer writer.Close() 
	
	opener := file_operations.NewFileOpener()	
	dir := t.TempDir()
	writePath := filepath.Join(dir,"WriteFile.gpx")
	writeFile, err := opener.Create(writePath)
	// Create the error writer instance
	writerError := syscall.EIO
	headerContent := `<?xml version="1.0" encoding="UTF-8"?><gpx version="1.1"
		   creator="RESTFULPi" xmlns="http://www.topografix.com/GPX/1/1"><trk><trkseg>`
	failPoint := len(headerContent) + 4
	errorWriter := NewErrorWriter(writeFile,failPoint,writerError)
	
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
							errorWriter)
	
	}()
	
	select {
		
		case err := <-errorChannel:
			if !errors.Is(err,syscall.EIO){
			
				t.Errorf("Expected: IO ErrorGot:%v ",err)
			}
			
		case <- time.After(1 * time.Second):
			t.Error("Time exceeded.")
	}
}

func TestBadNMEASentencesBadDateParserIntegration(t *testing.T){
	
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
		invalidSentence :="$GPRMC,123f19.50,A,4807.038,N,01131.000," + 
				   "E,022.4,084.4,210926,003.1,W*13\r\n"
		invalidBytes := []byte(invalidSentence)
		byteSentence := []byte(validSentence)
		for i := 0 ; i < 5 ; i++{
			if i == 2 {
				writer.Write(invalidBytes)
				continue
			}
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
			if len(gpxFile.Track.Segment.Points) != 4{
				t.Errorf("invalid length expected: %d got: %d",4,
						len(gpxFile.Track.Segment.Points))	
			}
			if math.Abs(expectedLatitude - latitude) > delta{
				t.Errorf("expected: %f got: %f",expectedLatitude,
								latitude)
			}
		case <- time.After(1 * time.Second):
			t.Error("Time exceeded.")
	}
	
}

func TestBadNMEASentencesInvalidParserIntegration(t *testing.T){

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
		invalidSentence :="$GPRMC,123519.50,V,4807.038,N,01131.000," + 
				   "E,022.4,084.4,210926,003.1,W*58\r\n"
		invalidBytes := []byte(invalidSentence)

		byteSentence := []byte(validSentence)
		for i := 0 ; i < 5 ; i++{
			if i == 2 {
				writer.Write(invalidBytes)
				continue
			}
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
			if len(gpxFile.Track.Segment.Points) != 4{
				t.Errorf("invalid length expected: %d got: %d",4,
						len(gpxFile.Track.Segment.Points))	
			}
			if math.Abs(expectedLatitude - latitude) > delta{
				t.Errorf("expected: %f got: %f",expectedLatitude,
								latitude)
			}
		case <- time.After(1 * time.Second):
			t.Error("Time exceeded.")
	}
}

func TestBadNMEASentencesBadCheckSumParserIntegration(t *testing.T){

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
		invalidSentence :="$GPRMC,123519.50,A,4807.038,N,01131.000," + 
				   "E,022.4,084.4,210926,003.1,W*58\r\n"
		invalidBytes := []byte(invalidSentence)

		byteSentence := []byte(validSentence)
		for i := 0 ; i < 5 ; i++{
			if i == 2 {
				writer.Write(invalidBytes)
				continue
			}
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
			if len(gpxFile.Track.Segment.Points) != 4{
				t.Errorf("invalid length expected: %d got: %d",4,
						len(gpxFile.Track.Segment.Points))	
			}
			if math.Abs(expectedLatitude - latitude) > delta{
				t.Errorf("expected: %f got: %f",expectedLatitude,
								latitude)
			}
		case <- time.After(1 * time.Second):
			t.Error("Time exceeded.")
	}
}
//TODO: Test Error case for bad NEMA sentences bad lat 
func TestBadNMEASentencesBadLatitudeParserIntegration(t *testing.T){

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
		invalidSentence :="$GPRMC,123519.50,A,4f07.038,N,01131.000," + 
				   "E,022.4,084.4,210926,003.1,W*1E\r\n"
		invalidBytes := []byte(invalidSentence)

		byteSentence := []byte(validSentence)
		for i := 0 ; i < 5 ; i++{
			if i == 2 {
				writer.Write(invalidBytes)
				continue
			}
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
			if len(gpxFile.Track.Segment.Points) != 4{
				t.Errorf("invalid length expected: %d got: %d",4,
						len(gpxFile.Track.Segment.Points))	
			}
			if math.Abs(expectedLatitude - latitude) > delta{
				t.Errorf("expected: %f got: %f",expectedLatitude,
								latitude)
			}
		case <- time.After(1 * time.Second):
			t.Error("Time exceeded.")
	}
}
//TODO: Test Error case for bad NEMA sentences bad long
//TODO: Test Error case for bad NEMA sentences blank fields
