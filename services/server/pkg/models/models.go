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

// BeforeCreate will set an ID to a V1 UUID and set time for updated/created dates
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	t := time.Now()
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	b.ID = uuid

	b.CreatedAt = t
	b.UpdatedAt = t
	return
}

type User struct {
	Base
	Username     string `gorm:"unique;not null"`
	Password     string
	Token        string
	RefreshToken string
}

type Channel struct {
	Base
	Name   string    `gorm:"unique;not null"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
}

type Moderator struct {
	Base
	Role      string
	ChannelID uuid.UUID `gorm:"type:uuid;not null"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
}
