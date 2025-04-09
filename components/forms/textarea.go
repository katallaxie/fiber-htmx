package forms

import htmx "github.com/katallaxie/fiber-htmx"

// TextareaProps are the properties for the textarea component
type TextareaProps struct {
	ClassNames  htmx.ClassNames
	Placeholder string
	Disabled    bool
	Name        string
	Value       string
}

// Textarea is a textarea component
func Textarea(p TextareaProps, children ...htmx.Node) htmx.Node {
	return htmx.Textarea(
		htmx.Merge(
			htmx.ClassNames{
				"textarea": true,
			},
			p.ClassNames,
		),
		htmx.Attribute(
			"placeholder",
			p.Placeholder,
		),
		htmx.Attribute(
			"name",
			p.Name,
		),
		htmx.Attribute(
			"value",
			p.Value,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// TextareaBordered is a textarea component with a border
func TextareaBordered(p TextareaProps, children ...htmx.Node) htmx.Node {
	return htmx.Textarea(
		htmx.Merge(
			htmx.ClassNames{
				"textarea":          true,
				"textarea-bordered": true,
			},
			p.ClassNames,
		),
		htmx.Attribute(
			"placeholder",
			p.Placeholder,
		),
		htmx.Attribute(
			"name",
			p.Name,
		),
		htmx.Attribute(
			"value",
			p.Value,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// TextareaGhost is a textarea component with a ghost border
func TextareaGhost(p TextareaProps, children ...htmx.Node) htmx.Node {
	return htmx.Textarea(
		htmx.Merge(
			htmx.ClassNames{
				"textarea":       true,
				"textarea-ghost": true,
			},
			p.ClassNames,
		),
		htmx.Attribute(
			"placeholder",
			p.Placeholder,
		),
		htmx.Attribute(
			"name",
			p.Name,
		),
		htmx.Attribute(
			"value",
			p.Value,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// TextareaPrimary is a textarea component with a primary border
func TextareaPrimary(p TextareaProps, children ...htmx.Node) htmx.Node {
	return htmx.Textarea(
		htmx.Merge(
			htmx.ClassNames{
				"textarea":         true,
				"textarea-primary": true,
			},
			p.ClassNames,
		),
		htmx.Attribute(
			"placeholder",
			p.Placeholder,
		),
		htmx.Attribute(
			"name",
			p.Name,
		),
		htmx.Attribute(
			"value",
			p.Value,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// TextareaSecondary is a textarea component with a secondary border
func TextareaSecondary(p TextareaProps, children ...htmx.Node) htmx.Node {
	return htmx.Textarea(
		htmx.Merge(
			htmx.ClassNames{
				"textarea":           true,
				"textarea-secondary": true,
			},
			p.ClassNames,
		),
		htmx.Attribute(
			"placeholder",
			p.Placeholder,
		),
		htmx.Attribute(
			"name",
			p.Name,
		),
		htmx.Attribute(
			"value",
			p.Value,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// TextareaSuccess is a textarea component with a success border
func TextareaSuccess(p TextareaProps, children ...htmx.Node) htmx.Node {
	return htmx.Textarea(
		htmx.Merge(
			htmx.ClassNames{
				"textarea":         true,
				"textarea-success": true,
			},
			p.ClassNames,
		),
		htmx.Attribute(
			"placeholder",
			p.Placeholder,
		),
		htmx.Attribute(
			"name",
			p.Name,
		),
		htmx.Attribute(
			"value",
			p.Value,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// TextareaWarning is a textarea component with a warning border
func TextareaWarning(p TextareaProps, children ...htmx.Node) htmx.Node {
	return htmx.Textarea(
		htmx.Merge(
			htmx.ClassNames{
				"textarea":         true,
				"textarea-warning": true,
			},
			p.ClassNames,
		),
		htmx.Attribute(
			"placeholder",
			p.Placeholder,
		),
		htmx.Attribute(
			"name",
			p.Name,
		),
		htmx.Attribute(
			"value",
			p.Value,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// TextareaError is a textarea component with an error border
func TextareaError(p TextareaProps, children ...htmx.Node) htmx.Node {
	return htmx.Textarea(
		htmx.Merge(
			htmx.ClassNames{
				"textarea":       true,
				"textarea-error": true,
			},
			p.ClassNames,
		),
		htmx.Attribute(
			"placeholder",
			p.Placeholder,
		),
		htmx.Attribute(
			"name",
			p.Name,
		),
		htmx.Attribute(
			"value",
			p.Value,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}
