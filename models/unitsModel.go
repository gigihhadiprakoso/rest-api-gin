package models

import (
	"github.com/jinzhu/gorm"
)

type Units struct {
	gorm.Model
	Name string
	Symbol string
	UnitTypeID uint
	CompanyID uint
	IsDeleted int8
}

