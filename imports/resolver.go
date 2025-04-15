package imports

import (
	"context"

	"github.com/katallaxie/pkg/utilx"
)

// Imports is a struct that represents an import map for JavaScript modules.
// sse https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script/type/importmap
type Imports struct {
	// Imports is a map of module specifiers to URLs.
	Imports map[string]string `json:"imports"`
	// Scopes is a map of scopes to maps of module specifiers to URLs.
	Scopes map[string]map[string]string `json:"scopes,omitempty"`
	// Integrity is a map of URLs to integrity metadata.
	Integrity map[string]string `json:"integrity,omitempty"`
}

// NewImports creates a new Imports instance.
func NewImports() Imports {
	return Imports{
		Imports:   make(map[string]string),
		Scopes:    make(map[string]map[string]string),
		Integrity: make(map[string]string),
	}
}

// Builder is a struct that represents a builder for an import map.
type Builder interface {
	// Packages adds packahges to the import map.
	Packages(pkgs ...*ExactPackage) Builder
	// Require a package to the import map.
	Require(pkgs ...*Require) Builder
	// Build returns the import map.
	Build(ctx context.Context) (Imports, error)
}

// BuilderImpl is a struct that implements the Builder interface.
type BuilderImpl struct {
	pkgs     []*ExactPackage
	require  []*Require
	resolver Resolver
}

// New creates a new Imports instance.
func New(resolver Resolver) Builder {
	return &BuilderImpl{
		resolver: resolver,
	}
}

// Packages adds packages to the import map.
func (b *BuilderImpl) Packages(pkgs ...*ExactPackage) Builder {
	b.pkgs = append(b.pkgs, pkgs...)
	return b
}

// Require adds a package to the import map.
func (b *BuilderImpl) Require(pkgs ...*Require) Builder {
	b.require = append(b.require, pkgs...)
	return b
}

// Build builds the import map.
func (b *BuilderImpl) Build(ctx context.Context) (Imports, error) {
	imports := NewImports()

	for _, pkg := range b.pkgs {
		if err := b.resolver.Resolve(ctx, pkg); err != nil {
			return imports, err
		}
	}

	// TODO(katallaxie): add support for scopes and improve runtime performance
	for _, req := range b.require {
		for _, pkg := range b.pkgs {
			for _, file := range pkg.Files {
				switch f := file.(type) {
				case *FileJS:
					if req.File == f.LocalPath {
						imports.Imports[pkg.Name] = f.LocalPath
						if utilx.NotEmpty(req.As) {
							imports.Imports[req.As] = f.LocalPath
						}

						if utilx.NotEmpty(f.Integrity) {
							imports.Integrity[f.LocalPath] = f.Integrity
						}
					}
				default:
				}
			}
		}
	}

	return imports, nil
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

// Require is a struct that represents a package with its name and version.
type Require struct {
	// File is the file path of the package. This is the local path not the full url in the CDN.
	File string
	// Raw is the raw file path of the package.
	Raw string
	// As is the alias for the package.
	As string
}

// ExactPackage is a struct that represents a package with its name, version, and files.
type ExactPackage struct {
	// Name is the name of the package.
	Name string
	// Version is the version of the package.
	Version string
	// Files is a list of files associated with the package.
	Files []isFile_Type
}

type isFile_Type interface {
	isFile_Type()
}

// FileJS is a struct that represents a JavaScript file.
type FileJS struct {
	Path      string
	LocalPath string
	Integrity string
}

func (*FileJS) isFile_Type() {}

// FileCSS is a struct that represents a CSS file.
type FileCSS struct {
	Path      string
	LocalPath string
	Integrity string
}

func (*FileCSS) isFile_Type() {}

// FileUnkown is a struct that represents an unknown file type.
type FileUnkown struct {
	Path      string
	LocalPath string
	Integrity string
}

func (*FileUnkown) isFile_Type() {}
