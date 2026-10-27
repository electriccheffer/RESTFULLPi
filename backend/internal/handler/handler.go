package handler

import "net/http"
import "encoding/json"
import "io/fs"
import "time"
import "crypto/rand"
import "encoding/hex"
import "errors"
import "restfulpi/internal/models"


type FrontendHandler struct{

	fileServer http.Handler
	fileSystem fs.FS		

} 

func NewFrontendHandler(fileSystem fs.FS) *FrontendHandler {
	
	fileServer := http.FileServer(http.FS(fileSystem))
	frontHandler := &FrontendHandler{fileServer,fileSystem}
	return frontHandler 
	
}

func (fh *FrontendHandler) ServeHTTP(response http.ResponseWriter,request *http.Request) {
	
	path := request.URL.Path
	if path != "/"{
		
		if file,err := fh.fileSystem.Open(path[1:]) ; err != nil{
	
			request.URL.Path = "/"
		} else{
				
			stat, err := file.Stat()
			if err != nil || stat.IsDir() {
				request.URL.Path = "/"
			}
			_ = file.Close()  
		}
	}	
	
	fh.fileServer.ServeHTTP(response,request)
}

type StatusHandler struct{}

func NewStatusHandler() *StatusHandler{
	
	sh := &StatusHandler{}
	return sh
	
}

func (sh *StatusHandler) ServeHTTP(response http.ResponseWriter,request *http.Request){
	
	status := models.Status{Device:"RESTFULPi",Status:"Online"}
	encodedStatus,err := json.Marshal(status)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		http.Error(response,err.Error(),http.StatusInternalServerError)
		return
	}		
	response.WriteHeader(http.StatusOK)
	response.Write(encodedStatus)
		
	return

}

type GetLogsHandler struct{
	path string
	sessionManager SessionManagerService
}

func NewGetLogsHandler(path string,sessionManager SessionManagerService) *GetLogsHandler{
	lh := &GetLogsHandler{path:path,sessionManager:sessionManager}
	return lh 
}

func (glh *GetLogsHandler) ServeHTTP(response http.ResponseWriter,request *http.Request){
		
	sessions := glh.sessionManager.GetSessions() 
	var logs []models.Logs	
	for _, session := range sessions {
		log := models.Logs{Name:session.FileName,Id:session.Id}
		logs = append(logs,log)
	}
	
	jsonLogs,err := json.Marshal(logs)
	if err != nil{
				
	}
	response.WriteHeader(http.StatusOK)
	response.Write(jsonLogs)		
	return

	
}

type SessionStartHandler struct{
	
	manager SessionManagerService
	randomReader func(b []byte) (int, error)
	timeMaker func()(time.Time)	
}

func NewSessionStartHandler(mgr SessionManagerService) *SessionStartHandler{
	
	ssh := &SessionStartHandler{manager:mgr,randomReader:rand.Read,timeMaker:time.Now}
	return ssh

}

func (ssh *SessionStartHandler) ServeHTTP(response http.ResponseWriter,request *http.Request){

	now := ssh.timeMaker() 
	filePath := now.Format("01_02_06_15_04_05.gpx")			

	randomBytes := make([]byte,16)
	 _,err := ssh.randomReader(randomBytes)

	if err != nil{
			
	}
	id := hex.EncodeToString(randomBytes)
	
	session, err := ssh.manager.StartSession(id,filePath)
		
	if err != nil {
		
		retries := 3
		var sessionManagerError *SessionManagerError
		errors.As(err,&sessionManagerError)
		if sessionManagerError.Code == 2{

			response.Header().Set("Content-Type","application/json")
			response.Header().Set("X-Content-Type-Options","nosniff")
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
		if sessionManagerError.Code == 13 {
			
			response.Header().Set("Content-Type","application/json")
			response.Header().Set("X-Content-Type-Options","nosniff")
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
		if sessionManagerError.Code == 17{

			for attempt := 1 ; attempt < retries && err !=nil; attempt++{
			
				now = now.Add(1 * time.Second)
				filePath = now.Format("01_02_06_15_04_05.gpx")
				session, err = ssh.manager.StartSession(id,filePath)	
			}
			if err == nil{
				jsonSession, err := json.Marshal(session)
				if err != nil{


				}
	
				response.Header().Set("Content-Type","application/json")
				response.Header().Set("X-Content-Type-Options","nosniff")
				response.WriteHeader(http.StatusCreated)
				response.Write(jsonSession)
				return
	 
			}else{
				response.Header().Set("Content-Type","application/json")
				response.Header().Set("X-Content-Type-Options","nosniff")
				response.WriteHeader(http.StatusConflict)			
				return
			}
		}
		if sessionManagerError.Code == 409{
			retries = 3	
			for attempt := 1 ; attempt < retries && err != nil ; attempt++{
				randomBytes = make([]byte,16)
				_,err = ssh.randomReader(randomBytes)
				id = hex.EncodeToString(randomBytes)
				session, err = ssh.manager.StartSession(id,filePath)	
			}
			if err == nil{
				
				jsonSession, err := json.Marshal(session)
				if err != nil{

				}
	
				response.Header().Set("Content-Type","application/json")
				response.Header().Set("X-Content-Type-Options","nosniff")
				response.WriteHeader(http.StatusCreated)
				response.Write(jsonSession)
	 			return
			}else{
				response.Header().Set("Content-Type","application/json")
				response.Header().Set("X-Content-Type-Options","nosniff")
				response.WriteHeader(http.StatusConflict)			
				return
			}
				
		}			
	
	}

	jsonSession, err := json.Marshal(session)
	if err != nil{


	}
	
	response.Header().Set("Content-Type","application/json")
	response.Header().Set("X-Content-Type-Options","nosniff")
	response.WriteHeader(http.StatusCreated)
	response.Write(jsonSession)
	return
}
