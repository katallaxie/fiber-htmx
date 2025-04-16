package main

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-playground/validator/v10"
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/alerts"
	"github.com/katallaxie/fiber-htmx/components/avatars"
	"github.com/katallaxie/fiber-htmx/components/buttons"
	"github.com/katallaxie/fiber-htmx/components/drawers"
	"github.com/katallaxie/fiber-htmx/components/dropdowns"
	"github.com/katallaxie/fiber-htmx/components/forms"
	"github.com/katallaxie/fiber-htmx/components/heroicons"
	"github.com/katallaxie/fiber-htmx/components/menus"
	"github.com/katallaxie/fiber-htmx/components/mockups"
	"github.com/katallaxie/fiber-htmx/components/navbars"
	"github.com/katallaxie/fiber-htmx/components/swaps"
	"github.com/katallaxie/fiber-htmx/components/tables"
	"github.com/katallaxie/fiber-htmx/components/tabs"
	"github.com/katallaxie/fiber-htmx/components/tailwind"
	"github.com/katallaxie/fiber-htmx/components/typography"
	"github.com/katallaxie/fiber-htmx/components/utils"
	"github.com/katallaxie/fiber-htmx/components/validate"
	"github.com/katallaxie/pkg/utilx"
)

// DemoRow ...
type DemoRow struct {
	ID   int
	Name string
}

//nolint:mnd
var demoRows = []DemoRow{
	{
		ID:   1,
		Name: "Name 1",
	},
	{
		ID:   2,
		Name: "Name 2",
	},
	{
		ID:   3,
		Name: "Name 3",
	},
	{
		ID:   4,
		Name: "Name 4",
	},
	{
		ID:   5,
		Name: "Name 5",
	},
	{
		ID:   6,
		Name: "Name 6",
	},
	{
		ID:   7,
		Name: "Name 7",
	},
	{
		ID:   8,
		Name: "Name 8",
	},
	{
		ID:   9,
		Name: "Name 9",
	},
	{
		ID:   10,
		Name: "Name 10",
	},
}

type RegisterFormProps struct {
	errors validate.Errors
}

func RegisterForm(props RegisterFormProps) htmx.Node {
	return htmx.Form(
		htmx.ClassNames{
			"group": true,
		},
		htmx.HxPost("/register"),
		htmx.HxSwap("outerHTML"),
		htmx.Target("cloesest div"),
		htmx.NoValidate(),
		forms.FormControl(
			forms.FormControlProps{},
			forms.FormControlLabel(
				forms.FormControlLabelProps{},
				forms.FormControlLabelText(
					forms.FormControlLabelTextProps{
						ClassNames: htmx.ClassNames{
							"label-text": true,
						},
					},
					htmx.Text("Email"),
				),
			),
			forms.TextInput(
				forms.TextInputProps{
					Error:       props.errors.Field("email"),
					Name:        "email",
					Value:       "",
					Type:        "email",
					Placeholder: "indy@example.com",
					ClassNames: htmx.ClassNames{
						"peer": true,
						"invalid:[&:not(:placeholder-shown):not(:focus)]:border-red-500": true,
					},
				},
				htmx.Attribute("type", "email"),
				htmx.Required(),
			),
			htmx.Span(
				htmx.ClassNames{
					"mt-2":         true,
					"hidden":       true,
					"text-sm":      true,
					"text-red-500": true,
					"peer-[&:not(:placeholder-shown):not(:focus):invalid]:block": true,
				},
				htmx.Text("A valid email is required."),
			),
			htmx.If(
				props.errors.HasError("email"),
				htmx.Div(
					htmx.ClassNames{
						"text-error": true,
					},
					htmx.Text("A valid email is required."),
				),
			),
		),
		forms.FormControl(
			forms.FormControlProps{},
			forms.FormControlLabel(
				forms.FormControlLabelProps{},
				forms.FormControlLabelText(
					forms.FormControlLabelTextProps{
						ClassNames: htmx.ClassNames{
							"label-text": true,
						},
					},
					htmx.Text("Option 1"),
				),
			),
			forms.Radio(
				forms.RadioProps{
					Name:  "option",
					Value: "1",
					Error: props.errors.Field("option"),
				},
				htmx.Required(),
			),
			htmx.If(
				props.errors.HasError("option"),
				htmx.Div(
					htmx.ClassNames{
						"text-error": true,
					},
					htmx.Text("A valid option is required."),
				),
			),
		),
		buttons.Button(
			buttons.ButtonProps{
				Type: "submit",
				ClassNames: htmx.ClassNames{
					"group-invalid:pointer-events-none group-invalid:opacity-30": true,
				},
			},
			htmx.Text("Register"),
		),
	)
}

