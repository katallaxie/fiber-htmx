// Code generated by herogen; DO NOT EDIT.
package heroicons

import htmx "github.com/katallaxie/fiber-htmx"

func SlashSolid(p IconProps) htmx.Node {
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
			htmx.Attribute("d", "M15.256 3.04243C15.6453 3.18399 15.8461 3.61434 15.7046 4.00364L9.7046 20.504C9.56304 20.8933 9.13271 21.0942 8.74342 20.9526C8.35414 20.811 8.15331 20.3807 8.29487 19.9914L14.2948 3.49099C14.4364 3.1017 14.8667 2.90087 15.256 3.04243Z"),
		),
	)
}
