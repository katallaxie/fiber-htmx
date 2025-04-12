package htmx_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImportMap(t *testing.T) {
	t.Parallel()

	m := htmx.ImportMap(htmx.Imports{
		Imports: map[string]string{
			"htmx": "https://unpkg.com/htmx.org@1.9.3",
		},
	})

	var buf bytes.Buffer
	err := m.Render(&buf)
	require.NoError(t, err)

	assert.Equal(t, `<script type="importmap">{"imports":{"htmx":"https://unpkg.com/htmx.org@1.9.3"}}</script>`, buf.String())
}
