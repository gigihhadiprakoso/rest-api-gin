package models

import (
	"github.com/jinzhu/gorm"
)

type PurchaseDetails struct {
	gorm.Model
	PurchaseID uint
	ProductID uint
	Qty int
	IsDeleted int8
}