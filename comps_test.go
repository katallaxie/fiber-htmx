package htmx_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/pkg/conv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := htmx.Element("div").Render(io.Discard)
		assert.NoError(b, err)
	}
}

func Benchmark_AttrRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := htmx.Attribute("href", "/").Render(io.Discard)
		assert.NoError(b, err)
	}
}

func ExampleElement_tag() {
	_ = htmx.Element("div").Render(os.Stdout)
	// Output: <div></div>
}

func ExampleAttribute_bool() {
	_ = htmx.Attribute("disabled", conv.String(true)).Render(os.Stdout)
	// Output: disabled="true"
}

func ExampleAttribute_name_value() {
	_ = htmx.Attribute("href", "/").Render(os.Stdout)
	// Output: href="/"
}

func Test_AttrRender(t *testing.T) {
	tests := []struct {
		desc  string
		name  string
		value string
		out   string
	}{
		{
			desc:  "name only",
			name:  "data-tip",
			value: "",
			out:   ` data-tip=""`,
		},
		{
			desc:  "name and value",
			name:  "href",
			value: "/",
			out:   ` href="/"`,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var bb bytes.Buffer

			err := htmx.Attribute(test.name, test.value).Render(&bb)
			require.NoError(t, err)
			assert.Equal(t, test.out, bb.String())
		})
	}
}

func TestNodeFunc(t *testing.T) {
	t.Run("implements fmt.String", func(t *testing.T) {
		fn := htmx.NodeFunc(func(w io.Writer) error {
			_, _ = w.Write([]byte("hello"))
			return nil
		})

		assert.Implements(t, (*fmt.Stringer)(nil), fn)
		assert.Equal(t, "hello", fmt.Sprint(fn))
	})
}

func TestIf(t *testing.T) {
	tests := []struct {
		desc string
		cond bool
		in   htmx.Node
		out  string
	}{
		{
			desc: "true",
			cond: true,
			in:   htmx.Element("div"),
			out:  "<div><div></div></div>",
		},
		{
			desc: "false",
			cond: false,
			in:   htmx.Element("div"),
			out:  "<div></div>",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var bb bytes.Buffer

			err := htmx.Div(htmx.If(test.cond, test.in)).Render(&bb)
			require.NoError(t, err)
			assert.Equal(t, test.out, bb.String())
		})
	}
}

func TestKeyExists(t *testing.T) {
	tests := []struct {
		desc string
		kv   map[string]string
		fn   func(k, v string) htmx.Node
		in   string
		out  string
	}{
		{
			desc: "exists",
			kv:   map[string]string{"foo": "bar"},
			fn: func(k, v string) htmx.Node {
				return htmx.Attribute(k, v)
			},
			in:  "foo",
			out: "<div foo=\"bar\"></div>",
		},
		{
			desc: "missing",
			kv:   map[string]string{},
			fn: func(k, v string) htmx.Node {
				return htmx.Attribute(k, v)
			},
			in:  "foo",
			out: "<div></div>",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var bb bytes.Buffer

			err := htmx.Div(htmx.KeyExists(test.kv, test.in, test.fn)).Render(&bb)
			require.NoError(t, err)
			assert.Equal(t, test.out, bb.String())
		})
	}
}

func TestFragment(t *testing.T) {
	tests := []struct {
		desc string
		in   []htmx.Node
		out  string
	}{
		{
			desc: "empty",
			in:   nil,
			out:  "",
		},
		{
			desc: "single",
			in:   []htmx.Node{htmx.Element("div")},
			out:  "<div></div>",
		},
		{
			desc: "multiple",
			in:   []htmx.Node{htmx.Element("div"), htmx.Element("span")},
			out:  "<div></div><span></span>",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var bb bytes.Buffer

			err := htmx.Fragment(test.in...).Render(&bb)
			require.NoError(t, err)
			assert.Equal(t, test.out, bb.String())
		})
	}
}

func TestRaw(t *testing.T) {
	tests := []struct {
		desc string
		in   string
		out  string
	}{
		{
			desc: "empty",
			in:   "",
			out:  "",
		},
		{
			desc: "text",
			in:   "hello",
			out:  "hello",
		},
		{
			desc: "html",
			in:   "<div></div>",
			out:  "<div></div>",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var bb bytes.Buffer

			err := htmx.Raw(test.in).Render(&bb)
			require.NoError(t, err)
			assert.Equal(t, test.out, bb.String())
		})
	}
}
