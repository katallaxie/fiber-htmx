package csrf

import (
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/pkg/utilx"
)

const defaultTokenName = "csrftoken"

// Props is the struct that holds the CSRF properties.
type Props struct {
	// Token is the CSRF token
	Token string
	// Name is the name of the CSRF token
	Name string
}

// Token is the struct that holds the CSRF properties.
func Token(props Props) htmx.Node {
	return htmx.Input(
		htmx.Type("hidden"),
		htmx.IfElse(utilx.NotEmpty(props.Name), htmx.Name(props.Name), htmx.Name(defaultTokenName)),
		htmx.Value(props.Token),
	)
}
