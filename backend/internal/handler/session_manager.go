package handler

import "syscall"
import "path/filepath"
import "errors"
import "os"
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
	_, err := sm.opener.Open(sm.deviceReadPath)
	if err != nil{
		
		if errors.Is(err,os.ErrPermission){
			return nil, NewSessionManagerError(int(syscall.EACCES),
				   "Permission Denied creating file: " + sm.deviceReadPath)
		}

	}
	joinedWriteFilePath := filepath.Join(sm.upperWritePath,filePath)
	_, err = sm.opener.Create(joinedWriteFilePath)
	if err != nil{
		if errors.Is(err,os.ErrPermission){
			return nil, NewSessionManagerError(int(syscall.EACCES),
				   "Permission Denied creating file: " + joinedWriteFilePath)
		}	
	}
	session := &models.Session{FileName:filePath,Id:id}
	return session,nil	
}
