package models

import (
	"github.com/jinzhu/gorm"
)

type Suppliers struct {
	gorm.Model
	Name string
	CompanyID uint
	IsDeleted int8
}