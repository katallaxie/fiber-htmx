package htmx

// Path creates an SVG path element with the specified children.
// It is a convenience function that calls the Element function with "path" as the tag name.
func Path(children ...Node) Node {
	return Element("path", children...)
}

// G creates an SVG group element with the
// specified children. It is a convenience function that calls the Element function
// with "g" as the tag name.
func G(children ...Node) Node {
	return Element("g", children...)
}

// SVG creates an SVG element with the specified children.
// It sets the "xmlns" attribute to "http://www.w3.org/2000/svg".
// The children are rendered inside a group element.
func SVG(children ...Node) Node {
	return Element("svg", Attribute("xmlns", "http://www.w3.org/2000/svg"), Group(children...))
}

// ClipRule returns an SVG attribute node for specifying the clip rule.
// The clip rule determines how the clipping path is applied to the SVG element.
// The value of the clip rule is specified as a string.
// Example usage: ClipRule("evenodd").
func ClipRule(v string) Node {
	return Attribute("clip-rule", v)
}

// D returns an SVG attribute node for specifying the path data.
// The path data defines the shape of the SVG path element.
// The value of the path data is specified as a string.
// Example usage: D("M10 10 L20 20").
func D(v string) Node {
	return Attribute("d", v)
}

// Fill returns an SVG attribute node for specifying the fill color.
// The fill color determines the color used to fill the SVG element.
// The value of the fill color is specified as a string.
// Example usage: Fill("red").
func Fill(v string) Node {
	return Attribute("fill", v)
}

// FillRule returns an SVG attribute node for specifying the fill rule.
// The fill rule determines how the interior of the SVG element is filled.
// The value of the fill rule is specified as a string.
// Example usage: FillRule("evenodd").
func FillRule(v string) Node {
	return Attribute("fill-rule", v)
}

// Stroke returns an SVG attribute node for specifying the stroke color.
// The stroke color determines the color used to stroke the SVG element.
// The value of the stroke color is specified as a string.
// Example usage: Stroke("blue").
func Stroke(v string) Node {
	return Attribute("stroke", v)
}

// StrokeWidth returns an SVG attribute node for specifying the stroke width.
// The stroke width determines the thickness of the stroke applied to an SVG element.
// The value of the stroke width is specified as a string.
// Example usage: StrokeWidth("2px").
func StrokeWidth(v string) Node {
	return Attribute("stroke-width", v)
}

// ViewBox returns an SVG attribute node for specifying the viewBox.
// The viewBox defines the position and size of the SVG element's viewport.
// The value of the viewBox is specified as a string.
// Example usage: ViewBox("0 0 100 100").
func ViewBox(v string) Node {
	return Attribute("viewBox", v)
}
