package models

type Actress struct {
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
	Photo  string `json:"photo"`
}
