package htmx_test

import (
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/internal/assert"
)

func Test_Async(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{
			name: "async",
			want: " async",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.Async()
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_AutoFocus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{
			name: "autofocus",
			want: " autofocus",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := htmx.AutoFocus()
			assert.Equal(t, test.want, a)
		})
	}
}

func Test_HxPost(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
		url  string
	}{
		{
			name: "hx-post",
			want: " hx-post=\"\"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := htmx.HxPost(test.url)
			assert.Equal(t, test.want, h)
		})
	}
}

func Test_HxPut(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
		url  string
	}{
		{
			name: "hx-put",
			want: " hx-put=\"\"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := htmx.HxPut(test.url)
			assert.Equal(t, test.want, h)
		})
	}
}

func Test_HxPatch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
		url  string
	}{
		{
			name: "hx-patch",
			want: " hx-patch=\"\"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := htmx.HxPatch(test.url)
			assert.Equal(t, test.want, h)
		})
	}
}

func Test_HxDelete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
		url  string
	}{
		{
			name: "hx-delete",
			want: " hx-delete=\"\"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := htmx.HxDelete(test.url)
			assert.Equal(t, test.want, h)
		})
	}
}

func Test_HxGet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
		url  string
	}{
		{
			name: "hx-get",
			want: " hx-get=\"\"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := htmx.HxGet(test.url)
			assert.Equal(t, test.want, h)
		})
	}
}

func Test_HxTrigger(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		want  string
		event []string
	}{
		{
			name:  "hx-trigger",
			want:  " hx-trigger=\"\"",
			event: []string{},
		},
		{
			name:  "hx-trigger",
			want:  " hx-trigger=\"click\"",
			event: []string{"click"},
		},
		{
			name:  "hx-trigger",
			want:  " hx-trigger=\"click onload\"",
			event: []string{htmx.TriggerClick, htmx.TriggerLoad},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := htmx.HxTrigger(test.event...)
			assert.Equal(t, test.want, h)
		})
	}
}

func Test_HxSwap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		want   string
		target string
	}{
		{
			name:   "hx-swap",
			want:   " hx-swap=\"\"",
			target: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := htmx.HxSwap(test.target)
			assert.Equal(t, test.want, h)
		})
	}
}

func Test_HxIndicator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		want   string
		target string
	}{
		{
			name: "hx-indicator",
			want: " hx-indicator=\"\"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := htmx.HxIndicator(test.target)
			assert.Equal(t, test.want, h)
		})
	}
}
