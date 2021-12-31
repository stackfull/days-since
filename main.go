package main

import (
	"log"

	"github.com/kataras/iris/v12"

	since "github.com/stackfull/days-since/src"
)

func main() {
	app := iris.New()

	app.PartyFunc("/since", since.Router())

	// $ go-bindata -prefix "app/build" -fs ./app/build/...
	app.HandleDir("/", AssetFile())

	if err := app.Listen(":8080"); err != nil {
		log.Fatal("app.Listen", err)
	}
}
