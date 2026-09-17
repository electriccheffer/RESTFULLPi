package file_operations 

import "testing"
import "math"
import "errors"

func TestParseCoordinateLatitudeNorth(t *testing.T){

	coordinate := "9012.3456"
	expected := 90.205760
	hemisphere := "N"	
	latitude := true
	gpsParser := NewGPSParser()
	delta := 0.000001	
	result,err := gpsParser.parseCoordinate(hemisphere,coordinate,latitude)
	if err != nil {

		t.Errorf("Unexpected error: %s",err.Error())
	}
	if math.Abs(result -  expected) > delta{
	
		t.Errorf("Incorrect coordinate returned got:%f expected:%f",result,expected)
	}
	
}

func TestParseCoordinateLatitudeSouth(t *testing.T){

	coordinate := "9012.3456"
	expected := -90.205760
	hemisphere := "S"	
	latitude := true
	gpsParser := NewGPSParser()
	delta := 0.000001	
	result,err := gpsParser.parseCoordinate(hemisphere,coordinate,latitude)
	if err != nil {

		t.Errorf("Unexpected error: %s",err.Error())
	}
	if math.Abs(result -  expected) > delta{
	
		t.Errorf("Incorrect coordinate returned got:%f expected:%f",result,expected)
	}
	
}

func TestParseCoordinateLongitudeEast(t *testing.T){

	coordinate := "07103.5340"
	expected := 71.058900
	hemisphere := "E"
	latitude := false
	
	gpsParser := NewGPSParser()
	delta := 0.000001
	result, err := gpsParser.parseCoordinate(hemisphere,coordinate,latitude)
	if err != nil {
		t.Errorf("Unexpected error: %s",err.Error())
	}
	if math.Abs(result - expected) > delta{
		t.Errorf("Incorrect coordinate returned got:%f expected:%f",result,expected)
	}

}

func TestParseCoordinateLongitudeWest(t *testing.T){

	coordinate := "07103.5340"
	expected := -71.058900
	hemisphere := "W"
	latitude := false
	
	gpsParser := NewGPSParser()
	delta := 0.000001
	result, err := gpsParser.parseCoordinate(hemisphere,coordinate,latitude)
	if err != nil {
		t.Errorf("Unexpected error: %s",err.Error())
	}
	if math.Abs(result - expected) > delta{
		t.Errorf("Incorrect coordinate returned got:%f expected:%f",result,expected)
	}
}

func TestParseCoordinteErrorCaseEmptyCoordinate(t *testing.T){

	coordinate := ""
	hemisphere := "W"
	latitude := false
	
	gpsParser := NewGPSParser()
	_, err := gpsParser.parseCoordinate(hemisphere,coordinate,latitude)
	if err == nil{
		t.Error("Error expected none thrown")
	}
	if err != nil{
		var parserErr *GPSParserError
		if !errors.As(err,&parserErr){
			t.Errorf("Error of incorrect type:%s",err.Error())
		} else{
			if parserErr.Code != 101{
				t.Errorf("GPSParserError wrong code expected:%d got:%d",
					101,parserErr.Code)
			}
		}
	}

}
