package models

import (
	"github.com/jinzhu/gorm"
)

type Warehouses struct {
	gorm.Model
	Name string
	CompanyID uint
	IsDeleted int8
}