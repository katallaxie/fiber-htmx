package status_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/status"
	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	tests := []struct {
		name     string
		props    status.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    status.Props{},
			expected: "<div class=\"status\" role=\"status\" aria-label=\"\"></div>",
		},
		{
			name:     "with class",
			props:    status.Props{ClassNames: htmx.ClassNames{"custom-class": true}},
			expected: "<div class=\"custom-class status\" role=\"status\" aria-label=\"\"></div>",
		},
		{
			name:     "with aria label",
			props:    status.Props{AriaLabel: "test"},
			expected: "<div class=\"status\" role=\"status\" aria-label=\"test\"></div>",
		},
		{
			name:     "with animate bounce",
			props:    status.Props{AnimateBounce: true},
			expected: "<div class=\"animate-bounce status\" role=\"status\" aria-label=\"\"></div>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := status.Status(
				tt.props,
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}
