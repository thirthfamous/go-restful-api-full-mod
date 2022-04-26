package helper

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/belajar_golang_restful_api_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	PanicIfError(err)

	return db
}

func TruncateCategory(db *gorm.DB) {
	db.Exec("TRUNCATE category")
}
