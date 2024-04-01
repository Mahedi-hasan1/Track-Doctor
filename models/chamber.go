package models

type Chamber struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address  string `json:"address"`
	District   string `json:"district"`
	Upzila string `json:"upzila"`
}