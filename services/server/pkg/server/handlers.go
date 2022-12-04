package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginResponse struct {
	UserID       string `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *server) Login(c echo.Context) error {
	req := loginRequest{}
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := s.UserController.Login(req.Username, req.Password)
	if err != nil {
		return err
	}

	resp := &loginResponse{
		UserID:       user.ID.String(),
		Token:        user.Token,
		RefreshToken: user.RefreshToken,
	}

	return c.JSON(http.StatusOK, resp)
}

type signupRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type signupResponse struct {
	UserID       string `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *server) Signup(c echo.Context) error {
	req := signupRequest{}
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := s.UserController.Signup(req.Username, req.Password)
	if err != nil {
		return err
	}

	resp := &signupResponse{
		UserID:       user.ID.String(),
		Token:        user.Token,
		RefreshToken: user.RefreshToken,
	}

	return c.JSON(http.StatusOK, resp)
}
