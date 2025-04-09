package navbars

import htmx "github.com/katallaxie/fiber-htmx"

// NavbarProps are the properties of the Navbar component
type NavbarProps struct {
	ClassNames htmx.ClassNames
}

// Navbar represents a Daisy UI Navbar component.
// see: https://daisyui.com/components/navbar
func Navbar(props NavbarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"navbar":      true,
				"bg-base-100": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// NavbarStartProps represents the properties of the NavbarStart component
type NavbarStartProps struct {
	ClassNames htmx.ClassNames
}

// NavbarStart represents a Daisy UI NavbarStart component.
func NavbarStart(props NavbarStartProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"navbar-start": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// NavbarCenterProps are the properties of the NavbarCenter component
type NavbarCenterProps struct {
	ClassNames htmx.ClassNames
}

// NavbarCenter represents a Daisy UI NavbarCenter component.
func NavbarCenter(props NavbarCenterProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"navbar-center": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// NavbarEndProps are the properties of the NavbarEnd component
type NavbarEndProps struct {
	ClassNames htmx.ClassNames
}

// NavbarEnd represents a Daisy UI NavbarEnd component.
func NavbarEnd(props NavbarEndProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"navbar-end": true,
			},
		),
		htmx.Group(children...),
	)
}
