package models

import (
	"github.com/jinzhu/gorm"
)

type Companies struct {
	gorm.Model
	Name string
	IsDeleted int8
}