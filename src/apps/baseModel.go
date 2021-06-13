package apps

import (
	"time"
  "gorm.io/gorm"
)

//*Example define a base (for models)

type BaseModel struct {
	UUID      string         `gorm:"type:varchar(36);primaryKey;not null" json:"uuid"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}