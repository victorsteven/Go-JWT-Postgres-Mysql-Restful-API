package services

import "github.com/jinzhu/gorm"

type DbInstance struct {
	DB *gorm.DB
}

func NewDbInstance(DB *gorm.DB) *DbInstance {
	return &DbInstance{DB}
}
