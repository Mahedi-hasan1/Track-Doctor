package models

type Doctor struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Digree  string `json:"digree"`
	Specs   string `json:"specs"`
	Picture string `json:"picture"`
}
