package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	UserID       string `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *server) Login(c echo.Context) error {
	var req loginRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "Unprocessable Entity",
			"message": "Invalid json provided",
		})
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
