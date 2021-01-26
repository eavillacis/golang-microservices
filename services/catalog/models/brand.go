package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Brand struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
