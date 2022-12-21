package database

import (
	"fmt"
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
func New(reacttv_dsn string, logger *logrus.Logger) (*Postgres, error) {
	conn, err := gorm.Open(postgres.Open(reacttv_dsn))
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

// Create a Database (requires LOGIN grants on user)
func Create(admin_dsn string, database_name string, log *logrus.Logger) error {
	db, err := gorm.Open(postgres.Open(admin_dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Check if DB Exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", database_name)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		return rs.Error
	}

	// Create DB
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		log.Infof("Creating Database %s...", database_name)
		stmt := fmt.Sprintf("CREATE DATABASE %s;", database_name)
		if rs := db.Exec(stmt); rs.Error != nil {
			return rs.Error
		}

		// Close DB Connection
		sql, err := db.DB()
		defer func() {
			_ = sql.Close()
		}()
		if err != nil {
			return err
		}
	}
	return nil
}
