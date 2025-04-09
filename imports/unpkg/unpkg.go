package unpkg

import (
	"fmt"

	"github.com/katallaxie/fiber-htmx/imports"
)

const (
	// URL is the URL of the unpkg provider.
	URL = "https://unpkg.com"
)

type unpkgProvider struct {
	imports.UnimplementedProvider
}

// New returns a new unpkg provider.
func New() *unpkgProvider {
	return &unpkgProvider{}
}

// ToURL returns the URL of the provider.
func (p *unpkgProvider) ToURL(pkg imports.ExactPackage) string {
	return fmt.Sprintf("%s/%s@%s", URL, pkg.Name(), pkg.Version())
}
