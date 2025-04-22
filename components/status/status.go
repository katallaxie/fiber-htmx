package status

import htmx "github.com/katallaxie/fiber-htmx"

// StatusProps is a struct that contains the properties of the Status component.
type StatusProps struct {
	// AriaLabel is the aria-label attribute for the status element.
	AriaLabel string
	// AnimateBounce is a boolean that indicates whether to apply the animate-bounce class.
	AnimateBounce bool

	htmx.ClassNames
}

// Status is a component that renders a status indicator.
func Status(props StatusProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"animate-bounce": props.AnimateBounce,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// StatusPrimary is a component that renders a primary status indicator.
func StatusPrimary(props StatusProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"status-primary": true,
				"animate-bounce": props.AnimateBounce,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// StatusSecondary is a component that renders a secondary status indicator.
func StatusSecondary(props StatusProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":           true,
				"status-secondary": true,
				"animate-bounce":   props.AnimateBounce,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// StatusAccent is a component that renders an accent status indicator.
func StatusAccent(props StatusProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"status-accent":  true,
				"animate-bounce": props.AnimateBounce,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// StatusInfo is a component that renders an info status indicator.
func StatusInfo(props StatusProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"status-info":    true,
				"animate-bounce": props.AnimateBounce,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// StatusSuccess is a component that renders a success status indicator.
func StatusSuccess(props StatusProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"status-success": true,
				"animate-bounce": props.AnimateBounce,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// StatusWarning is a component that renders a warning status indicator.
func StatusWarning(props StatusProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"status-warning": true,
				"animate-bounce": props.AnimateBounce,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// StatusError is a component that renders an error status indicator.
func StatusError(props StatusProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"status-error":   true,
				"animate-bounce": props.AnimateBounce,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// StatusNeutral is a component that renders a neutral status indicator.
func StatusNeutral(props StatusProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"status-neutral": true,
				"animate-bounce": props.AnimateBounce,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}
