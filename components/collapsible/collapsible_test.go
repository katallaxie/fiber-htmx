package collapsible_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/collapsible"
	"github.com/stretchr/testify/require"
)

func TestCollapsible(t *testing.T) {
	tests := []struct {
		name     string
		props    collapsible.CollapseProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    collapsible.CollapseProps{},
			expected: "<div class=\"bg-base-200 collapse\" tabindex=\"0\"></div>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := collapsible.Collapse(
				collapsible.CollapseProps{},
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}
