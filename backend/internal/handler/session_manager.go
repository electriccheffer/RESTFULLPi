package handler

import "restfulpi/internal/models"
import "restfulpi/internal/file_operations"

type SessionManagerService interface{

	StartSession(id string,filePath string) (*models.Session,error)
} 

type SessionManager struct{
	
	upperWritePath string
	deviceReadPath string
	sessions map[string]*models.Session
	opener file_operations.Opener
}

func NewSessionManager(wp string,rp string,op file_operations.Opener)*SessionManager{

	sm := &SessionManager{upperWritePath:wp,deviceReadPath:rp,opener:op}
	return sm
}

func (sm *SessionManager) StartSession(id string, filePath string)(*models.Session,error){

	_,exists := sm.sessions[id]
	if exists {
		
		// collision case throw error

	}	
		
	// check for file errors 
	
	session := &models.Session{FileName:filePath,Id:id}
	return session,nil	
}
