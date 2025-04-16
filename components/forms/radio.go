package forms

import (
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/pkg/utilx"
)

// RadioProps represents the properties for a radio element.
type RadioProps struct {
	Name     string // The name of the radio element.
	Value    string // The value of the radio element.
	Checked  bool   // Whether the radio element is checked.
	Disabled bool
	Error    error

	htmx.ClassNames // The class names for the radio element.
}

// Radio generates a radio element based on the provided properties.
func Radio(p RadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"radio":       true,
				"radio-error": utilx.NotEmpty(p.Error),
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "radio"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// RadioSuccess component represents a successful radio element.
func RadioSuccess(p RadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"radio":         true,
				"radio-success": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "radio"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// RadioInfo component represents a info radio element.
func RadioInfo(p RadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"radio":      true,
				"radio-info": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "radio"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// RadioWarning component represents a warning radio element.
func RadioWarning(p RadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"radio":         true,
				"radio-warning": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "radio"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// RadioError component represents an error radio element.
func RadioError(p RadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"radio":       true,
				"radio-error": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "radio"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// RadioPrimary component represents a primary radio element.
func RadioPrimary(p RadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"radio":         true,
				"radio-primary": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "radio"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// RadioSecondary component represents a secondary radio element.
func RadioSecondary(p RadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"radio":           true,
				"radio-secondary": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "radio"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// RadioAccent component represents an accent radio element.
func RadioAccent(p RadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"radio":        true,
				"radio-accent": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "radio"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}
