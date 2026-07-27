package file_operations  

import "io/fs"
import "os"

type DirectoryRead struct{
	
	path string	

}

func NewDirectoryRead(path string) *DirectoryRead{

	dr := &DirectoryRead{path}
	return dr 
}

func (dr *DirectoryRead) GetFiles() ([]os.DirEntry,error){
	
	fileSystem := os.DirFS(dr.path)		
	entries, err := fs.ReadDir(fileSystem,".")
	return entries,err

}

 
