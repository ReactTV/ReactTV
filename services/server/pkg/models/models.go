package models

import (
	"time"

	"github.com/google/uuid"
	gorm "gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate will set a ID to a V1 UUID
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	b.ID = uuid
	return
}

type User struct {
	Base
	Username     string
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Role         string
	Token        string
	RefreshToken string
}
