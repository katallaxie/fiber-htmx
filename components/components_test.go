package components_test

import (
	"bytes"
	"fmt"

	"github.com/katallaxie/fiber-htmx/components/accordions"
	"github.com/katallaxie/fiber-htmx/components/buttons"
)

func ExampleButton() {
	btn := buttons.Button(
		buttons.ButtonProps{},
	)

	var buf bytes.Buffer
	_ = btn.Render(&buf)

	fmt.Println(buf.String())
	// Output:
	// <button class="btn" type="button"></button>
}

func ExampleAccordion() {
	accordion := accordions.Accordion(
		accordions.AccordionProps{},
	)

	var buf bytes.Buffer
	_ = accordion.Render(&buf)
	fmt.Println(buf.String())
	// Output:
	// <div class="bg-base-200 border border-base-300 collapse"><input type="radio" name=""></div>
}
