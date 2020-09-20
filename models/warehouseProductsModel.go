package models

import (
	"github.com/jinzhu/gorm"
)

type WarehouseProducts struct {
	gorm.Model
	Name string
	WarehouseID uint
	ProductID uint
	Stock int
	IsDeleted int8
}