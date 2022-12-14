package commands

import (
	"fmt"
	"server/pkg/api"
	"server/pkg/api/router"
	"server/pkg/channel"
	"server/pkg/database"
	"server/pkg/user"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Port          string `required help:"Port the server listens on" default:"8000"`
	DBHost        string `required env:"DB_HOST" help:"Postgres hostname or IP address"`
	DBUser        string `required env:"DB_USER" help:"Postgres username"`
	DBPassword    string `required env:"DB_PASSWORD" help"Postgres password"`
	DBName        string `required env:"DB_NAME" help"Postgres database name"`
	DBDefaultName string `required env:"DB_NAME_POSTGRES" help"The default Postgres database name (usually 'postgres')"`
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

	// Create database (reacttv) if it doesn't exist
	default_dsn := fmt.Sprintf("postgresql://%v:%v@%v/%v", flag.DBUser, flag.DBPassword, flag.DBHost, flag.DBDefaultName)
	err = database.Create(default_dsn, flag.DBName, logger)
	if err != nil {
		return err
	}

	// Connect to racttv database
	reacttv_dsn := fmt.Sprintf("postgresql://%v:%v@%v/%v", flag.DBUser, flag.DBPassword, flag.DBHost, flag.DBName)
	db, err := database.New(reacttv_dsn, logger)
	if err != nil {
		return err
	}

	// Migrate tables in reacttv database
	err = db.Migrate()
	if err != nil {
		return err
	}

	// Controllers
	uc := user.NewUserController(db, logger)
	cc := channel.NewChannelController(db, logger)

	// Router / Handlers
	r := router.NewRouter(logger, uc, cc)

	// API Server
	s := api.NewServer(logger, flag.Port, r)

	err = s.Serve()
	if err != nil {
		return err
	}
	return nil
}
