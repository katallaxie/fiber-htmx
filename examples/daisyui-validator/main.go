package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/forms"
	"github.com/spf13/cobra"
	"github.com/zeiss/pkg/server"
)

type Page struct {
	htmx.Node
}

func NewPage(title string) *Page {
	return &Page{
		Node: htmx.HTML5(
			htmx.HTML5Props{
				Title: title,
				Head: []htmx.Node{
					htmx.Link(
						htmx.Href("https://cdn.jsdelivr.net/npm/daisyui@5"),
						htmx.Rel("stylesheet"),
						htmx.Type("text/css"),
					),
					htmx.Script(
						htmx.Src("https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"),
					),
				},
			},
			htmx.Body(
				htmx.Div(
					htmx.ClassNames{
						"flex":         true,
						"flex-col":     true,
						"items-center": true,
						"p-4":          true,
					},
					forms.TextInput(
						forms.TextInputProps{
							ClassNames: htmx.ClassNames{
								"m-4": true,
							},
							Validator:   true,
							Type:        "email",
							Placeholder: "me@example.com",
						},
					),
					forms.TextInput(
						forms.TextInputProps{
							ClassNames: htmx.ClassNames{
								"m-4":  true,
								"mb-0": true,
							},
							Validator:   true,
							Type:        "email",
							Placeholder: "me@example.com",
						},
					),
					forms.ValidatorHint(
						forms.ValidatorHintProps{},
						htmx.Text("Enter valid email address"),
					),
				),
			),
		),
	}
}

var Addr string

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, _ []string) error {
		return run(cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&Addr, "addr", ":3000", "addr")
	rootCmd.SilenceUsage = true
}

type webSrv struct{}

func (w *webSrv) Start(_ context.Context, _ server.ReadyFunc, _ server.RunFunc) func() error {
	return func() error {
		app := fiber.New()
		app.Use(requestid.New())
		app.Use(logger.New())
		app.Use(recover.New())

		app.Get("/", htmx.NewHandler(NewPage("Hello World")))

		err := app.Listen(Addr)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	}
}

func run(ctx context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	webSrv := &webSrv{}

	s, _ := server.WithContext(ctx)
	s.Listen(webSrv, false)

	err := s.Wait()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
