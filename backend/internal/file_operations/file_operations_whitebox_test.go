package file_operations 

import "testing"
import "math"

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

//TODO: logitude case
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

//TODO: error cases
