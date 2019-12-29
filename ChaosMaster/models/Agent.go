package models

type Agent struct {
	Id      string `json:"id"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Enabled bool   `json:"enabled"`
	Status  string `json:"status"`
}
