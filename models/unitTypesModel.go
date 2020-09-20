package models

import (
	"github.com/jinzhu/gorm"
)

type UnitTypes struct {
	gorm.Model
	Name string
	CompanyID uint
	IsDeleted int8
}