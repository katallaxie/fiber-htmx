package imports

import (
	htmx "github.com/katallaxie/fiber-htmx"
)

// ImportsProp ...
type ImportsProp struct {
	Imports htmx.Imports
}

// Imports ...
func Imports(props ImportsProp) htmx.Node {
	return htmx.Fragment(
		htmx.Script(
			htmx.Src("https://unpkg.com/es-module-shims@1.10.0/dist/es-module-shims.js"),
			htmx.CrossOrigin("anonymous"),
			htmx.Integrity("sha384-BTO8nLHukFlPxTSib9wgQyLgd2oYLxp24Goxje82QeHp7cwyUtgx4Z32PCEb3Q09"),
		),
		htmx.ImportMap(props.Imports),
	)
}
