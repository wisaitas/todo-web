package main

import "github.com/wisaitas/todo-web/internal/initial"

func main() {
	app := initial.InitializeApp()

	app.SetupMiddlewares()

	app.SetupRoutes()

	app.Run()
}
