package database

import (
	"server/pkg/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Conn *gorm.DB
	log  *logrus.Logger
}

// New connects to the Postgres Server and returns a pointer to the connection
func New(dsn string, logger *logrus.Logger) (*Postgres, error) {
	conn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	db := &Postgres{Conn: conn, log: logger}
	return db, nil
}

// Migrate will migrate all database Models
func (db *Postgres) Migrate() error {
	db.log.Info("Migrating Users...")
	err := db.Conn.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	db.log.Info("Migrating Channels...")
	err = db.Conn.AutoMigrate(&models.Channel{})
	if err != nil {
		return err
	}

	db.log.Info("Migrating Moderators...")
	err = db.Conn.AutoMigrate(&models.Moderator{})
	if err != nil {
		return err
	}
	return nil
}
