package htmx

import (
	"encoding/json"

	"github.com/katallaxie/fiber-htmx/imports"
	"github.com/katallaxie/pkg/conv"
	"github.com/katallaxie/pkg/errorx"
	"github.com/zeiss/pkg/utilx"
)

// ImportMap is a struct that represents an import map.
// sse https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script/type/importmap
type ImportMap struct {
	// Imports is a map of module specifiers to URLs.
	Imports map[string]string `json:"imports"`
	// Scopes is a map of scopes to maps of module specifiers to URLs.
	Scopes map[string]map[string]string `json:"scopes,omitempty"`
	// Integrity is a map of URLs to integrity metadata.
	Integrity map[string]string `json:"integrity,omitempty"`
}

// NewImportMap creates a new import map.
func NewImportMap() ImportMap {
	return ImportMap{
		Imports:   make(map[string]string),
		Scopes:    make(map[string]map[string]string),
		Integrity: make(map[string]string),
	}
}

// ImportsProp ...
type ImportsProp struct {
	Pkgs     []imports.ExactPackage
	Requires []imports.Require
	Resolver imports.Resolver
}

// Imports ...
func Imports(props ImportsProp) Node {
	im := NewImportMap()

	for i := range props.Pkgs {
		if err := props.Resolver.Resolve(&props.Pkgs[i]); err != nil {
			panic(err)
		}
	}

	// TODO(katallaxie): add support for scopes and improve runtime performance
	for _, req := range props.Requires {
		for _, pkg := range props.Pkgs {
			for _, file := range pkg.Files {
				switch f := file.(type) {
				case *imports.FileJS:
					if req.File == f.LocalPath {
						im.Imports[pkg.Name] = f.Path
						if utilx.NotEmpty(req.As) {
							im.Imports[req.As] = f.Path
						}

						if utilx.NotEmpty(f.Integrity) {
							im.Integrity[f.LocalPath] = f.Integrity
						}
					}
				default:
				}
			}
		}
	}

	j := errorx.Must(json.Marshal(im))

	return Fragment(
		Script(
			Src("https://unpkg.com/es-module-shims@1.10.0/dist/es-module-shims.js"),
			CrossOrigin("anonymous"),
			Integrity("sha384-BTO8nLHukFlPxTSib9wgQyLgd2oYLxp24Goxje82QeHp7cwyUtgx4Z32PCEb3Q09"),
		),
		Script(
			Attribute("type", "importmap"),
			Raw(conv.String(j)),
		),
	)
}
