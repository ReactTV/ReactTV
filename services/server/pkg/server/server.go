package server

import (
	"fmt"
	"server/pkg/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type server struct {
	Log            *logrus.Logger
	Port           string
	UserController *user.UserController
}

func NewServer(l *logrus.Logger, port string, db *gorm.DB) *server {
	return &server{
		Log:            l,
		Port:           port,
		UserController: user.NewUserController(db, l),
	}
}

func (s *server) Serve() error {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Setup custom validator for performing data validation on request payload
	e.Validator = &CustomValidator{}

	// Middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			s.Log.WithFields(logrus.Fields{
				"URI":     values.URI,
				"status":  values.Status,
				"latency": values.Latency.Milliseconds(),
				"agent":   values.UserAgent,
			}).Info("request")

			return nil
		},
	}))

	// Initialize Routes
	s.InitializeUserRoutes(e)

	// Starting Server
	s.Log.WithField("port", s.Port).Info("Starting HTTP Server...")
	err := e.Start(fmt.Sprintf(":%s", s.Port))
	if err != nil {
		return err
	}
	return nil
}
