package main

import (
	"log"
	"os"

	handler "github.com/katallaxie/fiber-htmx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	reload "github.com/katallaxie/fiber-reload"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/spf13/pflag"
)

var addr = ":3000"

func Demo() htmx.Node {
	return htmx.HTML5(
		htmx.HTML5Props{
			Title:    "Demo",
			Language: "en",
			Head: []htmx.Node{
				htmx.Script(
					htmx.Src("https://unpkg.com/fiber-reload@0.9.0/reload.js"),
					htmx.Type("text/javascript"),
				),
			},
		},
		htmx.Body(
			buttons.Button(buttons.ButtonProps{}, htmx.Text("Demo")),
		),
	)
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	pflag.StringVar(&addr, "addr", addr, "address to listen on")

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(reload.Environment(reload.Development))
	reload.WithHotReload(app)

	app.Get("/", handler.NewHandler(Demo()))

	err := app.Listen(addr)
	if err != nil {
		log.Fatal(err)
	}
}
