package swaps

import htmx "github.com/katallaxie/fiber-htmx"

// Props contains the properties for the swap component.
type Props struct {
	ClassNames htmx.ClassNames
	Value      string
}

// On is a component for the on state of the swap component.
func On(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"swap-on": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Off is a component for the off state of the swap component.
func Off(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"swap-off": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Swap is a component for the htmx swap extension.
func Swap(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"swap": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
