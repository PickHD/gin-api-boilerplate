package examples

import (
	"gorm.io/gorm"
)

type ExampleBase struct {
  DB *gorm.DB
}

func NewExampleBase(db *gorm.DB) *ExampleBase {
  return &ExampleBase{
    DB: db,
  }
}