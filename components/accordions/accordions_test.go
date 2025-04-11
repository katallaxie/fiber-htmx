package accordions

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/stretchr/testify/require"
)

func TestAccordion(t *testing.T) {
	t.Parallel()

	component := Accordion(
		AccordionProps{
			ClassNames: htmx.ClassNames{
				"bg-base-200": true,
			},
			Name: "test",
		},
	)

	var buf bytes.Buffer
	err := component.Render(&buf)
	require.NoError(t, err)

	require.Equal(t, "<div class=\"bg-base-200 collapse\"></div>", buf.String())
}
