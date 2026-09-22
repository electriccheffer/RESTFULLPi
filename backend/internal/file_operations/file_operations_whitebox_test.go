package file_operations 

import "testing"
import "math"
import "errors"
import "time"

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

func TestParseCoordinateErrorCaseEmptyCoordinate(t *testing.T){

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

func TestParseCoordinateErrorCaseEmptyHemisphere(t *testing.T){

	coordinate := "07103.5340"
	hemisphere := ""
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

func TestParseCoordinateErrorCaseInvalidHemisphere(t *testing.T){

	coordinate := "07103.5340"
	hemisphere := "B"
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

func TestParseTimeSuccess(t *testing.T){

	expected := time.Date(2026,time.September,21,12,1,30,300000000,time.UTC)
	day := "210926"
	dayTime := "120130.30"
	parser := NewGPSParser()
	result,err := parser.parseTime(dayTime,day)
	if err != nil {
		t.Errorf("Unexpected error:%s ", err.Error())
	}
	if !result.Equal(expected){

		t.Errorf("got:%s expected: %s",result,expected)
	}

}

func TestValidateChecksumSuccess(t *testing.T){

	validSentence := "$GPRMC,123519.50,A,4807.038,"+
			 "N,01131.000,E,022.4,084.4,210926,003.1,W*40"
	
	parser := NewGPSParser()
	got := parser.validateChecksum(validSentence)		
	if !got{
		t.Error("expected true returned false")
	}
}

func TestValidateChecksumFailure(t *testing.T){

	invalidSentence := "$GPRMC,23519.50,A,4807.038,"+
			 "N,01131.000,E,022.4,084.4,210926,003.1,W*40"
	
	parser := NewGPSParser()
	got := parser.validateChecksum(invalidSentence)		

	if got{
		t.Error("expected false returned false")
	}
}

func TestValidateChecksumNoStar(t *testing.T){

	invalidSentence := "$GPRMC,123519.50,A,4807.038,"+
			 "N,01131.000,E,022.4,084.4,210926,003.1,W40"
	
	parser := NewGPSParser()
	got := parser.validateChecksum(invalidSentence)		
	if got{
		t.Error("expected false returned true")
	}
}
