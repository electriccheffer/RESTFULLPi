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

	sm := &SessionManager{upperWritePath:wp,
				deviceReadPath:rp,opener:op,
				sessions:make(map[string]*models.Session)}
	return sm
}

func (sm *SessionManager) StartSession(id string, filePath string)(*models.Session,error){

	_,exists := sm.sessions[id]
	if exists {
		
		return nil, NewSessionManagerError(409,"Id in session table already exists")

	}	
		
	// check for file errors 
	_, err := sm.opener.Open(sm.deviceReadPath)
	if err != nil{
		
		if errors.Is(err,os.ErrPermission){
			return nil, NewSessionManagerError(int(syscall.EACCES),
				   "Permission Denied opening serial file:" + sm.deviceReadPath)
		}
		if errors.Is(err,os.ErrNotExist){
			return nil, NewSessionManagerError(int(syscall.ENOENT),
				    "Serial file does not exist:: " + sm.deviceReadPath)
		}

	}
	joinedWriteFilePath := filepath.Join(sm.upperWritePath,filePath)
	_, err = sm.opener.Create(joinedWriteFilePath)
	if err != nil{
		if errors.Is(err,os.ErrPermission){
			return nil, NewSessionManagerError(int(syscall.EACCES),
				   "Permission Denied creating file: " + joinedWriteFilePath)
		}	
		if errors.Is(err,os.ErrNotExist){
			return nil, NewSessionManagerError(int(syscall.ENOENT),
				    "Directory does not exist: " + joinedWriteFilePath)
		}
		if errors.Is(err,os.ErrExist){

			return nil, NewSessionManagerError(int(syscall.EEXIST),
				    "Write file already exists")
		}
	}

	session := &models.Session{FileName:filePath,Id:id}
	sm.sessions[id] = session
	return session,nil	
}
