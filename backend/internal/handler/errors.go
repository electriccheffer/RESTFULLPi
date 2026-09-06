package handler 

import "fmt"

type SessionManagerError struct{
	Code int
	Message string
}

func (e *SessionManagerError) Error()string{
	return fmt.Sprintf("SessionManagerError %d: %s",e.Code,e.Message)
}

func NewSessionManagerError(c int, m string) *SessionManagerError {

	return &SessionManagerError{Code:c,Message:m}
}
