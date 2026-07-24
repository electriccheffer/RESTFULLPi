package router 

import "testing"
import "net/http"
import "net/http/httptest"

func TestAlwaysTrue(t *testing.T) {

	if(true != true){
		t.Errorf("Always true test fails")
	}

}

func TestRouterCallsHandler(t *testing.T){

	
	handler := http.HandlerFunc(
			func(writer http.ResponseWriter,request *http.Request){

				writer.WriteHeader(http.StatusOK)
				writer.Write([]byte("Not the app"))	

			})
	
	router := NewRouter(handler)
	request := httptest.NewRequest(http.MethodGet,"/",nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)

	if response.Code != http.StatusOK {
		
		t.Errorf("status: %d want %d",response.Code,http.StatusOK)
	}
	
		
}

