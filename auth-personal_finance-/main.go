package main

import (
	"auth/config"
	"auth/internal/app"
)

func main() {
	cfg := config.Load()
	app.Run(&cfg)
}
