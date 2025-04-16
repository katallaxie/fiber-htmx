package csrf

import htmx "github.com/katallaxie/fiber-htmx"

const defaultTokenName = "CSRFToken"

// TokenProps is the struct that holds the CSRF properties.
type TokenProps struct {
	// Token is the CSRF token
	Token string
	// Name is the name of the CSRF token
	Name string
}

// Token is the struct that holds the CSRF properties.
func Token(props TokenProps) htmx.Node {
	if props.Name == "" {
		props.Name = defaultTokenName
	}

	return htmx.Input(
		htmx.Type("hidden"),
		htmx.Name(props.Name),
		htmx.Value(props.Token),
	)
}
