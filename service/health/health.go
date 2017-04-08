package health

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/xorkevin/glock"
	"net/http"
	"time"
)

type (
	// Health is a health service for monitoring
	Health struct {
	}
)

// New creates a new Health service
func New() *Health {
	return &Health{}
}

// Mount is a collection of routes for healthchecks
func (h *Health) Mount(conf glock.Config, r *echo.Group, l *logrus.Logger) error {
	r.GET("/check", func(c echo.Context) error {
		t, err := time.Now().MarshalText()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, string(t))
	})
	r.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, conf.Version)
	})
	r.GET("/ping", func(c echo.Context) error {
		l.Debug("Ping")
		return c.String(http.StatusOK, "Pong")
	})
	return nil
}
