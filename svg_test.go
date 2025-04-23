package htmx_test

import (
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/internal/assert"
)

func Test_SVG(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "svg",
			want: `<svg xmlns="http://www.w3.org/2000/svg"></svg>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.SVG()
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_ClipRule(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "clip-rule",
			want: ` clip-rule="evenodd"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.ClipRule("evenodd")
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_Fill(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "fill",
			want: ` fill="black"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.Fill("black")
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_FillRule(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "fill-rule",
			want: ` fill-rule="evenodd"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.FillRule("evenodd")
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_Stroke(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "stroke",
			want: ` stroke="black"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.Stroke("black")
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_StrokeWidth(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "stroke-width",
			want: ` stroke-width="2"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.StrokeWidth("2")
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_ViewBox(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "view-box",
			want: ` viewBox="0 0 100 100"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.ViewBox("0 0 100 100")
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_G(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "g",
			want: `<g></g>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.G()
			assert.Equal(t, test.want, a)
		})
	}
}
