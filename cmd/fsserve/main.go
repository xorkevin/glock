package main

import (
	"github.com/xorkevin/glock/staticfs"
)

func main() {
	s := staticfs.New(staticfs.NewConfig())
	s.Start(3000)
}
