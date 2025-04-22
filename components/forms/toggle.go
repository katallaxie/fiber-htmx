package forms

import htmx "github.com/katallaxie/fiber-htmx"

// ToggleProps represents the properties for a toggle element.
type ToggleProps struct {
	ClassNames htmx.ClassNames // The class names for the toggle element.
	Name       string          // The name of the toggle element.
	Value      string          // The value of the toggle element.
	Disabled   bool            // Whether the toggle element is disabled.
	Checked    bool            // Whether the toggle element is checked.
	Validator  bool            // Whether the toggle element is a validator.
}

// Toggle returns a toggle element based on the provided properties.
func Toggle(p ToggleProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"toggle":    true,
				"validator": p.Validator,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "checkbox"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.Group(children...),
	)
}

// ToggleSuccess is a component that displays a success toggle.
func ToggleSuccess(p ToggleProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"toggle":         true,
				"toggle-success": true,
				"validator":      p.Validator,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "checkbox"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.Group(children...),
	)
}

// ToggleWarning is a component that displays a warning toggle.
func ToggleWarning(p ToggleProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"toggle":         true,
				"toggle-warning": true,
				"validator":      p.Validator,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "checkbox"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.Group(children...),
	)
}

// ToggleInfo is a component that displays an info toggle.
func ToggleInfo(p ToggleProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"toggle":      true,
				"toggle-info": true,
				"validator":   p.Validator,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "checkbox"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.Group(children...),
	)
}

// ToggleError is a component that displays an error toggle.
func ToggleError(p ToggleProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"toggle":       true,
				"toggle-error": true,
				"validator":    p.Validator,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "checkbox"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.Group(children...),
	)
}

// ToggleLabel is a component that displays a label for a toggle.
func ToggleLabel(p ToggleProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"form-control": true,
			},
		),
		htmx.Label(
			htmx.Merge(
				htmx.ClassNames{
					"label":          true,
					"cursor-pointer": true,
				},
			),
			htmx.Span(
				htmx.ClassNames{
					"label-text": true,
				},
				htmx.Group(children...),
			),
			htmx.Input(
				htmx.Merge(
					htmx.ClassNames{
						"toggle": true,
					},
					p.ClassNames,
				),
				htmx.Attribute("type", "checkbox"),
				htmx.Attribute("name", p.Name),
				htmx.Attribute("value", p.Value),
				htmx.If(p.Disabled, htmx.Disabled()),
				htmx.If(p.Checked, htmx.Checked()),
			),
		),
	)
}
