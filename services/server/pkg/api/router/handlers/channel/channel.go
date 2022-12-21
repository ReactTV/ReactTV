package channel

import (
	"net/http"
	"server/pkg/channel"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	channel *channel.ChannelController
}

func NewHandler(c *channel.ChannelController) *Handler {
	return &Handler{channel: c}
}

type createChannelRequest struct {
	ID   string `json:"user_id" validate:"required"`
	Name string `json:"channel_name" validate:"required"`
}

type createChannelResponse struct {
	ID string `json:"channel_id"`
}

// CreateChannel godoc
//
//	@Summary		Create a Channel for a user
//	@Accept			json
//	@Produce		json
//	@Param			request	body createChannelRequest true "Create Channel Request"
//	@Success		200	{object} createChannelResponse
//	@Router			/channel/create [post]
func (c *Handler) CreateChannel(e echo.Context) error {
	req := createChannelRequest{}
	err := e.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = e.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user_id, err := uuid.Parse(req.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	chann, err := c.channel.CreateChannel(user_id, req.Name)
	if err != nil {
		return err
	}

	resp := &createChannelResponse{
		ID: chann.ID.String(),
	}

	return e.JSON(http.StatusOK, resp)
}
