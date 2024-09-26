package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// HealthHandler is the request handler for the health endpoint.
type HealthHandler interface {
	Healthz(c echo.Context) error
}

type healthHandler struct{}

// NewHealth returns a new instance of the health handler.
func NewHealth() HealthHandler {
	return &healthHandler{}
}

func (t *healthHandler) Healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, time.Now())
}