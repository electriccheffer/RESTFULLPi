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


