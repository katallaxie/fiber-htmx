package kbds

import htmx "github.com/katallaxie/fiber-htmx"

// Props contains the properties for the kbds component.
type Props struct {
	htmx.ClassNames
}

// Kbd is a component for the htmx kbds extension.
func Kbd(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Kbd(
		htmx.Merge(
			htmx.ClassNames{
				"kbd": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ExtraSmall is a component for the htmx kbds extension.
func ExtraSmall(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Kbd(
		htmx.Merge(
			htmx.ClassNames{
				"kbd":    true,
				"kbd-xs": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Small is a component for the htmx kbds extension.
func Small(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Kbd(
		htmx.Merge(
			htmx.ClassNames{
				"kbd":    true,
				"kbd-sm": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Medium is a component for the htmx kbds extension.
func Medium(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Kbd(
		htmx.Merge(
			htmx.ClassNames{
				"kbd":    true,
				"kbd-md": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Large is a component for the htmx kbds extension.
func Large(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Kbd(
		htmx.Merge(
			htmx.ClassNames{
				"kbd":    true,
				"kbd-lg": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ExtraLarge is a component for the htmx kbds extension.
func ExtraLarge(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Kbd(
		htmx.Merge(
			htmx.ClassNames{
				"kbd":    true,
				"kbd-xl": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
