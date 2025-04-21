package stacks_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/stacks"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"stack\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"custom-class stack\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"stack\">child</div>",
		},
		{
			name:     "with classes and children",
			classes:  htmx.ClassNames{"custom-class": true},
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"custom-class stack\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := stacks.Stack(
				stacks.StackProps{
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

func TestStackTop(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"stack stack-top\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"custom-class stack stack-top\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"stack stack-top\">child</div>",
		},
		{
			name:     "with classes and children",
			classes:  htmx.ClassNames{"custom-class": true},
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"custom-class stack stack-top\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := stacks.StackTop(
				stacks.StackProps{
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

func TestStackStart(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"stack stack-start\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"custom-class stack stack-start\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"stack stack-start\">child</div>",
		},
		{
			name:     "with classes and children",
			classes:  htmx.ClassNames{"custom-class": true},
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"custom-class stack stack-start\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := stacks.StackStart(
				stacks.StackProps{
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

func TestStackEnd(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"stack stack-end\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"custom-class stack stack-end\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"stack stack-end\">child</div>",
		},
		{
			name:     "with classes and children",
			classes:  htmx.ClassNames{"custom-class": true},
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"custom-class stack stack-end\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := stacks.StackEnd(
				stacks.StackProps{
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
