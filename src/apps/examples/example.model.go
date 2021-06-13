package examples

import (
	"gin-api-boilerplate/src/apps"
	"gin-api-boilerplate/src/utils"

	"gorm.io/gorm"
)

//*Create models,hooks here

type Example struct {
	apps.BaseModel
	Name string `gorm:"varchar(255);not null"`
}

func (e *Example) BeforeCreate(tx *gorm.DB) (err error) {
	e.UUID=utils.GenerateUuid()

	return nil
}