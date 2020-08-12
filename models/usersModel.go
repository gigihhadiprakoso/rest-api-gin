package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Name string
	Username string
	Password string
	CompanyID uint
	IsDeleted int8
}