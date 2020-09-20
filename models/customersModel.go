package models

import (
	"github.com/jinzhu/gorm"
)

type Customers struct {
	gorm.Model
	Name string
	CompanyID uint
	IsDeleted int8
}