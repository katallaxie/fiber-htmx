// Code generated by herogen; DO NOT EDIT.
package heroicons

import htmx "github.com/katallaxie/fiber-htmx"

func ArrowLongDownOutline(p IconProps) htmx.Node {
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
			htmx.Attribute("d", "M15.75 17.25L12 21M12 21L8.25 17.25M12 21L12 3"),
		),
	)
}
