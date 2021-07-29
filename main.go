package main

import (
	"log"

	"github.com/nikitamirzani323/fiberwebsite/routers"
)

func main() {
	app := routers.Init()

	log.Fatal(app.Listen(":3000"))
}
