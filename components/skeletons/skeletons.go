package skeletons

import (
	"fmt"

	htmx "github.com/katallaxie/fiber-htmx"
)

// SkeletonProps is a struct that contains the props of the skeleton component.
type SkeletonProps struct {
	Width  int
	Height int

	htmx.ClassNames
}

// Skeleton is a component that renders a skeleton element.
func Skeleton(p SkeletonProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"skeleton":                    true,
				fmt.Sprintf("w-%d", p.Width):  true,
				fmt.Sprintf("h-%d", p.Height): true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
