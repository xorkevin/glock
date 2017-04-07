package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

type (
	// Server is an http gateway
	Server struct {
		i   *echo.Echo
		log *logrus.Logger
	}
)

// New creates a new Server
func New() *Server {
	// logger
	l := logrus.New()
	l.Formatter = &logrus.TextFormatter{}
	l.Out = os.Stdout

	// http server instance
	i := echo.New()

	// middleware
	i.Use(middleware.Logger())
	i.Use(middleware.Recover())
	i.Use(middleware.RemoveTrailingSlash())

	// routes
	i.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// api routes
	api := i.Group("/api")
	api.GET("/health", func(c echo.Context) error {
		t, err := time.Now().MarshalText()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, string(t))
	})
	api.GET("/ping", func(c echo.Context) error {
		l.Debug("Ping")
		return c.String(http.StatusOK, "Pong")
	})

	return &Server{
		i:   i,
		log: l,
	}
}

// Start starts the server at the specified port
func (s *Server) Start(port uint) error {
	s.i.Logger.Fatal(s.i.Start(":" + strconv.Itoa(int(port))))
	return nil
}
