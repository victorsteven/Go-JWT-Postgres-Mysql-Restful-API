package database

import (
	"github.com/jinzhu/gorm"
	"github.com/victorsteven/fullstack/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
