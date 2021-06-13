package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

//*Initialize database
func Init() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s TimeZone=Asia/Shanghai", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("db error: (Init) ", err)
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)

	DB = db
	return DB
}

//*Return instance of database
func GetDB() *gorm.DB {
	return DB
}
