package router

import (
	"fmt"
	"net/http"
	channel_handler "server/pkg/api/router/handlers/channel"
	user_handler "server/pkg/api/router/handlers/user"
	"server/pkg/channel"
	"server/pkg/user"

	_ "server/pkg/api/docs"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	log            *logrus.Logger
	userHandler    *user_handler.Handler
	channelHandler *channel_handler.Handler
}

func NewRouter(l *logrus.Logger, uc *user.UserController, cc *channel.ChannelController) *Router {
	return &Router{
		log:            l,
		userHandler:    user_handler.NewHandler(uc),
		channelHandler: channel_handler.NewHandler(cc),
	}
}

const (
	// Swagger (api docs)
	swaggerBaseRoute = "/docs"
	swaggerAllRoute  = "/*"

	// User
	userBaseRoute = "/user"
	loginRoute    = "/login"
	signupRoute   = "/signup"

	// Channel
	channelBaseRoute   = "/channel"
	createChannelRoute = "/create"
)

func (r *Router) InitializeUserRoutes(e *echo.Echo) {
	user := e.Group(userBaseRoute)

	r.log.WithFields(logrus.Fields{"Method": http.MethodPost, "Route": fmt.Sprint(userBaseRoute, loginRoute)}).Info("Added HTTP Route")
	user.Add(http.MethodPost, loginRoute, r.userHandler.Login)

	r.log.WithFields(logrus.Fields{"Method": http.MethodPost, "Route": fmt.Sprint(userBaseRoute, signupRoute)}).Info("Added HTTP Route")
	user.Add(http.MethodPost, signupRoute, r.userHandler.Signup)
}

func (r *Router) InitializeChannelRoutes(e *echo.Echo) {
	channel := e.Group(channelBaseRoute)

	r.log.WithFields(logrus.Fields{"Method": http.MethodPost, "Route": fmt.Sprint(channelBaseRoute, createChannelRoute)}).Info("Added HTTP Route")
	channel.Add(http.MethodPost, createChannelRoute, r.channelHandler.CreateChannel)
}

func (r *Router) InitializeSwaggerRoutes(e *echo.Echo) {
	channel := e.Group(swaggerBaseRoute)

	r.log.WithFields(logrus.Fields{"Method": http.MethodGet, "Route": fmt.Sprint(swaggerBaseRoute, swaggerAllRoute)}).Info("Added HTTP Route")
	channel.Add(http.MethodGet, swaggerAllRoute, echoSwagger.WrapHandler)
}
