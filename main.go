package main

import (
	"github.com/codeS2-com/codeS2-backend/app"
)

func main() {
	//config := config.GetConfig()

	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
