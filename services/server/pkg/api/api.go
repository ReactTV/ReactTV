package api

import (
	"fmt"
	"server/pkg/api/router"
	"server/pkg/api/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type server struct {
	port   string
	log    *logrus.Logger
	router *router.Router
}

func NewServer(l *logrus.Logger, port string, r *router.Router) *server {
	return &server{
		log:    l,
		port:   port,
		router: r,
	}
}

func (s *server) Serve() error {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Setup custom validator for performing data validation on request payload
	e.Validator = &utils.CustomValidator{}

	// Middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			s.log.WithFields(logrus.Fields{
				"URI":     values.URI,
				"status":  values.Status,
				"latency": values.Latency.Milliseconds(),
				"agent":   values.UserAgent,
			}).Info("request")

			return nil
		},
	}))

	// Initialize Routes
	s.router.InitializeSwaggerRoutes(e)
	s.router.InitializeUserRoutes(e)
	s.router.InitializeChannelRoutes(e)

	// Starting Server
	s.log.WithField("port", s.port).Info("Starting HTTP Server...")
	err := e.Start(fmt.Sprintf(":%s", s.port))
	if err != nil {
		return err
	}
	return nil
}
