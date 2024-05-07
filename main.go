package main

import (
	"embed"

	"github.com/c-malecki/some-game-thing/core"
)

//go:embed all:data/*
var dataFS embed.FS

func main() {
	core.RunGame(dataFS)
}
