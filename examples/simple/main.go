package main

import (
	"log"
	"os"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/buttons"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/pflag"
	reload "github.com/zeiss/fiber-reload"
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

	app.Get("/", htmx.NewHandler(Demo()))

	err := app.Listen(addr)
	if err != nil {
		log.Fatal(err)
	}
}
