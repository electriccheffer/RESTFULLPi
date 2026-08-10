package models

type Status struct{
	
	Device string `json:"device"`
	Status string `json:"status"`
}

type Logs struct{
	Name string `json:"name"`
}
