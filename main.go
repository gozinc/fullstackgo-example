package main

import (
	"fullstackgo/zinc"
	"log/slog"
)

func main() {
	app := zinc.New()

	err := app.Start()
	if err != nil {
		slog.Error(err.Error())
	}
}
