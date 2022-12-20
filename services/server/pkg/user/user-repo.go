package user

import (
	"errors"
	"server/pkg/database"
	"server/pkg/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db  *database.Postgres
	log *logrus.Logger
}

func NewUserRepository(db *database.Postgres, l *logrus.Logger) *UserRepository {
	return &UserRepository{db: db, log: l}
}

func (u *UserRepository) FindUserByUsername(username string) (models.User, error) {
	var user models.User
	result := u.db.Conn.First(&user, "username = ?", username)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, gorm.ErrRecordNotFound
		} else {
			return user, errors.New("error occurred while fetching user")
		}
	}
	return user, nil
}

func (u *UserRepository) UpdateUserAuthByID(token string, refreshToken string, id uuid.UUID) (models.User, error) {
	var user models.User
	result := u.db.Conn.First(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, gorm.ErrRecordNotFound
		} else {
			return user, errors.New("error occurred while fetching user")
		}
	}

	user.Token = token
	user.RefreshToken = refreshToken

	tx := u.db.Conn.Save(&user)
	if tx.RowsAffected != 1 {
		return user, errors.New("error occurred while updating user")
	}
	return user, nil
}

func (ur *UserRepository) CreateUser(user models.User) (models.User, error) {
	result := ur.db.Conn.FirstOrCreate(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, gorm.ErrRecordNotFound
		} else {
			return user, errors.New("error occurred while fetching user")
		}
	}

	// User created successfully
	if result.RowsAffected == 1 {
		return user, nil
	}

	// user already exists
	return user, errors.New("user already exists")
}
