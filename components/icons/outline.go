// Package icons provides a set of icons that can be used in the application.
// The icons are implemented as functions that return htmx.Node.
//
// https://heroicons.com
package icons

import (
	htmx "github.com/zeiss/fiber-htmx"
)

// IconProps ...
type IconProps struct {
	ClassNames htmx.ClassNames
}

// ChevronUpDownOutline ...
func ChevronUpDown(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.ClassNames{
			"w-6 h-6": true,
		}.Merge(p.ClassNames),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M8.25 15 12 18.75 15.75 15m-7.5-6L12 5.25 15.75 9"),
		),
	)
}

// ChevronUpOutline ...
func ChevronUp(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.ClassNames{
			"w-6 h-6": true,
		}.Merge(p.ClassNames),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M19.5 15 12 7.5 4.5 15"),
		),
	)
}

// ChevronDownOutline ...
func ChevronDown(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.ClassNames{
			"w-6 h-6": true,
		}.Merge(p.ClassNames),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M19.5 8.25 12 15.75 4.5 8.25"),
		),
	)
}

// ChevronLeftOutline ...
func ChevronLeft(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.ClassNames{
			"w-6 h-6": true,
		}.Merge(p.ClassNames),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M15 19.5 7.5 12 15 4.5"),
		),
	)
}

// ChevronRightOutline ...
func ChevronRight(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.ClassNames{
			"w-6 h-6": true,
		}.Merge(p.ClassNames),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M9 19.5 16.5 12 9 4.5"),
		),
	)
}

// SearchOutline ...
func Search(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"),
		),
	)
}
