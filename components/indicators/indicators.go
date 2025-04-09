package indicators

import htmx "github.com/katallaxie/fiber-htmx"

// IndicatorProps is a struct that contains the properties of an indicator.
type IndicatorProps struct {
	// ClassNames is a map of class names.
	ClassNames htmx.ClassNames
}

// Indicator is a function that returns an indicator.
func Indicator(props IndicatorProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"indicator": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// IndiciatorItemProps is a struct that contains the properties of an indicator item.
type IndicatorItemProps struct {
	// ClassNames is a map of class names.
	ClassNames htmx.ClassNames
}

// IndicatorItem is a function that returns an indicator item.
func IndicatorItem(props IndicatorItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"indicator-item": true,
				"badge":          true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// IndicatorItemPrimary is a function that returns a primary indicator item.
func IndicatorItemPrimary(props IndicatorItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"indicator-item": true,
				"badge":          true,
				"badge-primary":  true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// IndicatorItemSecondary is a function that returns a secondary indicator item.
func IndicatorItemSecondary(props IndicatorItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"indicator-item":  true,
				"badge":           true,
				"badge-secondary": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
