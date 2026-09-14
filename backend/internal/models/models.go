package models

import "time"
import "encoding/xml"

type Status struct{
	Device string `json:"device"`
	Status string `json:"status"`
}

type Logs struct{
	Name string `json:"name"`
}

type Session struct{
	FileName string `json:"name"`
	Id string `json:"id"`
}

type GPSPathPoint struct{
	Latitude float64 `xml:"trkpt"`
	Logitude float64 `xml:"lat.attr"`
	Elevation float64 `xml:ele.omitempty`
	Time time.Time `xml:"time.omitempty"`
	Valid bool `xml:"-"`
}
