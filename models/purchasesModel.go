package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Purchases struct {
	gorm.Model
	ReferenceNo string
	Date time.Time
	WarehouseID uint
	SupplierID uint
	CustomerID uint
	CompanyID uint
	IsDeleted int8
}