package glock

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type (
	// Config is the server configuration
	Config struct {
		Version  string
		LogLevel int
	}

	// Server is an http gateway
	Server struct {
		i      *echo.Echo
		log    *logrus.Logger
		config Config
	}
)

// NewConfig creates a new server configuration
func NewConfig(loglevel int) Config {
	return Config{
		Version:  os.Getenv("VERSION"),
		LogLevel: loglevel,
	}
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

func levelToLog(level int) logrus.Level {
	switch level {
	case LevelDebug:
		return logrus.DebugLevel
	case LevelInfo:
		return logrus.InfoLevel
	case LevelWarning:
		return logrus.WarnLevel
	case LevelError:
		return logrus.ErrorLevel
	case LevelFatal:
		return logrus.FatalLevel
	case LevelPanic:
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

// New creates a new Server
func New(config Config) *Server {
	// logger
	l := logrus.New()
	l.Formatter = &logrus.TextFormatter{}
	l.Out = os.Stdout
	l.Level = levelToLog(config.LogLevel)

	// http server instance
	i := echo.New()

	// middleware
	if config.LogLevel == LevelDebug {
		i.Use(middleware.Logger())
	}
	i.Use(middleware.Recover())
	i.Use(middleware.RemoveTrailingSlash())

	return &Server{
		i:      i,
		log:    l,
		config: config,
	}
}

// Start starts the server at the specified port
func (s *Server) Start(port uint) error {
	s.i.Logger.Fatal(s.i.Start(":" + strconv.Itoa(int(port))))
	return nil
}

type (
	// Service is an interface for services
	Service interface {
		Mount(c Config, r *echo.Group, l *logrus.Logger) error
	}
)

// MountRoute mounts a service
func (s *Server) MountRoute(path string, r Service, m ...echo.MiddlewareFunc) error {
	return r.Mount(s.config, s.i.Group(path, m...), s.log)
}
