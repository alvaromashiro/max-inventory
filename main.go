package main

import (
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(Settings.New),
		fx.Invoke(),
	)

	app.Run()
}
