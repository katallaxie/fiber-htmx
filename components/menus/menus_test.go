package menus_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/menus"
	"github.com/stretchr/testify/require"
)

func TestMenu(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    menus.Props
	}{
		{
			name:     "default",
			props:    menus.Props{},
			expected: `<ul class="menu"></ul>`,
		},
		{
			name:     "with class",
			props:    menus.Props{ClassNames: htmx.ClassNames{"bg-red-500": true}},
			expected: `<ul class="bg-red-500 menu"></ul>`,
		},
		{
			name:     "with multiple classes",
			props:    menus.Props{ClassNames: htmx.ClassNames{"bg-red-500": true, "text-white": true}},
			expected: `<ul class="bg-red-500 menu text-white"></ul>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := menus.Menu(
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

func TestMenuVertical(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    menus.Props
	}{
		{
			name:     "default",
			props:    menus.Props{},
			expected: `<ul class="menu menu-vertical"></ul>`,
		},
		{
			name:     "with class",
			props:    menus.Props{ClassNames: htmx.ClassNames{"bg-red-500": true}},
			expected: `<ul class="bg-red-500 menu menu-vertical"></ul>`,
		},
		{
			name:     "with multiple classes",
			props:    menus.Props{ClassNames: htmx.ClassNames{"bg-red-500": true, "text-white": true}},
			expected: `<ul class="bg-red-500 menu menu-vertical text-white"></ul>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := menus.MenuVertical(
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

func TestMenuHorizontal(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    menus.Props
	}{
		{
			name:     "default",
			props:    menus.Props{},
			expected: `<ul class="menu menu-horizontal"></ul>`,
		},
		{
			name:     "with class",
			props:    menus.Props{ClassNames: htmx.ClassNames{"bg-red-500": true}},
			expected: `<ul class="bg-red-500 menu menu-horizontal"></ul>`,
		},
		{
			name:     "with multiple classes",
			props:    menus.Props{ClassNames: htmx.ClassNames{"bg-red-500": true, "text-white": true}},
			expected: `<ul class="bg-red-500 menu menu-horizontal text-white"></ul>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := menus.MenuHorizontal(
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
		props    menus.ItemProps
	}{
		{
			name:     "default",
			props:    menus.ItemProps{},
			expected: `<li class=""></li>`,
		},
		{
			name:     "with class",
			props:    menus.ItemProps{ClassNames: htmx.ClassNames{"bg-red-500": true}},
			expected: `<li class="bg-red-500"></li>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := menus.Item(
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
