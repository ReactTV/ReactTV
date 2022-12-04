package user

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewUserRepository(db *gorm.DB, l *logrus.Logger) *UserRepository {
	return &UserRepository{db: db, log: l}
}
