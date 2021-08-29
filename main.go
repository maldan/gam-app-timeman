package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-timeman/internal/app/timeman"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
