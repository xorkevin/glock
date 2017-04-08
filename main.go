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

const (
	levelDebug = iota
	levelInfo
	levelWarn
	levelError
	levelFatal
	levelPanic
)

func envToLevel(e string) int {
	switch e {
	case "DEBUG":
		return levelDebug
	case "INFO":
		return levelInfo
	case "WARN":
		return levelWarn
	case "ERROR":
		return levelError
	case "FATAL":
		return levelFatal
	case "PANIC":
		return levelPanic
	default:
		return levelInfo
	}
}

func levelToLog(level int) logrus.Level {
	switch level {
	case levelDebug:
		return logrus.DebugLevel
	case levelInfo:
		return logrus.InfoLevel
	case levelWarn:
		return logrus.WarnLevel
	case levelError:
		return logrus.ErrorLevel
	case levelFatal:
		return logrus.FatalLevel
	case levelPanic:
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

// NewConfig creates a new server configuration
// It requires ENV vars:
//   VERSION
//   MODE
func NewConfig() Config {
	return Config{
		Version:  os.Getenv("VERSION"),
		LogLevel: envToLevel(os.Getenv("MODE")),
	}
}

// IsDebug returns if the configuration is in debug mode
func (c *Config) IsDebug() bool {
	return c.LogLevel == levelDebug
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
	if config.LogLevel == levelDebug {
		i.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
		}))
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
