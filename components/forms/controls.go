package forms

import htmx "github.com/katallaxie/fiber-htmx"

// FormControlProps represents the properties for a form control element.
type FormControlProps struct {
	ClassNames htmx.ClassNames // The class names for the form control element.
}

// FormControl generates a form control element based on the provided properties.
func FormControl(p FormControlProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"form-control": true,
				"w-full":       true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FormControlLabelProps represents the properties for a form control label element.
type FormControlLabelProps struct {
	ClassNames htmx.ClassNames // The class names for the form control label element.
}

// FormControlLabel generates a form control label element based on the provided properties.
func FormControlLabel(props FormControlLabelProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"label":          true,
				"cursor-pointer": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FormControlLabelTextProps represents the properties for a form control label text element.
type FormControlLabelTextProps struct {
	ClassNames htmx.ClassNames // The class names for the form control label text element.
}

// FormControlLabelText generates a form control label text element based on the provided properties.
func FormControlLabelText(props FormControlLabelTextProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"label-text": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FormControlLabelAltTextProps represents the properties for a form control label alt text element.
type FormControlLabelAltTextProps struct {
	ClassNames htmx.ClassNames // The class names for the form control label alt text element.
}

// FormControlLabelAltText generates a form control label alt text element based on the provided properties.
func FormControlLabelAltText(props FormControlLabelAltTextProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"label-text-alt": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
