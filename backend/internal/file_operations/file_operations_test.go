package file_operations_test

import  "path/filepath"
import  "restfulpi/internal/file_operations"
import "testing"

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

