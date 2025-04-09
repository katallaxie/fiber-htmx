package tabs

import (
	htmx "github.com/katallaxie/fiber-htmx"
)

// TabsProps is a struct that contains the properties of the Tabs component
type TabsProps struct {
	ID         string
	ClassNames htmx.ClassNames
}

// Tabs is a component that renders a list of tabs
func Tabs(props TabsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tabs": true,
			},
			props.ClassNames,
		),
		htmx.Role("tablist"),
		htmx.Group(children...),
	)
}

// TabsBoxed is a component that renders a list of tabs in a box
func TabsBoxed(props TabsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tabs":       true,
				"tabs-boxed": true,
			},
			props.ClassNames,
		),
		htmx.Role("tablist"),
		htmx.Group(children...),
	)
}

// TabProps is a struct that contains the properties of the Tab component
type TabProps struct {
	ClassNames htmx.ClassNames
	Name       string
	Label      string
	Active     bool
	Disabled   bool
}

// Tab is a component that renders a tab.
func Tab(props TabProps, children ...htmx.Node) htmx.Node {
	return htmx.Fragment(
		htmx.Input(
			htmx.Type("radio"),
			htmx.Name(props.Name),
			htmx.Role("tab"),
			htmx.Class("tab"),
			htmx.Aria("label", props.Label),
			htmx.If(props.Disabled, htmx.Disabled()),
			htmx.If(props.Active, htmx.Checked()),
		),
		htmx.Div(
			htmx.Merge(
				htmx.ClassNames{
					"tab-content": true,
					"p-10":        true,
				},
				props.ClassNames,
			),
			htmx.Role("tabpanel"),
			htmx.Group(children...),
		),
	)
}
