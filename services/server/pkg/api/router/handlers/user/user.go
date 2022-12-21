package user

import (
	"net/http"
	"server/pkg/user"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	user *user.UserController
}

func NewHandler(uc *user.UserController) *Handler {
	return &Handler{user: uc}
}

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginResponse struct {
	UserID       string `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// Login godoc
//
//	@Summary		Logs a user in
//	@Accept			json
//	@Produce		json
//	@Param			request	body loginRequest true "Login Request"
//	@Success		200	{object}	loginResponse
//	@Router			/user/login [post]
func (a *Handler) Login(c echo.Context) error {
	req := loginRequest{}
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := a.user.Login(req.Username, req.Password)
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

// Signup godoc
//
//	@Summary		Create a user
//	@Accept			json
//	@Produce		json
//	@Param			request	body signupRequest true "Signup Request"
//	@Success		200	{object} signupResponse
//	@Router			/user/signup [post]
func (a *Handler) Signup(c echo.Context) error {
	req := signupRequest{}
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := a.user.Signup(req.Username, req.Password)
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
