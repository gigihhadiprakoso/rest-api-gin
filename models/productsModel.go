package models

import (
	"github.com/jinzhu/gorm"
)

type Products struct {
	gorm.Model
	Name string
	CategoryID uint
	BrandID uint
	CompanyID uint
	UnitID uint
	IsDeleted int8
}