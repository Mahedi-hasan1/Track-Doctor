package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	District string `json:"district"`
	Upzila   string `json:"upzila"`
}
