package utils

import (
	"github.com/twinj/uuid"
)

//*Return a string uuid
func GenerateUuid() string {
	uuid := uuid.NewV4()

	return uuid.String()
}