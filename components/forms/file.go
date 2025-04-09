package forms

import htmx "github.com/katallaxie/fiber-htmx"

// FileInputProps represents the properties for a file input element.
type FileInputProps struct {
	ClassNames htmx.ClassNames
	Disabled   bool
}

// File generates a file input element based on the provided properties.
func FileInput(p FileInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("file"),
		htmx.Merge(
			htmx.ClassNames{
				"file-input": true,
				"w-full":     true,
				"max-w-xs":   true,
			},
			p.ClassNames,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// FileInputBordered is a component that displays a bordered file input.
func FileInputBordered(p FileInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("file"),
		htmx.Merge(
			htmx.ClassNames{
				"file-input":          true,
				"file-input-bordered": true,
				"w-full":              true,
				"max-w-xs":            true,
			},
			p.ClassNames,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// FileInputGhost is a component that displays a ghost file input.
func FileInputGhost(p FileInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("file"),
		htmx.Merge(
			htmx.ClassNames{
				"file-input":       true,
				"file-input-ghost": true,
				"w-full":           true,
				"max-w-xs":         true,
			},
			p.ClassNames,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// FileInputPrimary is a component that displays a primary file input.
func FileInputPrimary(p FileInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("file"),
		htmx.Merge(
			htmx.ClassNames{
				"file-input":          true,
				"file-input-bordered": true,
				"file-input-primary":  true,
				"w-full":              true,
				"max-w-xs":            true,
			},
			p.ClassNames,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// FileInputSecondary is a component that displays a secondary file input.
func FileInputSecondary(p FileInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("file"),
		htmx.Merge(
			htmx.ClassNames{
				"file-input":           true,
				"file-input-bordered":  true,
				"file-input-secondary": true,
				"w-full":               true,
				"max-w-xs":             true,
			},
			p.ClassNames,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// FileInputAccent is a component that displays an accent file input.
func FileInputAccent(p FileInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("file"),
		htmx.Merge(
			htmx.ClassNames{
				"file-input":          true,
				"file-input-bordered": true,
				"file-input-accent":   true,
				"w-full":              true,
				"max-w-xs":            true,
			},
			p.ClassNames,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// FileInputInfo is a component that displays an info file input.
func FileInputInfo(p FileInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("file"),
		htmx.Merge(
			htmx.ClassNames{
				"file-input":          true,
				"file-input-bordered": true,
				"file-input-info":     true,
				"w-full":              true,
				"max-w-xs":            true,
			},
			p.ClassNames,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// FileInputSuccess is a component that displays a success file input.
func FileInputSuccess(p FileInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("file"),
		htmx.Merge(
			htmx.ClassNames{
				"file-input":          true,
				"file-input-bordered": true,
				"file-input-success":  true,
				"w-full":              true,
				"max-w-xs":            true,
			},
			p.ClassNames,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// FileInputWarning is a component that displays a warning file input.
func FileInputWarning(p FileInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("file"),
		htmx.Merge(
			htmx.ClassNames{
				"file-input":          true,
				"file-input-bordered": true,
				"file-input-warning":  true,
				"w-full":              true,
				"max-w-xs":            true,
			},
			p.ClassNames,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// FileInputError is a component that displays an error file input.
func FileInputError(p FileInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("file"),
		htmx.Merge(
			htmx.ClassNames{
				"file-input":          true,
				"file-input-bordered": true,
				"file-input-error":    true,
				"w-full":              true,
				"max-w-xs":            true,
			},
			p.ClassNames,
		),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}
