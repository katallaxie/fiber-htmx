package forms

import htmx "github.com/katallaxie/fiber-htmx"

// FormProps represents the properties for a form element.
type FormProps struct {
	ClassNames htmx.ClassNames // The class names for the form element.
}

// Form returns a form element based on the provided properties.
// <form class="peer h-5 w-8"></form>.
func Form(p FormProps, children ...htmx.Node) htmx.Node {
	return htmx.Form(
		htmx.Merge(
			htmx.ClassNames{
				"form":  true,
				"group": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
