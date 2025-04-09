package imports

import htmx "github.com/katallaxie/fiber-htmx"

// Builder ...
type Builder interface {
	Imports() htmx.Imports
}

type builder struct {
	baseURL  string
	provider Provider
	registry Registry
	ignore   []string
	imports  htmx.Imports
}

// Opt ...
type Opt func(*builder)

// WithBaseURL ...
func WithBaseURL(baseURL string) Opt {
	return func(b *builder) {
		b.baseURL = baseURL
	}
}

// WithImports ...
func WithImports(imports htmx.Imports) Opt {
	return func(b *builder) {
		b.imports = imports
	}
}

// WithProvider ...
func WithProvider(provider Provider) Opt {
	return func(b *builder) {
		b.provider = provider
	}
}

// WithRegistry ...
func WithRegistry(registry Registry) Opt {
	return func(b *builder) {
		b.registry = registry
	}
}

// WithIgnore ...
func WithIgnore(ignore []string) Opt {
	return func(b *builder) {
		b.ignore = ignore
	}
}

// New ...
func New(opts ...Opt) Builder {
	b := new(builder)

	b.provider = &UnimplementedProvider{}
	b.imports = htmx.Imports{}
	b.registry = NPM
	b.ignore = []string{}

	for _, opt := range opts {
		opt(b)
	}

	return b
}

// Imports ...
func (b *builder) Imports() htmx.Imports {
	return htmx.Imports{}
}
