package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	gorm.Model
	FirstName string
	LastName string
	CompanyName string
	Username string
	Password string
	CompanyID uint
	IsDeleted int8
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *Users) AfterSave(db *gorm.DB) (err error) {
	company := Companies{
		Name: u.CompanyName,
	}
	result := db.Create(&company)
	if result.RowsAffected > 0 {
		var user []Users
		db.Model(&user).Where("id = ?", u.ID).Update("company_id", company.ID)
	}
	return
}