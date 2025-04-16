package accordions

import htmx "github.com/katallaxie/fiber-htmx"

// AccordionProps is a component that can be expanded and collapsed.
type AccordionProps struct {
	ClassNames htmx.ClassNames
	Name       string
	Checked    bool
}

// Accordion is a component that can be expanded and collapsed.
func Accordion(props AccordionProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"bg-base-200":     true,
				"border-base-300": true,
				"border":          true,
				"collapse":        true,
			},
			props.ClassNames,
		),
		AccordionRadio(
			AccordionRadioProps{
				Name:    props.Name,
				Checked: props.Checked,
			},
		),
		htmx.Group(children...),
	)
}

// AccordionArrow is a component that can be expanded and collapsed.
func AccordionArrow(props AccordionProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse":       true,
				"collapse-arrow": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// AccordionTitleProps is a component title.
type AccordionTitleProps struct {
	htmx.ClassNames
}

// AccordionTitle is a component that can be expanded and collapsed.
func AccordionTitle(props AccordionTitleProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse-title": true,
				"font-semibold":  true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// AccordionContentProps is a component that can be expanded and collapsed.
type AccordionContentProps struct {
	htmx.ClassNames
}

// AccordionContent is a component that can be expanded and collapsed.
func AccordionContent(props AccordionContentProps, children ...htmx.Node) htmx.Node {
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

// AccordionRadioProps is a component that can be expanded and collapsed.
type AccordionRadioProps struct {
	// Name is the name of the radio button.
	Name string
	// Checked is the checked state of the radio button.
	Checked bool

	htmx.ClassNames
}

// AccordionRadio is a component that can be expanded and collapsed.
func AccordionRadio(props AccordionRadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("radio"),
		htmx.Name(props.Name),
		htmx.If(props.Checked, htmx.Checked()),
	)
}
