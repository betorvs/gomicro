package controller

import (
	"net/http"

	"github.com/betorvs/gomicro/version"
	"github.com/labstack/echo/v4"
)

// Health struct
type Health struct {
	Status    string `json:"status"`
	Version   string `json:"version,omitempty"`
	BuildInfo string `json:"build_info,omitempty"`
}

// CheckHealth func to use by LB
func CheckHealth(c echo.Context) error {
	health := Health{}
	health.Status = "UP"
	health.Version = version.Version
	health.BuildInfo = version.BuildInfo
	return c.JSON(http.StatusOK, health)
}
