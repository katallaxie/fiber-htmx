package heros_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/buttons"
	"github.com/katallaxie/fiber-htmx/components/heros"
	"github.com/stretchr/testify/require"
)

func TestHero(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    heros.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<div class=\"hero\"></div>",
			props:    heros.Props{},
		},
		{
			name:     "with classes",
			expected: "<div class=\"custom-class hero\"></div>",
			props: heros.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
		},
		{
			name:     "with children",
			expected: "<div class=\"hero\">child</div>",
			props:    heros.Props{},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "with classes and children",
			expected: "<div class=\"custom-class hero\">child</div>",
			props: heros.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "with buttons",
			expected: "<div class=\"hero\"><button class=\"btn\" type=\"button\"></button></div>",
			props:    heros.Props{},
			children: []htmx.Node{buttons.Button(
				buttons.ButtonProps{},
			)},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := heros.Hero(
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

func TestHeroContent(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    heros.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<div class=\"hero-content\"></div>",
			props:    heros.Props{},
		},
		{
			name:     "with classes",
			expected: "<div class=\"custom-class hero-content\"></div>",
			props: heros.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
		},
		{
			name:     "with children",
			expected: "<div class=\"hero-content\">child</div>",
			props:    heros.Props{},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "with classes and children",
			expected: "<div class=\"custom-class hero-content\">child</div>",
			props: heros.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: []htmx.Node{htmx.Text("child")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := heros.Content(
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

func TestHeroOverlay(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    heros.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<div class=\"hero-overlay\"></div>",
			props:    heros.Props{},
		},
		{
			name:     "with classes",
			expected: "<div class=\"custom-class hero-overlay\"></div>",
			props: heros.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
		},
		{
			name:     "with children",
			expected: "<div class=\"hero-overlay\">child</div>",
			props:    heros.Props{},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "with classes and children",
			expected: "<div class=\"custom-class hero-overlay\">child</div>",
			props: heros.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: []htmx.Node{htmx.Text("child")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := heros.Overlay(
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
