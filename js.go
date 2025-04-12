package htmx

import (
	"encoding/json"

	"github.com/katallaxie/pkg/errorx"
	"github.com/zeiss/pkg/conv"
)

// Imports is following the WICG Import Maps proposal.
// sse https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script/type/importmap
type Imports struct {
	// Imports is a map of module specifiers to URLs.
	Imports map[string]string `json:"imports"`
	// Scopes is a map of scopes to maps of module specifiers to URLs.
	Scopes map[string]map[string]string `json:"scopes,omitempty"`
	// Integrity is a map of URLs to integrity metadata.
	Integrity map[string]string `json:"integrity,omitempty"`
}

// ImportMap returns a new ImportMaps object.
func ImportMap(m Imports) Node {
	b := errorx.Ignore(json.Marshal(m))

	return Script(
		Attribute("type", "importmap"),
		Raw(conv.String(b)),
	)
}
