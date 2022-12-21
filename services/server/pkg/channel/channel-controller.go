package channel

import (
	"server/pkg/database"
	"server/pkg/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ChannelController struct {
	repo *ChannelRepository
	log  *logrus.Logger
}

func NewChannelController(db *database.Postgres, l *logrus.Logger) *ChannelController {
	return &ChannelController{repo: NewChannelRepository(db, l), log: l}
}

func (c *ChannelController) CreateChannel(user_id uuid.UUID, name string) (*models.Channel, error) {
	channel, err := c.repo.CreateChannel(user_id, name)
	if err != nil {
		return nil, err
	}

	return &channel, nil
}

func (c *ChannelController) GetChannel(channel_id uuid.UUID) (*models.Channel, error) {
	channel, err := c.repo.FindChannelByID(channel_id)
	if err != nil {
		return nil, err
	}

	return &channel, nil
}
