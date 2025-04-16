package alerts

import (
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/heroicons"
)

// AlertProps is the type of the props for the Alert component
type AlertProps struct {
	// ClassNames are the class names for the alert component.
	ClassNames htmx.ClassNames
	// Icon is the icon for the alert component.
	Icon htmx.Node
}

// Alert is a component that displays an alert.
func Alert(p AlertProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert": true,
			},
			p.ClassNames,
		),
		htmx.Role("alert"),
		htmx.IfElse(
			p.Icon != nil,
			p.Icon,
			heroicons.InformationCircleOutline(
				heroicons.IconProps{},
			),
		),
		htmx.Span(
			htmx.Group(children...),
		),
	)
}

// Info is a component that displays an info alert.
func Info(p AlertProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert":      true,
				"alert-info": true,
			},
			p.ClassNames,
		),
		htmx.Role("alert"),
		htmx.IfElse(
			p.Icon != nil,
			p.Icon,
			heroicons.InformationCircleOutline(
				heroicons.IconProps{},
			),
		),
		htmx.Span(
			htmx.Group(children...),
		),
	)
}

// Success is a component that displays a success alert.
func Success(p AlertProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert":         true,
				"alert-success": true,
			},
			p.ClassNames,
		),
		htmx.Role("alert"),
		htmx.SVG(
			htmx.ClassNames{
				"stroke-success": true,
				"shrink-0":       true,
				"w-6":            true,
				"h-6":            true,
			},
			htmx.SVG(
				htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
				htmx.Attribute("fill", "none"),
				htmx.Attribute("viewBox", "0 0 24 24"),
				htmx.Path(
					htmx.Attribute("stroke-linecap", "round"),
					htmx.Attribute("stroke-linejoin", "round"),
					htmx.Attribute("stroke-width", "2"),
					htmx.Attribute("d", "M5 13l4 4L19 7"),
				),
			),
		),
		htmx.Span(
			htmx.Group(children...),
		),
	)
}

// Warning is a component that displays a warning alert.
func Warning(p AlertProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert":         true,
				"alert-warning": true,
			},
			p.ClassNames,
		),
		htmx.Role("alert"),
		htmx.SVG(
			htmx.ClassNames{
				"stroke-warning": true,
				"shrink-0":       true,
				"w-6":            true,
				"h-6":            true,
			},
			htmx.SVG(
				htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
				htmx.Attribute("fill", "none"),
				htmx.Attribute("viewBox", "0 0 24 24"),
				htmx.Path(
					htmx.Attribute("stroke-linecap", "round"),
					htmx.Attribute("stroke-linejoin", "round"),
					htmx.Attribute("stroke-width", "2"),
					htmx.Attribute("d", "M19 5L5 19"),
				),
			),
		),
		htmx.Span(
			htmx.Group(children...),
		),
	)
}

// Error is a component that displays an error alert.
func Error(p AlertProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert":       true,
				"alert-error": true,
			},
			p.ClassNames,
		),
		htmx.SVG(
			htmx.ClassNames{
				"stroke-error": true,
				"shrink-0":     true,
				"w-6":          true,
				"h-6":          true,
			},
			htmx.SVG(
				htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
				htmx.Attribute("fill", "none"),
				htmx.Attribute("viewBox", "0 0 24 24"),
				htmx.Path(
					htmx.Attribute("stroke-linecap", "round"),
					htmx.Attribute("stroke-linejoin", "round"),
					htmx.Attribute("stroke-width", "2"),
					htmx.Attribute("d", "M6 18L18 6M6 6l12 12"),
				),
			),
		),
		htmx.Span(
			htmx.Group(children...),
		),
	)
}
