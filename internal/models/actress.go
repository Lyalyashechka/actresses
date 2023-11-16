package models

type Actress struct {
	Uuid   string `json:"uuid" gorm:"column:uuid"`
	Name   string `json:"name" gorm:"column:name"`
	Rating int    `json:"rating" gorm:"column:rating"`
	Photo  string `json:"photo" gorm:"column:photo"`
}

func (Actress) TableName() string {
	return "actresses"
}
