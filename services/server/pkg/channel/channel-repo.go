package channel

import (
	"errors"
	"server/pkg/database"
	"server/pkg/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ChannelRepository struct {
	db  *database.Postgres
	log *logrus.Logger
}

func NewChannelRepository(db *database.Postgres, l *logrus.Logger) *ChannelRepository {
	return &ChannelRepository{db: db, log: l}
}

func (u *ChannelRepository) FindChannelByID(id uuid.UUID) (models.Channel, error) {
	var chann models.Channel
	result := u.db.Conn.First(&chann, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return chann, gorm.ErrRecordNotFound
		} else {
			return chann, errors.New("error occurred while fetching user")
		}
	}
	return chann, nil
}

// CreateChannel creates a new channel associated to a UserID
func (cr *ChannelRepository) CreateChannel(user_id uuid.UUID, name string) (models.Channel, error) {
	var ch models.Channel
	ch.UserID = user_id
	ch.Name = name
	result := cr.db.Conn.FirstOrCreate(&ch)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ch, gorm.ErrRecordNotFound
		} else {
			return ch, errors.New("error occurred while fetching channel")
		}
	}

	// Channel created successfully
	if result.RowsAffected == 1 {
		return ch, nil
	}

	// Channel already exists
	return ch, errors.New("channel already exists")
}
