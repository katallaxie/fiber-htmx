package lists_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/lists"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	tests := []struct {
		name     string
		props    lists.ListProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    lists.ListProps{},
			expected: "<ul class=\"list\"></ul>",
		},
		{
			name:     "with class",
			props:    lists.ListProps{ClassNames: htmx.ClassNames{"custom-class": true}},
			expected: "<ul class=\"custom-class list\"></ul>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := lists.List(
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

func TestListRow(t *testing.T) {
	tests := []struct {
		name     string
		props    lists.ListRowProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    lists.ListRowProps{},
			expected: "<li class=\"list-row\"></li>",
		},
		{
			name:     "with class",
			props:    lists.ListRowProps{ClassNames: htmx.ClassNames{"custom-class": true}},
			expected: "<li class=\"custom-class list-row\"></li>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := lists.ListRow(
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

func TestListHeader(t *testing.T) {
	tests := []struct {
		name     string
		props    lists.ListHeaderProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    lists.ListHeaderProps{},
			expected: "<li class=\"opacity-60 p-4 pb-2 text-xs tracking-wide\"></li>",
		},
		{
			name:     "with class",
			props:    lists.ListHeaderProps{ClassNames: htmx.ClassNames{"custom-class": true}},
			expected: "<li class=\"custom-class opacity-60 p-4 pb-2 text-xs tracking-wide\"></li>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := lists.ListHeader(
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
