package main

import (
	"log"
	c "gin-api-boilerplate/src/config"
	"os"
	"gin-api-boilerplate/src/apps/examples"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&examples.Example{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}

	db := c.Init()
	Migrate(db)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := APIRouter()
	r.Run(":" + os.Getenv("PORT"))
}
