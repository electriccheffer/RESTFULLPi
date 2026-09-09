package file_operations  

import "io/fs"
import "os"
import "io"


type FileHandle interface {

	io.ReadWriteCloser	
}

type Opener interface{
	
	Open(path string)(FileHandle,error)
	Create(path string)(FileHandle,error)
}


type FileOpener struct{}

func NewFileOpener()Opener{
		
	fo := &FileOpener{}
	return fo
}


func (fo *FileOpener) Open(path string)(FileHandle,error){
	
	file,err := os.Open(path)
	if err != nil {

		return nil,err
	}
	return file,nil
	
}

func (fo *FileOpener) Create(path string)(FileHandle,error){

	file,err := os.Create(path)
	if err != nil{	
		return nil, err
	}
	return file,nil
}

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
 
