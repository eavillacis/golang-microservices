package types

import (
	"github.com/gofrs/uuid"
)

// CreateBrandConfig ...
type CreateBrandConfig struct {
	Name        string
	ProgramCode uuid.UUID
}
