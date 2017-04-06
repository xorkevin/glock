package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"time"
)

type (
	// Server is a basic http gateway
	Server struct {
		i *echo.Echo
	}
)

// New creates a new Server
func New() *Server {
	i := echo.New()
	i.Use(middleware.Logger())
	i.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	i.GET("/health", func(c echo.Context) error {
		t, err := time.Now().MarshalText()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, string(t))
	})
	return &Server{
		i: i,
	}
}

// Start starts the server at the specified port
func (s *Server) Start(port string) {
	s.i.Logger.Fatal(s.i.Start(":" + port))
}
