package models

import (
	"github.com/google/uuid"
)

type Base struct {
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"id"`
}
