// Code generated by herogen; DO NOT EDIT.
package heroicons

import "github.com/katallaxie/fiber-htmx"

func MoonOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Merge(
			htmx.ClassNames{
				"h-5": true,
				"h-6": false,
				"w-5": true,
				"w-6": false,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M21.7519 15.0021C20.597 15.484 19.3296 15.7501 18 15.7501C12.6152 15.7501 8.25 11.3849 8.25 6.00011C8.25 4.67052 8.51614 3.40308 8.99806 2.24817C5.47566 3.71798 3 7.19493 3 11.2501C3 16.6349 7.36522 21.0001 12.75 21.0001C16.8052 21.0001 20.2821 18.5245 21.7519 15.0021Z"),
		),
	)
}
