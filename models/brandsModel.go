package models

import (
	"github.com/jinzhu/gorm"
)

type Brands struct {
	gorm.Model
	Name string
	CompanyID uint
	IsDeleted int8
}