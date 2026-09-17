package file_operations

import "fmt"

type GPSParserError struct{
	Code int
	Message string
}

func (e *GPSParserError) Error()string{

	return fmt.Sprintf("GPSParserError %d: %s",e.Code,e.Message)
}

func NewGPSParserError(c int, m string) *GPSParserError{

	return &GPSParserError{Code:c,Message:m}
}
