package menus

import htmx "github.com/katallaxie/fiber-htmx"

// MenuProps is the struct for the menu props.
type MenuProps struct {
	ClassNames htmx.ClassNames
}

// Menu is the component for the menu.
func Menu(p MenuProps, children ...htmx.Node) htmx.Node {
	return htmx.Ul(
		htmx.Merge(
			htmx.ClassNames{
				"menu":        true,
				"bg-base-200": true,
				"w-56":        true,
				"rounded-box": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// MenuItemProps is the struct for the menu item props.
type MenuItemProps struct {
	ClassNames htmx.ClassNames
}

// MenuItem is the component for the menu item.
func MenuItem(p MenuItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			htmx.ClassNames{
				"menu-item": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// MenuTitleProps is the struct for the menu title props.
type MenuTitleProps struct {
	ClassNames htmx.ClassNames
}

// MenuTitle is the component for the menu title.
func MenuTitle(p MenuTitleProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			htmx.ClassNames{
				"menu-title": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// MenuCollapsibleProps is the struct for the menu collapse props.
type MenuCollapsibleProps struct {
	ClassNames htmx.ClassNames
	Open       bool
}

// MenuCollapsible is the component for the menu collapse.
func MenuCollapsible(p MenuCollapsibleProps, children ...htmx.Node) htmx.Node {
	return htmx.Details(
		htmx.Merge(
			p.ClassNames,
		),
		htmx.If(p.Open, htmx.Attribute("open", "open")),
		htmx.Group(children...),
	)
}

// MenuCollapsibleSummaryProps is the struct for the menu collapse summary props.
type MenuCollapsibleSummaryProps struct {
	ClassNames htmx.ClassNames
}

// MenuCollapsibleSummary is the component for the menu collapse summary.
func MenuCollapsibleSummary(p MenuCollapsibleSummaryProps, children ...htmx.Node) htmx.Node {
	return htmx.Summary(
		htmx.Merge(
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// MenuLinkProps is the struct for the menu link props.
type MenuLinkProps struct {
	ClassNames htmx.ClassNames
	Href       string
	Active     bool
}

// MenuLink is the component for the menu link.
func MenuLink(p MenuLinkProps, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"active": p.Active,
			},
			p.ClassNames,
		),
		htmx.Href(p.Href),
		htmx.Group(children...),
	)
}
