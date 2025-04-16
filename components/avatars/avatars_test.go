package avatars_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/avatars"
	"github.com/stretchr/testify/require"
)

func TestAvatr(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    avatars.AvatarProps
	}{
		{
			name:     "default",
			props:    avatars.AvatarProps{},
			expected: `<div class="avatar"></div>`,
		},
		{
			name:     "with class",
			props:    avatars.AvatarProps{ClassNames: htmx.ClassNames{"bg-red-500": true}},
			expected: `<div class="avatar bg-red-500"></div>`,
		},
		{
			name:     "with multiple classes",
			props:    avatars.AvatarProps{ClassNames: htmx.ClassNames{"bg-red-500": true, "text-white": true}},
			expected: `<div class="avatar bg-red-500 text-white"></div>`,
		},
		{
			name:     "with children",
			props:    avatars.AvatarProps{},
			children: []htmx.Node{htmx.Text("Avatar Content")},
			expected: `<div class="avatar">Avatar Content</div>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := avatars.Avatar(
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

func TestAvatarGroup(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    avatars.AvatarProps
	}{
		{
			name:     "default",
			props:    avatars.AvatarProps{},
			expected: `<div class="-space-x-6 avatar-group"></div>`,
		},
		{
			name:     "with class",
			props:    avatars.AvatarProps{ClassNames: htmx.ClassNames{"bg-red-500": true}},
			expected: `<div class="-space-x-6 avatar-group bg-red-500"></div>`,
		},
		{
			name:     "with multiple classes",
			props:    avatars.AvatarProps{ClassNames: htmx.ClassNames{"bg-red-500": true, "text-white": true}},
			expected: `<div class="-space-x-6 avatar-group bg-red-500 text-white"></div>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := avatars.AvatarGroup(
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
