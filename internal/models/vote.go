package models

type Vote struct {
	Uuid   string `json:"uuid"`
	Rating int8   `json:"rating"`
}
