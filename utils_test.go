package htmx_test

import (
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name string
		in   []htmx.ClassNames
		out  htmx.ClassNames
	}{
		{
			name: "merge",
			in: []htmx.ClassNames{
				{"a": true, "b": false},
				{"b": true, "c": false},
			},
			out: htmx.ClassNames{"a": true, "b": true, "c": false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, htmx.Merge(tt.in...))
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	classNames := []htmx.ClassNames{
		{"a": true, "b": false},
		{"b": true, "c": false},
		{"c": true, "d": false},
		{"d": true, "e": false},
		{"e": true, "f": false},
		{"f": true, "g": false},
		{"g": true, "h": false},
		{"h": true, "i": false},
		{"i": true, "j": false},
		{"j": true, "k": false},
		{"k": true, "l": false},
		{"l": true, "m": false},
		{"m": true, "n": false},
		{"n": true, "o": false},
		{"o": true, "p": false},
		{"p": true, "q": false},
		{"q": true, "r": false},
		{"r": true, "s": false},
		{"s": true, "t": false},
		{"t": true, "u": false},
		{"u": true, "v": false},
		{"v": true, "w": false},
		{"w": true, "x": false},
		{"x": true, "y": false},
		{"y": true, "z": false},
		{"z": true, "a": false},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		htmx.Merge(classNames...)
	}
}

func TestJsonSerializeOrEmpty(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		out  string
	}{
		{
			name: "serialize",
			in:   map[string]interface{}{"a": 1, "b": "2"},
			out:  `{"a":1,"b":"2"}`,
		},
		{
			name: "empty",
			in:   nil,
			out:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, htmx.JsonSerializeOrEmpty(tt.in))
		})
	}
}
