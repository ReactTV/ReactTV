package user

import (
	"server/pkg/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserController struct {
	repo *UserRepository
	log  *logrus.Logger
}

func NewUserController(db *gorm.DB, l *logrus.Logger) *UserController {
	return &UserController{repo: NewUserRepository(db, l), log: l}
}

// Login controller for username and password authentication
func (c *UserController) Login(username string, password string) (models.User, error) {

	var user models.User
	return user, nil
}
