package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

// Connect replace x & y wih mysql credentials
func Connect() {
	d, err := gorm.Open("mysql", "x:y/simplerest?charset=utf8&parseTime=True&location=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
