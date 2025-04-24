package utils

import (
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/pkg/slices"
)

// PageProps is the props for the Page component.
type PageProps struct {
	htmx.HTML5Props
}

// Page is a component that renders a page with a title and a body.
func Page(props PageProps, children ...htmx.Node) htmx.Node {
	return htmx.HTML5(
		htmx.HTML5Props{
			Title:       props.Title,
			Description: props.Description,
			Language:    props.Language,
			Attributes:  props.Attributes,
			Head: slices.Append([]htmx.Node{
				htmx.Link(
					htmx.Href("https://cdn.jsdelivr.net/npm/daisyui@5"),
					htmx.Rel("stylesheet"),
					htmx.Type("text/css"),
				),
				htmx.Script(
					htmx.Src("https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"),
				),
			}, props.Head...),
		},
	)
}
