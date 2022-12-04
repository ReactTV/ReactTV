package commands

import (
	"fmt"
	"server/pkg/server"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DBHost     string `required env:"DB_HOST" help:"Postgres hostname or IP address"`
	DBUser     string `required env:"DB_USER" help:"Postgres username"`
	DBPassword string `required env:"DB_PASSWORD" help"Postgres password"`
	DBName     string `required env:"DB_NAME" help"Postgres database name"`
	Port       string `required help:"Port the server listens on" default:"8000"`
}

func (flag *Server) Run(globals *Globals) error {
	// Setup logger
	logger := logrus.New()
	lvl, err := logrus.ParseLevel(globals.LogLevel)
	if err != nil {
		logger.Fatal(err)
	}
	logger.SetLevel(lvl)

	switch globals.LogFormat {
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{})
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	}

	dburl := fmt.Sprintf("postgresql://%v:%v@%v/%v", flag.DBUser, flag.DBPassword, flag.DBHost, flag.DBName)
	db, err := gorm.Open(postgres.Open(dburl))
	if err != nil {
		return err
	}

	s := server.NewServer(logger, flag.Port, db)
	err = s.Serve()
	if err != nil {
		return err
	}
	return nil
}
