package accordions

import htmx "github.com/katallaxie/fiber-htmx"

// AccordionProps is a component that can be expanded and collapsed.
type AccordionProps struct {
	ClassNames htmx.ClassNames
	Name       string
}

// Accordion is a component that can be expanded and collapsed.
func Accordion(props AccordionProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse":    true,
				"bg-base-200": true,
			},
			props.ClassNames,
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
	ClassNames htmx.ClassNames
}

// AccordionTitle is a component that can be expanded and collapsed.
func AccordionTitle(props AccordionTitleProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse-title": true,
				"title-md":       true,
				"font-medium":    true,
			},
		),
		htmx.Group(children...),
	)
}

// AccordionContentProps is a component that can be expanded and collapsed.
type AccordionContentProps struct {
	ClassNames htmx.ClassNames
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
	ClassNames htmx.ClassNames
	Checked    bool
	Name       string
}

// AccordionRadio is a component that can be expanded and collapsed.
func AccordionRadio(props AccordionRadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("radio"),
		htmx.Attribute(
			"name",
			props.Name,
		),
	)
}
