package main

import (
	"github.com/xorkevin/glock/server"
)

func main() {
	g := server.New()
	g.Start("8080")
}