type ExampleForm struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required,min=10"`
}

func Page() htmx.Node {
	v := validator.New()

	form := &ExampleForm{
		Name:        "Demo",
		Description: "",
	}

	var verr validator.ValidationErrors
	err := v.Struct(form)
	errors.As(err, &verr)

	errz := validate.Errors(verr)

	return htmx.HTML5(
		htmx.HTML5Props{
			Title:    "index",
			Language: "en",
			Attributes: []htmx.Node{
				htmx.DataAttribute("theme", "light"),
			},
			Head: []htmx.Node{
				htmx.Link(
					htmx.Attribute("href", "https://cdn.jsdelivr.net/npm/daisyui/dist/full.css"),
					htmx.Attribute("rel", "stylesheet"),
					htmx.Attribute("type", "text/css"),
				),
				htmx.Link(
					htmx.Attribute("href", "/static/out.css"),
					htmx.Attribute("rel", "stylesheet"),
					htmx.Attribute("type", "text/css"),
				),
				htmx.Imports(htmx.ImportsProp{}),
				htmx.Script(
					htmx.Async(),
					htmx.CrossOrigin("anonymous"),
					htmx.Src("https://unpkg.com/es-module-shims@1.10.0/dist/es-module-shims.js"),
				),
				htmx.Script(
					htmx.Type("module"),
					htmx.Raw(`import htmx from "htmx";`),
					htmx.Raw(`htmx.logAll();`),
				),
				htmx.Script(
					htmx.Attribute("src", "https://unpkg.com/htmx-ext-sse@2.0.0/sse.js"),
					htmx.Attribute("type", "application/javascript"),
				),
				htmx.Script(
					htmx.Attribute("src", "/static/out.js"),
				),
			},
		},
		htmx.Body(
			htmx.HxBoost(true),
			htmx.HxHeaders(map[string]string{"X-CSRFToken": "123456"}),
			htmx.Fragment(
				htmx.Div(),
				htmx.Div(),
			),
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
										swaps.Swap(
											swaps.Props{
												ClassNames: htmx.ClassNames{
													"swap-rotate": true,
												},
											},
											htmx.Input(
												htmx.Class("theme-controller peer-invalid:text-red-500"),
												htmx.Value("dark"),
												htmx.Attribute("type", "checkbox"),
											),
											swaps.On(
												swaps.Props{},
												heroicons.MoonOutline(
													heroicons.IconProps{},
												),
											),
											swaps.Off(
												swaps.Props{},
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
									mockups.Code(
										mockups.CodeProps{},
									),
									swaps.Swap(
										swaps.Props{
											ClassNames: htmx.ClassNames{
												"swap-rotate": true,
											},
										},
										// Write a Hyperscript that toggles a checkbox inside the current element
										// 										htmx.HyperScript(`on click
										// writeText(my previousElementSibling's innerText) on navigator.clipboard
										// set <input/> in me to checked to not checked
										// wait 1s
										// set <input/> in me to checked to true`),
										htmx.Input(
											htmx.Value("copy"),
											htmx.Attribute("type", "checkbox"),
											htmx.Checked(),
										),
										swaps.On(
											swaps.Props{},
											heroicons.MoonOutline(
												heroicons.IconProps{},
											),
										),
										swaps.Off(
											swaps.Props{},
											heroicons.SunOutline(
												heroicons.IconProps{},
											),
										),
									),
								),
								htmx.Div(
									htmx.ClassNames{
										"bg-base-100": true,
									},
									tabs.Boxed(
										tabs.Props{},
										tabs.Tab(
											tabs.TabProps{
												Name:   "my_tabs_1",
												Label:  "Tab 1",
												Active: true,
											},
											htmx.Text("Tab 1"),
										),
										tabs.Tab(
											tabs.TabProps{
												Name:  "my_tabs_1",
												Label: "Tab 2",
											},
											htmx.Text("Tab 2"),
										),
										tabs.Tab(
											tabs.TabProps{
												Name:  "my_tabs_1",
												Label: "Tab 3",
											},
											htmx.Text("Tab 3"),
										),
									),
								),
								RegisterForm(RegisterFormProps{}),
								alerts.Info(
									alerts.AlertProps{},
									htmx.Text("Hello, World!"),
								),
								htmx.Div(
									htmx.HxSSE(),
									htmx.HxSSEConnect("/sse"),
									htmx.HxSSESwap("demo"),
								),
								htmx.Div(
									htmx.ID("events"),
								),
								htmx.Div(
									forms.TextInput(
										forms.TextInputProps{},
									),
									forms.FormControl(
										forms.FormControlProps{},
										forms.FormControlLabel(
											forms.FormControlLabelProps{},
											forms.FormControlLabelText(
												forms.FormControlLabelTextProps{
													ClassNames: htmx.ClassNames{
														"label-text": true,
													},
												},
												htmx.Text("Hello, World!"),
											),
										),
										forms.TextInput(
											forms.TextInputProps{
												Error: errz.Field("name"),
												Value: "Hello, World!",
											},
										),
										htmx.If(
											utilx.NotEmpty(errz.Field("name")),
											typography.Error(
												typography.Props{},
												htmx.Text("This field is required."),
											),
										),
									),
									forms.FormControl(
										forms.FormControlProps{},
										forms.FormControlLabel(
											forms.FormControlLabelProps{},
											forms.FormControlLabelText(
												forms.FormControlLabelTextProps{
													ClassNames: htmx.ClassNames{
														"label-text": true,
													},
												},
												htmx.Text("Hello, World!"),
											),
										),
										forms.TextInput(
											forms.TextInputProps{
												Error: errz.Field("Description"),
												Value: "Hello, World!",
											},
										),
										htmx.If(
											utilx.NotEmpty(errz.Field("Description")),
											typography.Error(
												typography.Props{},
												htmx.Text("This field is required."),
											),
										),
									),
									htmx.Div(
										htmx.Merge(
											htmx.ClassNames{
												"text-bg-error": true,
											},
										),
										htmx.ID("field_error"),
									),
									htmx.Form(
										htmx.Method("POST"),
										htmx.HxSwap("none"),
										buttons.Ghost(
											buttons.ButtonProps{},
											htmx.HxPost("/error"),
											htmx.Type("submit"),
											htmx.Text("Submit"),
										),
									),
								),
								htmx.Fallback(htmx.Markdown([]byte("## Hello, Markdown!")), func(_ error) htmx.Node {
									return htmx.Text("Fallback")
								}),
								htmx.Fallback(
									htmx.ErrorBoundary(
										func() htmx.Node {
											return utils.Panic(errors.New("this is a new panic"))
										},
									),
									func(err error) htmx.Node {
										return htmx.Text(err.Error())
									},
								),
								htmx.Raw(`<multi-select>Click me</multi-select>`),
								dropdowns.Dropdown(
									dropdowns.DropdownProps{
										ClassNames: htmx.ClassNames{},
									},
									dropdowns.DropdownButton(
										dropdowns.DropdownButtonProps{
											ClassNames: htmx.ClassNames{
												"btn-outline": true,
											},
										},
										htmx.ID("dropdown"),
										htmx.Text("Dropdown"),
									),
									dropdowns.DropdownMenuItems(
										dropdowns.DropdownMenuItemsProps{},
										forms.TextInputBordered(
											forms.TextInputProps{
												ClassNames: htmx.ClassNames{
													"input-sm": true,
												},
												Placeholder: "Search...",
											},
											htmx.HxPost("/dropdown"),
											htmx.HxTrigger("input changed delay:500ms, search"),
											htmx.HxTarget("#search-results"),
										),
										htmx.Div(
											htmx.ID("search-results"),
										),
									),
								),
								htmx.Div(
									tables.Table(
										tables.TableProps{
											Pagination: tables.TablePagination(
												tables.TablePaginationProps{},
												tables.Pagination(
													tables.PaginationProps{
														Total:  len(demoRows),
														Offset: 0,
														Limit:  10, //nolint:mnd
													},
												),
											),
										},
										tables.Columns[DemoRow]{
											{
												ID:          "id",
												AccessorKey: "ID",
												Header: func(_ tables.TableProps) htmx.Node {
													return htmx.Th(htmx.Text("ID"))
												},
												Cell: func(_ tables.TableProps, row DemoRow) htmx.Node {
													return htmx.Td(
														htmx.Text(strconv.Itoa(row.ID)),
													)
												},
											},
											{
												ID:          "name",
												AccessorKey: "Name",
												Header: func(_ tables.TableProps) htmx.Node {
													return htmx.Th(htmx.Text("Name"))
												},
												Cell: func(_ tables.TableProps, row DemoRow) htmx.Node {
													return htmx.Td(htmx.Text(row.Name))
												},
											},
										},
										demoRows,
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
	)
}

func main() {
	cwd, err := os.Getwd() // get the current directory using the built-in function
	if err != nil {
		panic(err)
	}

	//nolint:mnd
	f, err := os.OpenFile(filepath.Join(filepath.Clean(cwd), "examples/index.html"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = Page().Render(f)
	if err != nil {
		panic(err)
	}
}
