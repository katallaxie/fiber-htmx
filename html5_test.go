package htmx_test

import (
	"bytes"
	"io"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Benchmark_HTML5_Render(b *testing.B) {
	nodes := []htmx.Node{}

	for i := 0; i < 10000; i++ {
		nodes = append(nodes, htmx.Element("div"))
	}

	for i := 0; i < b.N; i++ {
		err := htmx.HTML5(
			htmx.HTML5Props{
				Title:       "Hello, World!",
				Description: "An example of an HTML5 document.",
				Language:    "en",
				Head: []htmx.Node{
					htmx.Meta(htmx.Charset("utf-8")),
					htmx.Meta(htmx.Name("viewport"), htmx.Content("width=device-width, initial-scale=1")),
					htmx.TitleElement(htmx.Text("Hello, World!")),
				},
			}, nodes...).Render(io.Discard)
		assert.NoError(b, err)
	}
}

func Benchmark_ClassNames_Render(b *testing.B) {
	c := htmx.ClassNames{"foo": true, "bar": true, "baz": true, "qux": true, "quux": true, "corge": true, "grault": true, "garply": true, "waldo": true, "fred": true, "plugh": true, "xyzzy": true, "thud": true}

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		buf.Reset()
		err := c.Render(&buf)
		require.NoError(b, err)
	}
}

func Test_ClassNames_Render(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		c    htmx.ClassNames
		want string
	}{
		{
			name: "empty",
			c:    htmx.ClassNames{},
			want: " class=\"\"",
		},
		{
			name: "single",
			c:    htmx.ClassNames{"foo": true},
			want: " class=\"foo\"",
		},
		{
			name: "multiple",
			c:    htmx.ClassNames{"foo": true, "bar": true},
			want: " class=\"bar foo\"",
		},
		{
			name: "overwrite",
			c:    htmx.ClassNames{"foo": false},
			want: " class=\"\"",
		},
		{
			name: "merge",
			c:    htmx.ClassNames{"foo": true},
			want: " class=\"foo\"",
		},
		{
			name: "merge overwrite",
			c:    htmx.ClassNames{"foo": true},
			want: " class=\"foo\"",
		},
		{
			name: "trim space",
			c:    htmx.ClassNames{" foo ": true},
			want: " class=\"foo\"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var b bytes.Buffer
			err := test.c.Render(&b)

			require.NoError(t, err)
			assert.Equal(t, test.want, b.String())
		})
	}
}
