package tabs_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/tabs"
	"github.com/stretchr/testify/require"
)

func TestTabs(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"tabs\" role=\"tablist\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"custom-class tabs\" role=\"tablist\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"tabs\" role=\"tablist\">child</div>",
		},
		{
			name:     "with classes and children",
			classes:  htmx.ClassNames{"custom-class": true},
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"custom-class tabs\" role=\"tablist\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := tabs.Tabs(
				tabs.Props{
					ClassNames: test.classes,
				},
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestBoxed(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"tabs tabs-boxed\" role=\"tablist\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"custom-class tabs tabs-boxed\" role=\"tablist\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"tabs tabs-boxed\" role=\"tablist\">child</div>",
		},
		{
			name:     "with classes and children",
			classes:  htmx.ClassNames{"custom-class": true},
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"custom-class tabs tabs-boxed\" role=\"tablist\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := tabs.Boxed(
				tabs.Props{
					ClassNames: test.classes,
				},
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}
