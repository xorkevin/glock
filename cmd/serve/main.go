package main

import (
	"github.com/xorkevin/glock"
	"github.com/xorkevin/glock/service/health"
)

func main() {
	g := glock.New(glock.NewConfig(glock.LevelDebug))

	hS := health.New()

	g.MountRoute("/api/health", hS)
	g.Start(8080)
}
