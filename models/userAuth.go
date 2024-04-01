package models

type UserAuth struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//pic,
}
