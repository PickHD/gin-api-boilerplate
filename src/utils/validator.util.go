package utils

import (
	"github.com/twinj/uuid"
)

func ValidateUuid(data string) bool {
	_, err := uuid.Parse(data)

	return err == nil
}