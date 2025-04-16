package main

import (
	"os"

	htmx "github.com/katallaxie/fiber-htmx"
)

func Demo() htmx.Node {
	return htmx.HTML5(
		htmx.HTML5Props{
			Title:    "error",
			Language: "en",
			Head:     []htmx.Node{},
		},
		htmx.Div(
			htmx.Text("Hello world!"),
		),
	)
}

func main() {
	_ = Demo().Render(os.Stdout)
}
