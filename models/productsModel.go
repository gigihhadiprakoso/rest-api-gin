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

// func (p *Products) AfterCreate(db *gorm.DB) (err error) {
// 	warehouse := Warehouses{
// 		Name: u.CompanyName,
// 		CompanyID: company.ID,
// 	}
// 	result = db.Create(&warehouse)
// 	return
// }