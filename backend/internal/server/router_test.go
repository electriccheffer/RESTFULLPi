package router_test

import "testing"
import "net/http"
import "net/http/httptest"
import "path/filepath"

import "restfulpi/internal/server"

func TestAlwaysTrue(t *testing.T) {

	if(true != true){
		t.Errorf("Always true test fails")
	}

}

func TestRouterCallsHandler(t *testing.T){
	testPath := filepath.Join("..","..","testData","browser","index.html");	
	router := router.NewRouter(testPath)
	request := httptest.NewRequest(http.MethodGet,"/",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)
	if response.Code != http.StatusOK {
		
		t.Errorf("status: %d want %d",response.Code,http.StatusOK)
	}
	
		
}

