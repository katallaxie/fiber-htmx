package kbds

import htmx "github.com/katallaxie/fiber-htmx"

// KbdProps contains the properties for the kbds component.
type KbdProps struct {
	ClassNames htmx.ClassNames
}

// Kbds is a component for the htmx kbds extension.
func Kbd(p KbdProps, children ...htmx.Node) htmx.Node {
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

// KbdsExtraSmall is a component for the htmx kbds extension.
func KbdsExtraSmall(p KbdProps, children ...htmx.Node) htmx.Node {
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

// KbdsSmall is a component for the htmx kbds extension.
func KbdsSmall(p KbdProps, children ...htmx.Node) htmx.Node {
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

// KbdsMedium is a component for the htmx kbds extension.
func KbdsMedium(p KbdProps, children ...htmx.Node) htmx.Node {
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

// KbdsLarge is a component for the htmx kbds extension.
func KbdsLarge(p KbdProps, children ...htmx.Node) htmx.Node {
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

// KbdsExtraLarge is a component for the htmx kbds extension.
func KbdsExtraLarge(p KbdProps, children ...htmx.Node) htmx.Node {
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
