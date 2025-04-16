package accordions_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/accordions"
	"github.com/stretchr/testify/require"
)

func TestAccordion(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		expected string
		checked  bool
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"bg-base-200 border border-base-300 collapse\"><input type=\"radio\" name=\"\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"bg-base-200 border border-base-300 collapse custom-class\"><input type=\"radio\" name=\"\"></div>",
		},
		{
			name:     "checked",
			classes:  nil,
			expected: "<div class=\"bg-base-200 border border-base-300 collapse\"><input type=\"radio\" name=\"\" checked=\"checked\"></div>",
			checked:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := accordions.Accordion(
				accordions.AccordionProps{
					ClassNames: test.classes,
					Checked:    test.checked,
				},
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestAccordionTitle(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"collapse-title font-semibold\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"collapse-title custom-class font-semibold\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			expected: "<div class=\"collapse-title font-semibold\">child</div>",
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "with classes and children",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"collapse-title custom-class font-semibold\">child</div>",
			children: []htmx.Node{htmx.Text("child")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := accordions.AccordionTitle(
				accordions.AccordionTitleProps{
					ClassNames: test.classes,
				},
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestAccordionContent(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"collapse-content\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"collapse-content custom-class\"></div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := accordions.AccordionContent(
				accordions.AccordionContentProps{
					ClassNames: test.classes,
				},
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}
