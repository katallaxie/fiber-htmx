package progress

import (
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/pkg/conv"
)

// ProgressProps is a struct that contains the props of the progress component
type ProgressProps struct {
	ClassNames htmx.ClassNames
	Value      int
	Max        int
}

// Progress is a component that renders a progress element
func Progress(p ProgressProps, children ...htmx.Node) htmx.Node {
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

// ProgressPrimary is a component that renders a primary progress element
func ProgressPrimary(p ProgressProps, children ...htmx.Node) htmx.Node {
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

// ProgressSecondary is a component that renders a secondary progress element
func ProgressSecondary(p ProgressProps, children ...htmx.Node) htmx.Node {
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

// ProgressSuccess is a component that renders a success progress element
func ProgressSuccess(p ProgressProps, children ...htmx.Node) htmx.Node {
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

// ProgressInfo is a component that renders a info progress element
func ProgressInfo(p ProgressProps, children ...htmx.Node) htmx.Node {
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

// ProgressWarning is a component that renders a warning progress element
func ProgressWarning(p ProgressProps, children ...htmx.Node) htmx.Node {
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

// ProgressIntermediate is a component that renders a intermediate progress element
func ProgressIntermediate(p ProgressProps, children ...htmx.Node) htmx.Node {
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
