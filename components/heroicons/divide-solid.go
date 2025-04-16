// Code generated by herogen; DO NOT EDIT.
package heroicons

import htmx "github.com/katallaxie/fiber-htmx"

func DivideSolid(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Merge(
			htmx.ClassNames{
				"h-5": false,
				"h-6": true,
				"w-5": false,
				"w-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("d", "M10.8737 5.24812C10.8737 4.62682 11.3773 4.12305 11.9987 4.12305C12.6201 4.12305 13.1237 4.62682 13.1237 5.24812C13.1237 5.86942 12.6201 6.3732 11.9987 6.3732C11.3773 6.3732 10.8737 5.86942 10.8737 5.24812ZM3.74854 11.9983C3.74854 11.5841 4.08433 11.2483 4.49855 11.2483H19.4983C19.9125 11.2483 20.2483 11.5841 20.2483 11.9983C20.2483 12.4126 19.9125 12.7484 19.4983 12.7484H4.49855C4.08433 12.7484 3.74854 12.4126 3.74854 11.9983ZM10.8743 18.751C10.8743 18.1297 11.378 17.6259 11.9994 17.6259C12.6207 17.6259 13.1244 18.1297 13.1244 18.751C13.1244 19.3723 12.6207 19.8761 11.9994 19.8761C11.378 19.8761 10.8743 19.3723 10.8743 18.751Z"),
		),
	)
}
