package unpkg

import (
	"testing"

	"github.com/katallaxie/fiber-htmx/imports"
	"github.com/stretchr/testify/require"
)

func TestResolve(t *testing.T) {
	cdn := New()

	pkg := &imports.Package{
		Name:    "htmx.org",
		Version: "2.0.4",
	}

	err := cdn.Resolve(t.Context(), pkg)
	require.NoError(t, err)
}
