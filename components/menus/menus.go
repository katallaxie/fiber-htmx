package menus

import htmx "github.com/katallaxie/fiber-htmx"

// Props is the struct for the menu props.
type Props struct {
	htmx.ClassNames
}

// Menu is the component for the menu.
func Menu(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Ul(
		htmx.Merge(
			htmx.ClassNames{
				"menu": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// MenuVertical is the component for the vertical menu.
func MenuVertical(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Ul(
		htmx.Merge(
			htmx.ClassNames{
				"menu":          true,
				"menu-vertical": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// MenuHorizontal is the component for the horizontal menu.
func MenuHorizontal(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Ul(
		htmx.Merge(
			htmx.ClassNames{
				"menu":            true,
				"menu-horizontal": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ItemProps is the struct for the menu item props.
type ItemProps struct {
	htmx.ClassNames
}

// Item is the component for the menu item.
func Item(p ItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// TitleProps is the struct for the menu title props.
type TitleProps struct {
	htmx.ClassNames
}

// Title is the component for the menu title.
func Title(p TitleProps, children ...htmx.Node) htmx.Node {
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

// CollapsibleProps is the struct for the menu collapse props.
type CollapsibleProps struct {
	Open bool
	htmx.ClassNames
}

// Collapsible is the component for the menu collapse.
func Collapsible(p CollapsibleProps, children ...htmx.Node) htmx.Node {
	return htmx.Details(
		htmx.Merge(
			p.ClassNames,
		),
		htmx.If(p.Open, htmx.Attribute("open", "open")),
		htmx.Group(children...),
	)
}

// CollapsibleSummaryProps is the struct for the menu collapse summary props.
type CollapsibleSummaryProps struct {
	htmx.ClassNames
}

// CollapsibleSummary is the component for the menu collapse summary.
func CollapsibleSummary(p CollapsibleSummaryProps, children ...htmx.Node) htmx.Node {
	return htmx.Summary(
		htmx.Merge(
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// LinkProps is the struct for the menu link props.
type LinkProps struct {
	Href   string
	Active bool
	htmx.ClassNames
}

// Link is the component for the menu link.
func Link(p LinkProps, children ...htmx.Node) htmx.Node {
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
