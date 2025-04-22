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
	"github.com/katallaxie/fiber-htmx/components/buttons"
	"github.com/katallaxie/fiber-htmx/components/heroicons"
	"github.com/katallaxie/fiber-htmx/components/lists"
	"github.com/spf13/cobra"
	"github.com/zeiss/pkg/server"
)

type Page struct {
	htmx.Node
}

func NewPage(title, body string) *Page {
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
					lists.List(
						lists.ListProps{},
						lists.ListHeader(
							lists.ListHeaderProps{},
							htmx.Text("Most played songs this week"),
						),
						lists.ListRow(
							lists.ListRowProps{},
							htmx.Div(
								htmx.ClassNames{},
								htmx.Img(
									htmx.ClassNames{
										"size-10":     true,
										"rounded-box": true,
									},
									htmx.Src("https://img.daisyui.com/images/profile/demo/1@94.webp"),
								),
							),
							htmx.Div(
								htmx.ClassNames{},
								htmx.Div(
									htmx.ClassNames{},
									htmx.Text("Dio Lupa"),
								),
								htmx.Div(
									htmx.ClassNames{
										"text-xs":       true,
										"uppercase":     true,
										"font-semibold": true,
										"opacity-60":    true,
									},
									htmx.Text("Remaining Reason"),
								),
							),
							buttons.Ghost(
								buttons.ButtonProps{},
								heroicons.PlayOutline(
									heroicons.IconProps{},
								),
							),
							buttons.Ghost(
								buttons.ButtonProps{},
								heroicons.HeartOutline(
									heroicons.IconProps{},
								),
							),
						),
						lists.ListRow(
							lists.ListRowProps{},
							htmx.Div(
								htmx.ClassNames{},
								htmx.Img(
									htmx.ClassNames{
										"size-10":     true,
										"rounded-box": true,
									},
									htmx.Src("https://img.daisyui.com/images/profile/demo/1@94.webp"),
								),
							),
							htmx.Div(
								htmx.ClassNames{},
								htmx.Div(
									htmx.ClassNames{},
									htmx.Text("Dio Lupa"),
								),
								htmx.Div(
									htmx.ClassNames{
										"text-xs":       true,
										"uppercase":     true,
										"font-semibold": true,
										"opacity-60":    true,
									},
									htmx.Text("Remaining Reason"),
								),
							),
							buttons.Ghost(
								buttons.ButtonProps{},
								heroicons.PlayOutline(
									heroicons.IconProps{},
								),
							),
							buttons.Ghost(
								buttons.ButtonProps{},
								heroicons.HeartOutline(
									heroicons.IconProps{},
								),
							),
						),
						lists.ListRow(
							lists.ListRowProps{},
							htmx.Div(
								htmx.ClassNames{},
								htmx.Img(
									htmx.ClassNames{
										"size-10":     true,
										"rounded-box": true,
									},
									htmx.Src("https://img.daisyui.com/images/profile/demo/1@94.webp"),
								),
							),
							htmx.Div(
								htmx.ClassNames{},
								htmx.Div(
									htmx.ClassNames{},
									htmx.Text("Dio Lupa"),
								),
								htmx.Div(
									htmx.ClassNames{
										"text-xs":       true,
										"uppercase":     true,
										"font-semibold": true,
										"opacity-60":    true,
									},
									htmx.Text("Remaining Reason"),
								),
							),
							buttons.Ghost(
								buttons.ButtonProps{},
								heroicons.PlayOutline(
									heroicons.IconProps{},
								),
							),
							buttons.Ghost(
								buttons.ButtonProps{},
								heroicons.HeartOutline(
									heroicons.IconProps{},
								),
							),
						),
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

		app.Get("/", htmx.NewHandler(NewPage("Hello World", "Hello World")))

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
