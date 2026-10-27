package router_test

import "testing"
import "net/http"
import "net/http/httptest"
import "path/filepath"
import "os"
import "log"
import "restfulpi/internal/server"
import "restfulpi/internal/file_operations"
import "restfulpi/internal/handler"

func TestAlwaysTrue(t *testing.T) {

	if(true != true){
		t.Errorf("Always true test fails")
	}

}

func TestRouterCallsHandler(t *testing.T){
	
	homeDir,err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v", err)
	}
	gpsLoggingPath := filepath.Join(homeDir, "gpsLogs")
	readPath := filepath.Join("/dev","ttyACM0")
	fileOpener := file_operations.NewFileOpener()	
	sessionManager := handler.NewSessionManager(gpsLoggingPath,readPath,fileOpener)
	router := router.NewRouter(sessionManager)
	request := httptest.NewRequest(http.MethodGet,"/",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	if response.Code != http.StatusOK {
		
		t.Errorf("status: %d want %d",response.Code,http.StatusOK)
	}
	
		
}

