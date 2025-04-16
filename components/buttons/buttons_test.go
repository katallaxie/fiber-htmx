package buttons_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/buttons"
	"github.com/stretchr/testify/require"
)

func TestButton(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    buttons.ButtonProps
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<button class=\"btn\" type=\"button\"></button>",
			props:    buttons.ButtonProps{},
		},
		{
			name:     "with classes",
			expected: "<button class=\"btn custom-class\" type=\"button\"></button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
		},
		{
			name:     "with children",
			expected: "<button class=\"btn\" type=\"button\">child</button>",
			props:    buttons.ButtonProps{},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "with classes and children",
			expected: "<button class=\"btn custom-class\" type=\"button\">child</button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "disabled",
			expected: "<button class=\"btn\" type=\"button\" disabled=\"disabled\"></button>",
			props:    buttons.ButtonProps{Disabled: true},
		},
		{
			name:     "disabled with classes",
			expected: "<button class=\"btn custom-class\" type=\"button\" disabled=\"disabled\"></button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
				Disabled:   true,
			},
		},
		{
			name:     "disabled with children",
			expected: "<button class=\"btn\" type=\"button\" disabled=\"disabled\">child</button>",
			props:    buttons.ButtonProps{Disabled: true},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "disabled with classes and children",
			expected: "<button class=\"btn custom-class\" type=\"button\" disabled=\"disabled\">child</button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
				Disabled:   true,
			},
			children: []htmx.Node{htmx.Text("child")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := buttons.Button(
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

func TestPrimary(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    buttons.ButtonProps
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<button class=\"btn btn-primary\" type=\"button\"></button>",
			props:    buttons.ButtonProps{Type: "button"},
		},
		{
			name:     "with classes",
			expected: "<button class=\"btn btn-primary custom-class\" type=\"button\"></button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
		},
		{
			name:     "with children",
			expected: "<button class=\"btn btn-primary\" type=\"button\">child</button>",
			props:    buttons.ButtonProps{Type: "button"},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "with classes and children",
			expected: "<button class=\"btn btn-primary custom-class\" type=\"button\">child</button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "disabled",
			expected: "<button class=\"btn btn-primary\" type=\"button\" disabled=\"disabled\"></button>",
			props:    buttons.ButtonProps{Type: "button", Disabled: true},
		},
		{
			name:     "disabled with classes",
			expected: "<button class=\"btn btn-primary custom-class\" type=\"button\" disabled=\"disabled\"></button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
				Disabled:   true,
			},
		},
		{
			name:     "disabled with children",
			expected: "<button class=\"btn btn-primary\" type=\"button\" disabled=\"disabled\">child</button>",
			props:    buttons.ButtonProps{Type: "button", Disabled: true},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "disabled with classes and children",
			expected: "<button class=\"btn btn-primary custom-class\" type=\"button\" disabled=\"disabled\">child</button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
				Disabled:   true,
			},
			children: []htmx.Node{htmx.Text("child")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := buttons.Primary(
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

func TestSecondary(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    buttons.ButtonProps
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<button class=\"btn btn-secondary\" type=\"button\"></button>",
			props:    buttons.ButtonProps{Type: "button"},
		},
		{
			name:     "with classes",
			expected: "<button class=\"btn btn-secondary custom-class\" type=\"button\"></button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
		},
		{
			name:     "with children",
			expected: "<button class=\"btn btn-secondary\" type=\"button\">child</button>",
			props:    buttons.ButtonProps{Type: "button"},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "with classes and children",
			expected: "<button class=\"btn btn-secondary custom-class\" type=\"button\">child</button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "disabled",
			expected: "<button class=\"btn btn-secondary\" type=\"button\" disabled=\"disabled\"></button>",
			props:    buttons.ButtonProps{Type: "button", Disabled: true},
		},
		{
			name:     "disabled with classes",
			expected: "<button class=\"btn btn-secondary custom-class\" type=\"button\" disabled=\"disabled\"></button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
				Disabled:   true,
			},
		},
		{
			name:     "disabled with children",
			expected: "<button class=\"btn btn-secondary\" type=\"button\" disabled=\"disabled\">child</button>",
			props:    buttons.ButtonProps{Type: "button", Disabled: true},
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "disabled with classes and children",
			expected: "<button class=\"btn btn-secondary custom-class\" type=\"button\" disabled=\"disabled\">child</button>",
			props: buttons.ButtonProps{
				Type:       "button",
				ClassNames: htmx.ClassNames{"custom-class": true},
				Disabled:   true,
			},
			children: []htmx.Node{htmx.Text("child")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := buttons.Secondary(
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
