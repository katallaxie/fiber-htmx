package csrf_test

import (
	"bytes"
	"testing"

	"github.com/katallaxie/fiber-htmx/components/csrf"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	tests := []struct {
		name     string
		props    csrf.Props
		expected string
	}{
		{
			name:     "default",
			props:    csrf.Props{},
			expected: `<input type="hidden" name="csrftoken" value="">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := csrf.Token(tt.props)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}
