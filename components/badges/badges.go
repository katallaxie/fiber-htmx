package badges

import htmx "github.com/zeiss/fiber-htmx"

// BadgeProps represents the properties for a badge element.
type BadgeProps struct {
	ClassNames map[string]bool // The class names for the badge element.
}

// Badge generates a badge element based on the provided properties.
func Badge(p BadgeProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.ClassNames{
			"badge": true,
		}.Merge(p.ClassNames),
		htmx.Group(children...),
	)
}

// Neutral generates a neutral badge element based on the provided properties.
func Neutral(p BadgeProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.ClassNames{
			"badge":         true,
			"badge-neutral": true,
		}.Merge(p.ClassNames),
		htmx.Group(children...),
	)
}

// Primary generates a primary badge element based on the provided properties.
func Primary(p BadgeProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.ClassNames{
			"badge":         true,
			"badge-primary": true,
		}.Merge(p.ClassNames),
		htmx.Group(children...),
	)
}

// Secondary generates a secondary badge element based on the provided properties.
func Secondary(p BadgeProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.ClassNames{
			"badge":           true,
			"badge-secondary": true,
		}.Merge(p.ClassNames),
		htmx.Group(children...),
	)
}

// Accent generates an accent badge element based on the provided properties.
func Accent(p BadgeProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.ClassNames{
			"badge":        true,
			"badge-accent": true,
		}.Merge(p.ClassNames),
		htmx.Group(children...),
	)
}

// Ghost generates a ghost badge element based on the provided properties.
func Ghost(p BadgeProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.ClassNames{
			"badge":       true,
			"badge-ghost": true,
		}.Merge(p.ClassNames),
		htmx.Group(children...),
	)
}
