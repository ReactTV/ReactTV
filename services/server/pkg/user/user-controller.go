package user

import (
	"server/pkg/models"
	"server/pkg/user/auth"
	"time"

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
func (c *UserController) Login(username string, password string) (*models.User, error) {
	foundUser, err := c.repo.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = auth.VerifyPassword(password, foundUser.Password)
	if err != nil {
		return nil, err
	}

	token, refreshToken, err := auth.GenerateAllTokens(foundUser.Username, foundUser.ID)
	if err != nil {
		return nil, err
	}

	updatedUser, err := c.repo.UpdateUserAuthByID(token, refreshToken, foundUser.ID)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (c *UserController) Signup(username string, password string) (*models.User, error) {
	var user models.User
	pwd, err := auth.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user.Username = username
	user.Password = pwd
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	u, err := c.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	token, refreshToken, err := auth.GenerateAllTokens(u.Username, u.ID)
	if err != nil {
		return nil, err
	}

	updatedUser, err := c.repo.UpdateUserAuthByID(token, refreshToken, u.ID)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}
