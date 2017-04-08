package health

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Route is a collection of routes for healthchecks
func Route(r *echo.Group, l *logrus.Logger) error {
	r.GET("/check", func(c echo.Context) error {
		t, err := time.Now().MarshalText()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, string(t))
	})
	r.GET("/ping", func(c echo.Context) error {
		l.Debug("Ping")
		return c.String(http.StatusOK, "Pong")
	})
	return nil
}
