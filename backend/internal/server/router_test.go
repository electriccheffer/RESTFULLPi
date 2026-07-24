package router 

import "testing"

func TestAlwaysTrue(t *testing.T) {

	if(true != true){
		t.Errorf("Always true test fails")
	}

}
