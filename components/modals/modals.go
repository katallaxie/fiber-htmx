package modals

import htmx "github.com/katallaxie/fiber-htmx"

// ModalProps contains the properties for the modal component.
type ModalProps struct {
	ClassNames htmx.ClassNames
	ID         string
}

// Modal is a component for the htmx modal extension.
func Modal(p ModalProps, children ...htmx.Node) htmx.Node {
	return htmx.Dialog(
		htmx.Merge(
			htmx.ClassNames{
				"modal": true,
			},
		),
		htmx.ID(p.ID),
		htmx.Div(
			htmx.Merge(
				htmx.ClassNames{
					"modal-box": true,
				},
			),
			htmx.Group(children...),
		),
	)
}

// ModalActionProps contains the properties for the modal actions component.
type ModalActionProps struct {
	ClassNames htmx.ClassNames
}

// ModalAction is a component for the htmx modal extension.
func ModalAction(p ModalActionProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"modal-action": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ModalCloseButtonProps contains the properties for the modal close button component.
type ModalCloseButtonProps struct {
	ClassNames htmx.ClassNames
}

// ModalCloseButton is a component for the htmx modal extension.
func ModalCloseButton(p ModalCloseButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Form(
		htmx.Method("dialog"),
		htmx.Button(
			htmx.Merge(
				htmx.ClassNames{
					"btn": true,
				},
				p.ClassNames,
			),
			htmx.Text("Close"),
		),
	)
}
