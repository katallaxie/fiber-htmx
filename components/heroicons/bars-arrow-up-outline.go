// Code generated by herogen; DO NOT EDIT.
package heroicons

import "github.com/katallaxie/fiber-htmx"

func BarsArrowUpOutline(p IconProps) htmx.Node {
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
			htmx.Attribute("d", "M3 4.5H17.25M3 9H12.75M3 13.5H8.25M13.5 12.75L17.25 9M17.25 9L21 12.75M17.25 9V21"),
		),
	)
}
