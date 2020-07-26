package models

type Categories struct {
	ID int `gorm:"primary_key"`
	Name string
}