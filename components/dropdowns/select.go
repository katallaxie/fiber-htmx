package dropdowns

import (
	"github.com/katallaxie/fiber-htmx/components/forms"

	htmx "github.com/katallaxie/fiber-htmx"
)

// SingleSelectProps ...
type SingleSelectProps struct {
	URL string
	ID  string
}

// SingleSelect ...
func SingleSelect(props SingleSelectProps, children ...htmx.Node) htmx.Node {
	return Dropdown(
		DropdownProps{
			ClassNames: htmx.ClassNames{},
		},
		DropdownButton(
			DropdownButtonProps{
				ClassNames: htmx.ClassNames{
					"btn-outline": true,
				},
			},
			htmx.ID(props.ID),
			htmx.Text("Dropdown"),
		),
		DropdownMenuItems(
			DropdownMenuItemsProps{},
			forms.TextInputBordered(
				forms.TextInputProps{
					ClassNames: htmx.ClassNames{
						"input-sm": true,
					},
					Placeholder: "Search...",
				},
				htmx.HxPost(props.URL),
				htmx.HxTrigger("input changed delay:500ms, search"),
				htmx.HxTarget("#search-results"),
			),
			htmx.Div(
				htmx.ID("search-results"),
			),
		),
		htmx.Group(children...),
	)
}

// SingleSelectOptionProps ...
type SingleSelectOptionProps struct {
	ID    string
	Name  string
	Value string
}

// SingleSelectOption ...
func SingleSelectOption(props SingleSelectOptionProps, children ...htmx.Node) htmx.Node {
	return DropdownMenuItem(
		DropdownMenuItemProps{},
		htmx.Input(
			htmx.Attribute("type", "hidden"),
			htmx.Attribute("name", props.Name),
			htmx.Value(props.Value),
		),
		htmx.A(
			htmx.Group(children...),
			htmx.HyperScript(`on click put parentElement.innerHTML of me into `+props.ID+`.innerHTML`),
		),
	)
}
