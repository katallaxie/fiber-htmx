package forms

import htmx "github.com/katallaxie/fiber-htmx"

// ValidatorHintProps is a struct that contains the properties of the ValidatorHint component.
type ValidatorHintProps struct {
	htmx.ClassNames
}

// ValidatorHint is a component that renders a hint for a validator.
func ValidatorHint(props ValidatorHintProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"validator-hint": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
