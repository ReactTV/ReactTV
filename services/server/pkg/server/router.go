package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	// Users
	usersBaseRoute = "/users"
	loginRoute     = "/login"
	signupRoute    = "/signup"
)

func (s *server) InitializeUserRoutes(e *echo.Echo) {
	// Users Routes
	users := e.Group(usersBaseRoute)

	s.Log.WithFields(logrus.Fields{"Method": http.MethodPost, "Route": fmt.Sprint(usersBaseRoute, loginRoute)}).Info("Added HTTP Route")
	users.Add(http.MethodPost, loginRoute, s.Login)

	s.Log.WithFields(logrus.Fields{"Method": http.MethodPost, "Route": fmt.Sprint(usersBaseRoute, signupRoute)}).Info("Added HTTP Route")
	users.Add(http.MethodPost, signupRoute, s.Signup)
}
