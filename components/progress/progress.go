package progress

import (
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/pkg/conv"
)

// Props is a struct that contains the props of the  component.
type Props struct {
	ClassNames htmx.ClassNames
	Value      int
	Max        int
}

// Progress is a component that renders a  element.
func Progress(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Progress(
		htmx.Merge(
			htmx.ClassNames{
				"progress": true,
				"w-56":     true,
			},
			p.ClassNames,
		),
		htmx.Attribute("max", conv.String(p.Max)),
		htmx.Attribute("value", conv.String(p.Value)),
		htmx.Group(children...),
	)
}

// Primary is a component that renders a primary element.
func Primary(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Progress(
		htmx.Merge(
			htmx.ClassNames{
				"progress":   true,
				"w-56":       true,
				"bg-primary": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("max", conv.String(p.Max)),
		htmx.Attribute("value", conv.String(p.Value)),
		htmx.Group(children...),
	)
}

// Secondary is a component that renders a secondary element.
func Secondary(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Progress(
		htmx.Merge(
			htmx.ClassNames{
				"progress":           true,
				"w-56":               true,
				"progress-secondary": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("max", conv.String(p.Max)),
		htmx.Attribute("value", conv.String(p.Value)),
		htmx.Group(children...),
	)
}

// Success is a component that renders a success element.
func Success(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Progress(
		htmx.Merge(
			htmx.ClassNames{
				"progress":         true,
				"w-56":             true,
				"progress-success": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("max", conv.String(p.Max)),
		htmx.Attribute("value", conv.String(p.Value)),
		htmx.Group(children...),
	)
}

// Info is a component that renders a info element.
func Info(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Progress(
		htmx.Merge(
			htmx.ClassNames{
				"progress":      true,
				"w-56":          true,
				"progress-info": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("max", conv.String(p.Max)),
		htmx.Attribute("value", conv.String(p.Value)),
		htmx.Group(children...),
	)
}

// Warning is a component that renders a warning element.
func Warning(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Progress(
		htmx.Merge(
			htmx.ClassNames{
				"progress":         true,
				"w-56":             true,
				"progress-warning": true,
			},
		),
		htmx.Attribute("max", conv.String(p.Max)),
		htmx.Attribute("value", conv.String(p.Value)),
		htmx.Group(children...),
	)
}

// Intermediate is a component that renders a intermediate element.
func Intermediate(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Progress(
		htmx.Merge(
			htmx.ClassNames{
				"progress": true,
				"w-56":     true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
