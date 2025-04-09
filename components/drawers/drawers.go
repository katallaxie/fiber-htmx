package drawers

import htmx "github.com/katallaxie/fiber-htmx"

// DrawerOpenProps is the props for the DrawerOpen component
type DrawerOpenProps struct {
	ID         string
	ClassNames htmx.ClassNames
}

// DrawerOpenButton is a component that renders a drawer open button
func DrawerOpenButton(p DrawerOpenProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"btn":           true,
				"btn-primary":   true,
				"drawer-button": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("for", p.ID),
		htmx.Group(children...),
	)
}

// DrawerSideProps is a component that renders a drawer side
type DrawerSideProps struct {
	ID         string
	ClassNames htmx.ClassNames
}

// DrawerSide is a component that renders a drawer side
func DrawerSide(p DrawerSideProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"drawer-side": true,
			},
			p.ClassNames,
		),
		htmx.Label(
			htmx.ClassNames{
				"drawer-overlay": true,
			},
			htmx.For(p.ID),
			htmx.Aria("label", "Close drawer"),
		),
		htmx.Group(children...),
	)
}

// DrawerSideMenuProps is the props for the DrawerSideMenu component
type DrawerSideMenuProps struct {
	ID         string
	ClassNames htmx.ClassNames
}

// DrawerSideMenu is a component that renders a drawer side menu
func DrawerSideMenu(p DrawerSideMenuProps, children ...htmx.Node) htmx.Node {
	return htmx.Ul(
		htmx.Merge(
			htmx.ClassNames{
				"menu":              true,
				"p-4":               true,
				"w-80":              true,
				"min-h-full":        true,
				"bg-base-200":       true,
				"text-base-content": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DrawerSideMenuItemProps is the props for the DrawerSideMenuItem component
type DrawerSideMenuItemProps struct {
	ID         string
	ClassNames htmx.ClassNames
}

// DrawerSideMenuItem is a component that renders a drawer side menu item
func DrawerSideMenuItem(p DrawerSideMenuItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			htmx.ClassNames{},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DrawerContentProps is the props for the DrawerContent component
type DrawerContentProps struct {
	ID         string
	ClassNames htmx.ClassNames
}

// DrawerContent is a component that renders a drawer content
func DrawerContent(p DrawerContentProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"drawer-content": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DrawerProps is the props for the Drawer component
type DrawerProps struct {
	ID         string
	ClassNames htmx.ClassNames
	DrawerSide []htmx.Node
}

// Drawer is a component that renders a drawer
func Drawer(p DrawerProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"drawer": true,
			},
			p.ClassNames,
		),
		htmx.Input(
			htmx.ID(p.ID),
			htmx.Attribute("type", "checkbox"),
			htmx.ClassNames{
				"drawer-toggle": true,
			},
		),
		htmx.Group(children...),
	)
}
