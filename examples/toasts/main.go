package main

import (
	"context"
	"log"
	"net/http"
	"os"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/avatars"
	"github.com/katallaxie/fiber-htmx/components/buttons"
	"github.com/katallaxie/fiber-htmx/components/drawers"
	"github.com/katallaxie/fiber-htmx/components/dropdowns"
	"github.com/katallaxie/fiber-htmx/components/heroicons"
	"github.com/katallaxie/fiber-htmx/components/menus"
	"github.com/katallaxie/fiber-htmx/components/navbars"
	"github.com/katallaxie/fiber-htmx/components/swap"
	"github.com/katallaxie/fiber-htmx/components/tailwind"
	"github.com/katallaxie/fiber-htmx/components/toasts"
	"github.com/katallaxie/pkg/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/cobra"
)

// Config ...
type Config struct {
	Flags *Flags
}

// Flags ...
type Flags struct {
	Addr string
}

var cfg = &Config{
	Flags: &Flags{},
}

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":3000", "addr")
	rootCmd.SilenceUsage = true
}

type exampleController struct {
	htmx.UnimplementedController
}

func (c *exampleController) Error(err error) error {
	return err
}

func (c *exampleController) Post() error {
	return toasts.Error("Using the htmx-toast component.")
}

func (c *exampleController) Get() error {
	return c.Render(
		htmx.HTML5(
			htmx.HTML5Props{
				Title:    "index",
				Language: "en",
				Attributes: []htmx.Node{
					htmx.DataAttribute("theme", "light"),
				},
				Head: []htmx.Node{
					htmx.Link(
						htmx.Attribute("href", "/assets/output.css"),
						htmx.Attribute("rel", "stylesheet"),
						htmx.Attribute("type", "text/css"),
					),
					htmx.Script(
						htmx.Attribute("src", "https://unpkg.com/htmx.org@2.0.2"),
						htmx.Attribute("type", "application/javascript"),
					),
				},
			},
			htmx.Body(
				htmx.HxBoost(true),
				htmx.HxHeaders(map[string]string{"X-CSRFToken": "123456"}),
				toasts.Toasts(),
				htmx.Div(
					htmx.ClassNames{},
					drawers.Drawer(
						drawers.DrawerProps{
							ID: "global-drawer",
							ClassNames: htmx.ClassNames{
								"lg:drawer-open": true,
							},
						},
						drawers.DrawerContent(
							drawers.DrawerContentProps{
								ID: "drawer",
							},
							htmx.Div(
								htmx.ClassNames{
									tailwind.OverflowAuto: true,
									tailwind.WFull:        true,
									tailwind.HScreen:      true,
									tailwind.MaxWFull:     true,
								},
								htmx.Div(
									htmx.ClassNames{
										tailwind.Flex:    true,
										tailwind.HFull:   true,
										tailwind.FlexCol: true,
										"bg-base-200":    true,
									},
									navbars.Navbar(
										navbars.NavbarProps{
											ClassNames: htmx.ClassNames{
												"navbar":      true,
												"z-10":        true,
												"border-b":    true,
												"px-3":        true,
												"bg-base-100": true,
											},
										},
										navbars.NavbarStart(
											navbars.NavbarStartProps{
												ClassNames: htmx.ClassNames{
													"gap-3": true,
												},
											},
											drawers.DrawerOpenButton(
												drawers.DrawerOpenProps{
													ID: "global-drawer",
													ClassNames: htmx.ClassNames{
														"btn-sm":      true,
														"btn-square":  true,
														"btn-primary": false,
													},
												},
												heroicons.Bars3Outline(
													heroicons.IconProps{},
												),
											),
										),
										navbars.NavbarEnd(
											navbars.NavbarEndProps{},
											swap.Swap(
												swap.SwapProps{
													ClassNames: htmx.ClassNames{
														"swap-rotate": true,
													},
												},
												htmx.Input(
													htmx.Class("theme-controller"),
													htmx.Value("dark"),
													htmx.Attribute("type", "checkbox"),
												),
												swap.SwapOn(
													swap.SwapProps{},
													heroicons.MoonOutline(
														heroicons.IconProps{},
													),
												),
												swap.SwapOff(
													swap.SwapProps{},
													heroicons.SunOutline(
														heroicons.IconProps{},
													),
												),
											),
											buttons.CircleSmall(
												buttons.ButtonProps{},
												heroicons.BellAlertOutline(
													heroicons.IconProps{},
												),
											),
											dropdowns.Dropdown(
												dropdowns.DropdownProps{
													ClassNames: htmx.ClassNames{
														"dropdown-end": true,
													},
												},
												dropdowns.DropdownButton(
													dropdowns.DropdownButtonProps{
														ClassNames: htmx.ClassNames{
															"btn-sm":     true,
															"btn-circle": true,
															"btn-ghost":  true,
														},
													},
													avatars.AvatarRoundSmall(
														avatars.AvatarProps{},
														htmx.Img(
															htmx.Attribute("src", "https://avatars.githubusercontent.com/u/570959?v=4"),
														),
													),
												),
												dropdowns.DropdownMenuItems(
													dropdowns.DropdownMenuItemsProps{},
													dropdowns.DropdownMenuItem(
														dropdowns.DropdownMenuItemProps{},
														htmx.A(
															htmx.Text("Profile"),
														),
													),
												),
											),
										),
									),
									htmx.Div(
										buttons.Button(
											buttons.ButtonProps{},
											htmx.HxPost("/notify"),
											htmx.Text("Notify"),
										),
									),
								),
							),
						),
						drawers.DrawerSide(
							drawers.DrawerSideProps{
								ID: "drawer",
							},
							drawers.DrawerSideMenu(
								drawers.DrawerSideMenuProps{
									ClassNames: htmx.ClassNames{
										"border-r":    true,
										"bg-base-100": true,
										"bg-base-200": false,
									},
								},
								htmx.Nav(
									htmx.Merge(
										htmx.ClassNames{},
									),
									menus.Menu(
										menus.MenuProps{
											ClassNames: htmx.ClassNames{
												"w-full":      true,
												"bg-base-200": false,
											},
										},
										menus.MenuItem(
											menus.MenuItemProps{},
											menus.MenuLink(
												menus.MenuLinkProps{},
												htmx.Text("Home"),
											),
										),
										menus.MenuTitle(
											menus.MenuTitleProps{},
											htmx.Text("Components"),
										),
										menus.MenuItem(
											menus.MenuItemProps{},
											menus.MenuCollapsible(
												menus.MenuCollapsibleProps{},
												menus.MenuCollapsibleSummary(
													menus.MenuCollapsibleSummaryProps{},
													htmx.Text("Buttons"),
												),
												menus.MenuItem(
													menus.MenuItemProps{},
													htmx.A(
														htmx.Attribute("href", "#"),
														htmx.Text("Item 1"),
													),
													htmx.A(
														htmx.Attribute("href", "#"),
														htmx.Text("Item 2"),
													),
													htmx.A(
														htmx.Attribute("href", "#"),
														htmx.Text("Item 3"),
													),
												),
											),
											menus.MenuCollapsible(
												menus.MenuCollapsibleProps{},
												menus.MenuCollapsibleSummary(
													menus.MenuCollapsibleSummaryProps{},
													htmx.Text("Forms"),
												),
												menus.MenuItem(
													menus.MenuItemProps{},
													htmx.A(
														htmx.Attribute("href", "#"),
														htmx.Text("Item 1"),
													),
													htmx.A(
														htmx.Attribute("href", "#"),
														htmx.Text("Item 2"),
													),
													htmx.A(
														htmx.Attribute("href", "#"),
														htmx.Text("Item 3"),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		),
	)
}

type webSrv struct{}

func (w *webSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		app := fiber.New(
			fiber.Config{
				ErrorHandler: toasts.DefaultErrorHandler,
			},
		)
		app.Use(requestid.New())
		app.Use(logger.New())
		app.Use(recover.New())
		app.Use("/assets", filesystem.New(filesystem.Config{
			Root: http.Dir("./dist"),
		}))

		app.Get("/", htmx.NewHxControllerHandler(func() htmx.Controller {
			return &exampleController{}
		}))

		app.Post("/notify", htmx.NewHxControllerHandler(func() htmx.Controller {
			return &exampleController{}
		}))

		err := app.Listen(cfg.Flags.Addr)
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

	s.Listen(webSrv, true)

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
