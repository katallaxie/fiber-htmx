package jsdeliver

import (
	"testing"

	"github.com/katallaxie/fiber-htmx/imports"
	"github.com/stretchr/testify/require"
	"github.com/zeiss/pkg/jsonx"
)

func TestResolve(t *testing.T) {
	cdn := New()

	pkg := &imports.ExactPackage{
		Name:    "htmx.org",
		Version: "2.0.4",
	}

	err := cdn.Resolve(t.Context(), pkg)
	require.NoError(t, err)
	require.NotNil(t, pkg)

	require.Equal(t, "htmx.org", pkg.Name)
	require.Equal(t, "2.0.4", pkg.Version)
	require.Len(t, pkg.Files, 35)

	b, err := jsonx.Prettify(pkg)
	require.NoError(t, err)
	require.NotEmpty(t, b)
}

func TestNew(t *testing.T) {
	t.Parallel()

	pkg := &imports.ExactPackage{
		Name:    "htmx.org",
		Version: "2.0.4",
	}

	resolver := New()

	b, err := imports.New(resolver).
		Packages(pkg).
		Require(&imports.Require{
			File: "dist/htmx.min.js",
		}).
		Build(t.Context())

	require.NoError(t, err)
	require.NotNil(t, b)
	require.Len(t, b.Imports, 1)
	require.Len(t, b.Integrity, 1)

	require.Contains(t, b.Imports, "htmx.org")
	require.Contains(t, b.Integrity, "dist/htmx.min.js")
}
