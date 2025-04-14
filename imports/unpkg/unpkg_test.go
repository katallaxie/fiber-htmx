package unpkg

import (
	"testing"

	"github.com/katallaxie/fiber-htmx/imports"
	"github.com/stretchr/testify/require"
)

func TestResolve(t *testing.T) {
	cdn := New()

	pkg := &imports.ExactPackage{
		Name:    "htmx.org",
		Version: "2.0.4",
		Require: []imports.Require{
			{
				Raw: "https://unpkg.com/htmx.org@2.0.4/dist/htmx.min.js",
			},
		},
	}

	err := cdn.Resolve(t.Context(), pkg)
	require.NoError(t, err)
	require.NotNil(t, pkg)

	require.Equal(t, "htmx.org", pkg.Name)
	require.Equal(t, "2.0.4", pkg.Version)
	require.Len(t, pkg.Files, 35)
}
