package breadcrumbs

import htmx "github.com/katallaxie/fiber-htmx"

// BreadcrumsbProps represents the properties for a breadcrumb element.
type BreadcrumbsProps struct {
	ClassNames htmx.ClassNames // The class names for the breadcrumb element.
}

// BreadCrumbs generates a breadcrumb element based on the provided properties.
func Breadcrumbs(p BreadcrumbsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"breadcrumbs": true,
				"text-sm":     true,
			},
			p.ClassNames,
		),
		htmx.Ul(
			htmx.Group(children...),
		),
	)
}

// BreadcrumbProps represents the properties for a breadcrumb item element.
type BreadcrumbProps struct {
	ClassNames htmx.ClassNames // The class names for the breadcrumb item element.
	Href       string          // The URL of the linked document.
	Rel        string          // The relationship between the current document and the linked document.
	Title      string          // The title of the linked document.
}

// BreadCrumb generates a breadcrumb item element based on the provided properties.
func Breadcrumb(p BreadcrumbProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			htmx.ClassNames{},
			p.ClassNames,
		),
		htmx.A(
			htmx.Href(p.Href),
			htmx.Rel(p.Rel),
			htmx.Text(p.Title),
		),
	)
}
