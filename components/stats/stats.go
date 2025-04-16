package stats

import htmx "github.com/katallaxie/fiber-htmx"

// Props contains the properties for the stats component.
type Props struct {
	htmx.ClassNames
}

// Stats is a component for the htmx stats extension.
//
//nolint:revive
func Stats(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stats":  true,
				"shadow": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Vertical is a component for the htmx stats extension.
func Vertical(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stats":          true,
				"shadow":         true,
				"stats-vertical": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// StatProps contains the properties for the stat component.
type StatProps struct {
	ClassNames htmx.ClassNames
}

// Stat is a component for the htmx stats extension.
func Stat(p StatProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stat": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// TitleProps contains the properties for the title component.
type TitleProps struct {
	ClassNames htmx.ClassNames
}

// Title is a component for the htmx stats extension.
func Title(p TitleProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stat-title": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ValueProps contains the properties for the value component.
type ValueProps struct {
	ClassNames htmx.ClassNames
}

// Value is a component for the htmx stats extension.
func Value(p ValueProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stat-value": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DescriptionProps contains the properties for the description component.
type DescriptionProps struct {
	ClassNames htmx.ClassNames
}

// Description is a component for the htmx stats extension.
func Description(p DescriptionProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stat-desc": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FigureProps contains the properties for the figure component.
type FigureProps struct {
	ClassNames htmx.ClassNames
}

// Figure is a component for the htmx stats extension.
func Figure(p FigureProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stat-figure": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ActiosnProps contains the properties for the action component.
type ActionsProps struct {
	ClassNames htmx.ClassNames
}

// Actions is a component for the htmx stats extension.
func Actions(p ActionsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"stat-actions": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
