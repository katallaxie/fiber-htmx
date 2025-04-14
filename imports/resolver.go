package imports

import (
	"context"

	"github.com/katallaxie/pkg/utilx"
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

// Resolver is the interface that must be implemented by all resolvers.
type Resolver interface {
	// Resolve resolves a package by its name and version.
	Resolve(ctx context.Context, pkg *ExactPackage) error
}

var _ Resolver = (*UnimplementedResolver)(nil)

// UnimplementedResolver is a resolver that is not implemented.
type UnimplementedResolver struct{}

// Resolve resolves a package by its name and version.
func (r *UnimplementedResolver) Resolve(ctx context.Context, pkg *ExactPackage) error {
	return nil
}

// Resolve resolves a list of packages using the provided resolver.
func Resolve(ctx context.Context, resolver Resolver, pkgs ...*ExactPackage) (Imports, error) {
	im := Imports{
		Imports: make(map[string]string),
	}

	for _, pkg := range pkgs {
		if err := resolver.Resolve(ctx, pkg); err != nil {
			return im, err
		}

		for _, req := range pkg.Require {
			name := pkg.Name
			file := req.File

			if utilx.NotEmpty(req.As) {
				name = req.As
			}

			im.Imports[name] = file
		}
	}

	return im, nil
}

// Require is a struct that represents a package with its name and version.
type Require struct {
	File string
	Raw  string
	As   string
}

// ExactPackage is a struct that represents a package with its name, version, and files.
type ExactPackage struct {
	Name    string
	Version string
	Files   []isFile_Type
	Require []Require
}

type isFile_Type interface {
	isFile_Type()
}

// FileJS is a struct that represents a JavaScript file.
type FileJS struct {
	Path      string
	LocalPath string
}

func (*FileJS) isFile_Type() {}

// FileCSS is a struct that represents a CSS file.
type FileCSS struct {
	Path      string
	LocalPath string
}

func (*FileCSS) isFile_Type() {}

// FileUnkown is a struct that represents an unknown file type.
type FileUnkown struct {
	Path      string
	LocalPath string
}

func (*FileUnkown) isFile_Type() {}
