package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	// Users
	loginBaseRoute = "/users"
	loginRoute     = "/login"
)

func (s *server) InitializeRoutes(e *echo.Echo) {
	// Users Routes
	users := e.Group(loginBaseRoute)

	s.Log.WithFields(logrus.Fields{"Method": http.MethodPost, "Route": fmt.Sprint(loginBaseRoute, loginRoute)}).Info("Added HTTP Route")
	users.Add(http.MethodPost, loginRoute, s.Login)
}
