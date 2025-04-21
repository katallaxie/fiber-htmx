package forms

import (
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/pkg/utilx"
)

// FilterFormProps is a struct that contains the properties of the Filter component.
type FilterFormProps struct {
	htmx.ClassNames
}

// FilterForm is a component that renders a filter element.
func FilterForm(props FilterFormProps, children ...htmx.Node) htmx.Node {
	return htmx.Form(
		htmx.Merge(
			htmx.ClassNames{
				"filter": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FilterProps is a struct that contains the properties of the Filter component.
type FilterProps struct {
	htmx.ClassNames
}

// Filter is a component that renders a filter element.
func Filter(props FilterFormProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"filter": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FilterResetButtonProps is a struct that contains the properties of the FilterReset component.
type FilterResetButtonProps struct {
	Value string
	htmx.ClassNames
}

// FilterResetButton is a component that renders a filter reset button.
func FilterResetButton(props FilterResetButtonProps, _ ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"btn":        true,
				"btn-square": true,
			},
			props.ClassNames,
		),
		htmx.Type("reset"),
		utilx.IfElse(utilx.NotEmpty(props.Value), htmx.Value(props.Value), htmx.Value("×")),
	)
}

// FilterButton is a component that renders a filter button.
type FilterButtonProps struct {
	Name  string
	Label string
	htmx.ClassNames
}

// FilterButton is a component that renders a filter button.
func FilterButton(props FilterButtonProps, _ ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"btn": true,
			},
			props.ClassNames,
		),
		htmx.Type("radio"),
		htmx.Name(props.Name),
		htmx.Attribute("aria-label", props.Label),
	)
}
