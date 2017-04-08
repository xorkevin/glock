package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
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

const (
	// LevelDebug is the Debug logging level
	LevelDebug = iota
	// LevelInfo is the Info logging level
	LevelInfo
	// LevelWarning is the Warning logging level
	LevelWarning
	// LevelError is the Error logging level
	LevelError
	// LevelFatal is the Fatal logging level
	LevelFatal
	// LevelPanic is the Panic logging level
	LevelPanic
)

// SetLoggingLevel sets the logger level
func (s *Server) SetLoggingLevel(level int) {
	k := logrus.InfoLevel
	switch level {
	case LevelDebug:
		k = logrus.DebugLevel
	case LevelInfo:
		k = logrus.InfoLevel
	case LevelWarning:
		k = logrus.WarnLevel
	case LevelError:
		k = logrus.ErrorLevel
	case LevelFatal:
		k = logrus.FatalLevel
	case LevelPanic:
		k = logrus.PanicLevel
	}
	s.log.Level = k
}

type (
	// Routes is a function that registers a group of routes
	Routes func(r *echo.Group, l *logrus.Logger) error
)

// RegisterRoute mounts a set of routes
func (s *Server) RegisterRoute(path string, r Routes, m ...echo.MiddlewareFunc) error {
	return r(s.i.Group(path, m...), s.log)
}
