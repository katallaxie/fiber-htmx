package avatars

import htmx "github.com/katallaxie/fiber-htmx"

// AvatarProps represents the properties for an avatar.
type AvatarProps struct {
	ClassNames htmx.ClassNames
}

// AvatarRounded generates an avatar based on the provided properties.
func AvatarRoundedMedium(p AvatarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.ClassNames{
				"w-24":    true,
				"rounded": true,
			},
			htmx.Group(children...),
		),
	)
}

// AvatarRoundedSmall generates an extra small avatar based on the provided properties.
func AvatarRoundedSmall(p AvatarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.ClassNames{
				"w-8":     true,
				"rounded": true,
			},
			htmx.Group(children...),
		),
	)
}

// AvatarRoundedLarge generates a large avatar based on the provided properties.
func AvatarRoundedLarge(p AvatarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.ClassNames{
				"w-32":    true,
				"rounded": true,
			},
			htmx.Group(children...),
		),
	)
}

// AvatarRoundMedium generates a round avatar based on the provided properties.
func AvatarRoundMedium(p AvatarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.ClassNames{
				"w-24":         true,
				"rounded-full": true,
			},
			htmx.Group(children...),
		),
	)
}

// AvatarRoundSmall generates an extra small round avatar based on the provided properties.
func AvatarRoundSmall(p AvatarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.ClassNames{
				"w-8":          true,
				"rounded-full": true,
			},
			htmx.Group(children...),
		),
	)
}

// AvatarRoundLarge generates a large round avatar based on the provided properties.
func AvatarRoundLarge(p AvatarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.ClassNames{
				"w-32":         true,
				"rounded-full": true,
			},
			htmx.Group(children...),
		),
	)
}
