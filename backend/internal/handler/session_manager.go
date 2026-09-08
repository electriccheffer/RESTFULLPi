package handler

import "restfulpi/internal/models"

type SessionManagerService interface{

	StartSession(id string,filePath string) (*models.Session,error)
} 

type SessionManager struct{
	
	upperWritePath string
	deviceReadPath string
	sessions map[string]*models.Session
}

func NewSessionManager(wp string,rp string)*SessionManager{

	sm := &SessionManager{upperWritePath:wp,deviceReadPath:rp}
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
