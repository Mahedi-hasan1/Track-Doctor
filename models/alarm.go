package models

type Alarm struct {
	Id      int    `json:"id"`
	DoctorID    int `json:"doctorId"`
	UserId  string `json:"userId"`
	District   string `json:"district"`
	Upzila string `json:"upzila"`
	Status string `json:"status"`
}