package htmx_test

import (
	"testing"

	"github.com/katallaxie/fiber-htmx/internal/assert"

	htmx "github.com/katallaxie/fiber-htmx"
)

func Test_MustacheTemplate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		v    string
		want string
	}{
		{
			name: "mustache-template",
			v:    "value",
			want: ` mustache-template="value"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.MustacheTemplate(test.v)
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_HandlebarsTemplate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		v    string
		want string
	}{
		{
			name: "handlebars-template",
			v:    "value",
			want: ` handlebars-template="value"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.HandlebarsTemplate(test.v)
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_NunjucksTemplate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		v    string
		want string
	}{
		{
			name: "nunjucks-template",
			v:    "value",
			want: ` nunjucks-template="value"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.NunjucksTemplate(test.v)
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_XSLTTemplat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		v    string
		want string
	}{
		{
			name: "xslt-template",
			v:    "value",
			want: ` xslt-template="value"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.XSLTTemplat(test.v)
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_HyperScript(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		v    string
		want string
	}{
		{
			name: "hyperscript",
			v:    "value",
			want: ` _="value"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.HyperScript(test.v)
			assert.Equal(t, test.want, a)
		})
	}
}
