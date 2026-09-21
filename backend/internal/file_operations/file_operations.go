package file_operations  

import "io/fs"
import "os"
import "io"
import "strconv"
import "time"

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
	
	flags := os.O_CREATE | os.O_EXCL | os.O_WRONLY
	mode := os.FileMode(0666)
	file,err := os.OpenFile(path,flags,mode)
	if err != nil{	
		return nil, err
	}
	return file,nil
}
	
type GPSParser struct{

	layout string	
}

func NewGPSParser()*GPSParser{
		
	gp := &GPSParser{layout:"020106 150405.00"}
	return gp
	
}


func (gp *GPSParser) ParseSentence(sentence string){

	//TODO: implement me 
}


func (gp *GPSParser) parseTime(clockTime string,date string)(time.Time,error){

	concatenatedTime := date + " " + clockTime
	parsedTime, err := time.ParseInLocation(gp.layout,concatenatedTime,time.UTC)
	if err != nil{

		return parsedTime,err	
	}
	return parsedTime,nil 
}


//TODO: parse checksum

func (gp *GPSParser) parseCoordinate(hemisphere string,
				     coordinate string,
				     latitude bool)(float64,error){
	if coordinate == ""{
	
		return 0, NewGPSParserError(101,"empty coordinate field")	
	}		
	switch hemisphere {

		case "N","S","E","W":
			break 
		default:
			return 0.0,NewGPSParserError(101,"invalid hemisphere")
	}		
	if latitude {
			
		degrees := coordinate[:2]
		minutes := coordinate[2:]
		numericalDegree,err := strconv.ParseFloat(degrees,64)
		if err != nil{
			return 0,err
		}
		numericalMinutes,err := strconv.ParseFloat(minutes,64)
		if err != nil{
			return 0,err
		}
		result := numericalDegree + (numericalMinutes/60.0) 
		if hemisphere == "N"{
			return result,nil
		}
		if hemisphere == "S"{
			return result * -1.0,nil
		}
	}else if !latitude {
		
		degrees := coordinate[:3]
		minutes := coordinate[3:]
		numericalDegree, err := strconv.ParseFloat(degrees,64)
		if err != nil {
			return 0, err
		}		
		numericalMinutes,err := strconv.ParseFloat(minutes,64)
		if err != nil{

			return 0, err
		}
		result := numericalDegree + (numericalMinutes/60.0)
		if hemisphere == "E"{
			return result,nil	
		}	
		if hemisphere == "W"{
			return result * -1.0,nil
		}
	}
	return 0,NewGPSParserError(127,"unknown error unable to parse coordinate")		
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
 
