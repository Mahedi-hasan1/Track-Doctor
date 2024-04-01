package models

type Schedule struct {
	DoctorId      int    `json:"DoctorId"`
	ChamberId    string `json:"chamberId"`
	Fee  string `json:"fee"`
	Status   string `json:"status"`
	Days string `json:"days"`
	//pic,
}