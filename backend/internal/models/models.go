package models

import "time"

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

type GPSTrackPoint struct{
	Latitude float64 `xml:"trkpt"`
	Longitude float64 `xml:"lat.attr"`
	Time time.Time `xml:"time.omitempty"`
}
