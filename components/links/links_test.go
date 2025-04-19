package links_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/links"
	"github.com/stretchr/testify/require"
)

func TestLink(t *testing.T) {
	tests := []struct {
		name     string
		props    links.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    links.Props{},
			expected: "<a class=\"link\" href=\"\" rel=\"\"></a>",
		},
		{
			name:     "active",
			props:    links.Props{Active: true},
			expected: "<a class=\"active link\" href=\"\" rel=\"\"></a>",
		},
		{
			name:     "with href",
			props:    links.Props{Href: "/test"},
			expected: "<a class=\"link\" href=\"/test\" rel=\"\"></a>",
		},
		{
			name:     "with rel",
			props:    links.Props{Rel: "noopener"},
			expected: "<a class=\"link\" href=\"\" rel=\"noopener\"></a>",
		},
		{
			name:     "with class",
			props:    links.Props{ClassNames: htmx.ClassNames{"custom-class": true}},
			expected: "<a class=\"custom-class link\" href=\"\" rel=\"\"></a>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := links.Link(
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
