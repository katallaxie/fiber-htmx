package forms

import (
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/pkg/utilx"
)

// TextInputProps represents the properties for a text input element.
type TextInputProps struct {
	ClassNames  htmx.ClassNames // The class names for the text input element.
	Disabled    bool            // Whether the text input element is disabled.
	Type        string          // The type of the text input element.
	Error       error           // The error message of the text input element.
	Icon        htmx.Node       // The icon of the text input element.
	Name        string          // The name of the text input element.
	Placeholder string          // The placeholder of the text input element.
	Value       string          // The value of the text input element.
}

// TextInput returns a text input element based on the provided properties.
func TextInput(p TextInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"input":       true,
				"input-error": utilx.NotEmpty(p.Error),
				"w-full":      true,
			},
			p.ClassNames,
		),
		htmx.IfElse(
			utilx.NotEmpty(p.Type),
			htmx.Attribute("type", p.Type),
			htmx.Attribute("type", "text"),
		),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.If(utilx.NotEmpty(p.Error), htmx.Attribute("aria-invalid", "true")),
		htmx.Attribute("placeholder", p.Placeholder),
		htmx.Group(children...),
	)
}

// TextInputBordered is a component that displays a bordered text input.
func TextInputBordered(p TextInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"input":          true,
				"input-bordered": true,
				"w-full":         true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "text"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Attribute("placeholder", p.Placeholder),
		htmx.Group(children...),
	)
}

// TextInputGhost is a component that displays a ghost text input.
func TextInputGhost(p TextInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"input":       true,
				"input-ghost": true,
				"w-full":      true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "text"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Attribute("placeholder", p.Placeholder),
		htmx.Group(children...),
	)
}

// TextInputPrimary is a component that displays a primary text input.
func TextInputPrimary(p TextInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"input":          true,
				"input-bordered": true,
				"input-primary":  true,
				"w-full":         true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "text"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Attribute("placeholder", p.Placeholder),
		htmx.Group(children...),
	)
}

// TextInputSecondary is a component that displays a secondary text input.
func TextInputSecondary(p TextInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"input":           true,
				"input-bordered":  true,
				"input-secondary": true,
				"w-full":          true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "text"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Attribute("placeholder", p.Placeholder),
		htmx.Group(children...),
	)
}

// TextInputAccent is a component that displays an accent text input.
func TextInputAccent(p TextInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"input":          true,
				"input-bordered": true,
				"input-accent":   true,
				"w-full":         true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "text"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Attribute("placeholder", p.Placeholder),
		htmx.Group(children...),
	)
}

// TextInputError is a component that displays an error text input.
func TextInputError(p TextInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"input":          true,
				"input-bordered": true,
				"input-error":    true,
				"w-full":         true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "text"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Attribute("placeholder", p.Placeholder),
		htmx.Group(children...),
	)
}

// TextInputSuccess is a component that displays a success text input.
func TextInputSuccess(p TextInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"input":          true,
				"input-bordered": true,
				"input-success":  true,
				"w-full":         true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "text"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Attribute("placeholder", p.Placeholder),
		htmx.Group(children...),
	)
}

// TextInputWarning is a component that displays a warning text input.
func TextInputWarning(p TextInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"input":          true,
				"input-bordered": true,
				"input-warning":  true,
				"w-full":         true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "text"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Attribute("placeholder", p.Placeholder),
		htmx.Group(children...),
	)
}

// TextInputWithIcon is a component that displays a text input with an icon.
func TextInputWithIcon(p TextInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"input":          true,
				"input-bordered": true,
				"flex":           true,
				"items-center":   true,
				"gap-2":          true,
			},
			p.ClassNames,
		),
		htmx.Input(
			htmx.Attribute("type", "text"),
			htmx.Attribute("name", p.Name),
			htmx.Attribute("value", p.Value),
			htmx.Merge(
				htmx.ClassNames{
					"grow": true,
				},
			),
			htmx.If(p.Disabled, htmx.Disabled()),
			htmx.Attribute("placeholder", p.Placeholder),
		),
		htmx.If(p.Icon != nil, p.Icon),
		htmx.Group(children...),
	)
}
