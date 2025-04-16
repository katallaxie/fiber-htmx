package joins_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/joins"
	"github.com/stretchr/testify/require"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"join\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"custom-class join\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"join\">child</div>",
		},
		{
			name:     "with classes and children",
			classes:  htmx.ClassNames{"custom-class": true},
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"custom-class join\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := joins.Join(
				joins.JoinProps{
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

func TestJoinVertical(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"join join-vertical\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"custom-class join join-vertical\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"join join-vertical\">child</div>",
		},
		{
			name:     "with classes and children",
			classes:  htmx.ClassNames{"custom-class": true},
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"custom-class join join-vertical\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := joins.JoinVertical(
				joins.JoinProps{
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

func TestJoinHorizontal(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"join join-horizontal\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"custom-class join join-horizontal\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"join join-horizontal\">child</div>",
		},
		{
			name:     "with classes and children",
			classes:  htmx.ClassNames{"custom-class": true},
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"custom-class join join-horizontal\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := joins.JoinHorizontal(
				joins.JoinProps{
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
