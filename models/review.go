package models

type Review struct {
	Id       int    `json:"id"`
	Body     string `json:"body"`
	UserId   int    `json:"userId"`
	DoctorId int    `json:"doctorId"`
}
