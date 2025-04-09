package csrf

import htmx "github.com/katallaxie/fiber-htmx"

const defaultTokenName = "CSRFToken"

// CsrfTokenProps is the struct that holds the CSRF properties
type CsrfTokenProps struct {
	// Token is the CSRF token
	Token string
	// Name is the name of the CSRF token
	Name string
}

// CsrfToken is the struct that holds the CSRF properties
func CsrfToken(props CsrfTokenProps) htmx.Node {
	if props.Name == "" {
		props.Name = defaultTokenName
	}

	return htmx.Input(
		htmx.Type("hidden"),
		htmx.Name(props.Name),
		htmx.Value(props.Token),
	)
}
