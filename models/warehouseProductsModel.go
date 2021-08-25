package models

import (
	"github.com/jinzhu/gorm"
)

type WarehouseProducts struct {
	gorm.Model
	WarehouseID uint
	ProductID uint
	Stock int
	IsDeleted int8
}

func (WarehouseProducts) TableName() string {
	return "warehouse_products"
}
