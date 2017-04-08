package main

import (
	rHealth "github.com/xorkevin/glock/routes/health"
	"github.com/xorkevin/glock/server"
)

func main() {
	g := server.New()
	g.SetLoggingLevel(server.LevelDebug)
	g.RegisterRoute("/api/health", rHealth.Route)
	g.Start(8080)
}
