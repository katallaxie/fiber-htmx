package htmx

import (
	"io"
	"sort"
	"strings"
)

// HTML5Props represents the properties for an HTML5 document.
type HTML5Props struct {
	// Title is the title of the HTML document.
	Title string
	// Description is the description of the HTML document.
	Description string
	// Language is the language of the HTML document.
	Language string
	// Head is the head section of the HTML document.
	Head []Node
	// Attributes are the attributes to be included in the HTML document.
	Attributes []Node
}

// HTML5 generates an HTML5 document based on the provided properties.
func HTML5(props HTML5Props, body ...Node) Node {
	return Doctype(
		HTML(
			If(props.Language != "", Lang(props.Language)),
			Group(props.Attributes...),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				TitleElement(Text(props.Title)),
				If(props.Description != "", Meta(Name("description"), Content(props.Description))),
				Group(props.Head...),
			),
			Body(body...),
		),
	)
}

// HxClassName represents a class name for htmx elements.
type HxClassName string

// String returns the string representation of the HxClassName.
func (c HxClassName) String() string {
	return string(c)
}

const (
	HxClassNameAdded     HxClassName = "htmx-added"
	HxClassNameIndicator HxClassName = "htmx-indicator"
	HxClassNameRequest   HxClassName = "htmx-request"
	HxClassNameSettling  HxClassName = "htmx-settling"
	HxClassNameSwapping  HxClassName = "htmx-swapping"
)

const (
	HxClssNameAddedSelector     = ".htmx-added"
	HxClssNameIndicatorSelector = ".htmx-indicator"
	HxClssNameRequestSelector   = ".htmx-request"
	HxClssNameSettlingSelector  = ".htmx-settling"
	HxClssNameSwappingSelector  = ".htmx-swapping"
)

// ClassNames represents a set of class names.
type ClassNames map[string]bool

// Render writes the class names to the provided writer.
func (c ClassNames) Render(w io.Writer) error {
	classes := make([]string, 0, len(c))

	for class, ok := range c {
		if !ok {
			continue
		}
		classes = append(classes, strings.TrimSpace(class))
	}

	sort.Strings(classes)

	return Class(strings.Join(classes, " ")).Render(w)
}

// Type returns the node type of the ClassNames.
func (c ClassNames) Type() NodeType {
	return AttributeType
}

// String returns the string representation of the ClassNames.
func (c ClassNames) String() string {
	var b strings.Builder

	_ = c.Render(&b)

	return b.String()
}
