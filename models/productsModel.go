package models

type Products struct {
	ID int `gorm:"primary_key"`
	Name string
}