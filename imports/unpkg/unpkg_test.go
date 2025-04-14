package unpkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResolve(t *testing.T) {
	cdn := New()

	pkg, err := cdn.Resolve(t.Context(), "htmx.org", "2.0.4")
	require.NoError(t, err)
	require.NotNil(t, pkg)

	require.Equal(t, "htmx.org", pkg.Name)
	require.Equal(t, "2.0.4", pkg.Version)
	require.Len(t, pkg.Files, 35)
}
