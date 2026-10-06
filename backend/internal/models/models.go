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

type GPSFile struct{
	
	XMLName xml.Name `xml:"gpx"`
	Version string `xml:"version,attr"`
	Creator string `xml:"creator,attr"`
	Xmlns string `xml:"xmlns,attr"`
	Track GPSTrack `xml:"trk"`
}

type GPSTrack struct{
	Name string `xml:"name,omitempty"`
	Segment GPSTrackSegment `xml:"trkseg"`
}

type GPSTrackSegment struct{
	
	Points []GPSTrackPoint `xml:"trkpt"`
}

type GPSTrackPoint struct{
	XMLName xml.Name `xml:"trkpt"`
	Latitude float64 `xml:"lat,attr"`
	Longitude float64 `xml:"lon,attr"`
	Time time.Time `xml:"time,omitempty"`
}

