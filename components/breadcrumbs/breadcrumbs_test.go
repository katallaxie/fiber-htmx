package breadcrumbs_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/breadcrumbs"
	"github.com/stretchr/testify/require"
)

func TestBreadcrumbs(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    breadcrumbs.Props
	}{
		{
			name:     "default",
			props:    breadcrumbs.Props{},
			expected: `<div class="breadcrumbs"></div>`,
		},
		{
			name:     "with class",
			props:    breadcrumbs.Props{ClassNames: htmx.ClassNames{"bg-red-500": true}},
			expected: `<div class="bg-red-500 breadcrumbs"></div>`,
		},
		{
			name:     "with multiple classes",
			props:    breadcrumbs.Props{ClassNames: htmx.ClassNames{"bg-red-500": true, "text-white": true}},
			expected: `<div class="bg-red-500 breadcrumbs text-white"></div>`,
		},
		{
			name:     "with children",
			props:    breadcrumbs.Props{},
			children: []htmx.Node{htmx.Text("Avatar Content")},
			expected: `<div class="breadcrumbs">Avatar Content</div>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := breadcrumbs.Breadcrumbs(
				test.props,
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestItems(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    breadcrumbs.Props
	}{
		{
			name:     "default",
			props:    breadcrumbs.Props{},
			expected: `<ul class=""></ul>`,
		},
		{
			name:     "with class",
			props:    breadcrumbs.Props{ClassNames: htmx.ClassNames{"bg-red-500": true}},
			expected: `<ul class="bg-red-500"></ul>`,
		},
		{
			name:     "with multiple classes",
			props:    breadcrumbs.Props{ClassNames: htmx.ClassNames{"bg-red-500": true, "text-white": true}},
			expected: `<ul class="bg-red-500 text-white"></ul>`,
		},
		{
			name:     "with children",
			props:    breadcrumbs.Props{},
			children: []htmx.Node{htmx.Text("Avatar Content")},
			expected: `<ul class="">Avatar Content</ul>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := breadcrumbs.Items(
				test.props,
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestItem(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    breadcrumbs.Props
	}{
		{
			name:     "default",
			props:    breadcrumbs.Props{},
			expected: `<li class=""></li>`,
		},
		{
			name:     "with class",
			props:    breadcrumbs.Props{ClassNames: htmx.ClassNames{"bg-red-500": true}},
			expected: `<li class="bg-red-500"></li>`,
		},
		{
			name:     "with multiple classes",
			props:    breadcrumbs.Props{ClassNames: htmx.ClassNames{"bg-red-500": true, "text-white": true}},
			expected: `<li class="bg-red-500 text-white"></li>`,
		},
		{
			name:     "with children",
			props:    breadcrumbs.Props{},
			children: []htmx.Node{htmx.Text("Avatar Content")},
			expected: `<li class="">Avatar Content</li>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := breadcrumbs.Item(
				test.props,
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}
