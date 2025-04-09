package tooltips

import htmx "github.com/katallaxie/fiber-htmx"

// TooltipProps ...
type TooltipProps struct {
	ClassNames htmx.ClassNames
	DataTip    string
}

// Tooltip ...
func Tooltip(p TooltipProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("data-tip", p.DataTip),
		htmx.Group(children...),
	)
}

// TooltipPrimary ...
func TooltipPrimary(p TooltipProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":         true,
				"tooltip-primary": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("data-tip", p.DataTip),
		htmx.Group(children...),
	)
}

// TooltipSecondary ...
func TooltipSecondary(p TooltipProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":           true,
				"tooltip-secondary": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("data-tip", p.DataTip),
		htmx.Group(children...),
	)
}

// TooltipSuccess ...
func TooltipSuccess(p TooltipProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":         true,
				"tooltip-success": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("data-tip", p.DataTip),
		htmx.Group(children...),
	)
}

// TooltipWarning ...
func TooltipWarning(p TooltipProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":         true,
				"tooltip-warning": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("data-tip", p.DataTip),
		htmx.Group(children...),
	)
}

// TooltipInfo ...
func TooltipInfo(p TooltipProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":      true,
				"tooltip-info": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("data-tip", p.DataTip),
		htmx.Group(children...),
	)
}

// TooltipError ...
func TooltipError(p TooltipProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":       true,
				"tooltip-error": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("data-tip", p.DataTip),
		htmx.Group(children...),
	)
}

// TooltipAccent ...
func TooltipAccent(p TooltipProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":        true,
				"tooltip-accent": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("data-tip", p.DataTip),
		htmx.Group(children...),
	)
}
