package collapsible

import (
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/pkg/conv"
)

// CollapseProps is a component that can be expanded and collapsed.
type CollapseProps struct {
	ClassNames htmx.ClassNames
	TabIndex   int
}

// Collapse is a component that can be expanded and collapsed.
func Collapse(props CollapseProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse":    true,
				"bg-base-200": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("tabindex", conv.String(props.TabIndex)),
		htmx.Group(children...),
	)
}

// CollapseArrow is a component that can be expanded and collapsed.
func CollapseArrow(p CollapseProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Attribute("tabindex", conv.String(p.TabIndex)),
		htmx.Merge(
			htmx.ClassNames{
				"collapse":        true,
				"collapse-arrow":  true,
				"bg-base-200":     true,
				"border":          true,
				"border-base-300": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// CollapseTitleProps is a component that can be expanded and collapsed.
type CollapseTitleProps struct {
	ClassNames htmx.ClassNames
}

// CollapseTitle is a component that can be expanded and collapsed.
func CollapseTitle(p CollapseTitleProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse-title": true,
				"title-md":       true,
				"font-medium":    true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// CollapseContentProps is a component that can be expanded and collapsed.
type CollapseContentProps struct {
	ClassNames htmx.ClassNames
}

// CollapseContent is a component that can be expanded and collapsed.
func CollapseContent(props CollapseContentProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse-content": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// CollapseCheckboxProps is a component that can be expanded and collapsed.
type CollapseCheckboxProps struct {
	Checked bool

	htmx.ClassNames
}

// CollapseCheckbox is a component that can be expanded and collapsed.
func CollapseCheckbox(props CollapseCheckboxProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(props.ClassNames),
		htmx.Type("checkbox"),
		htmx.Group(children...),
	)
}

// CollapseRadioProps is a component that can be expanded and collapsed.
type CollapseRadioProps struct {
	Checked bool
	Name    string

	htmx.ClassNames
}

// CollapseRadio is a component that can be expanded and collapsed.
func CollapseRadio(props CollapseRadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("radio"),
		htmx.If(props.Checked, htmx.Checked()),
		htmx.Name(props.Name),
		htmx.Group(children...),
	)
}
