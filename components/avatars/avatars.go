package avatars

import htmx "github.com/katallaxie/fiber-htmx"

// AvatarProps represents the properties for an avatar.
type AvatarProps struct {
	htmx.ClassNames
}

// Avatar generates an avatar based on the provided properties.
func Avatar(p AvatarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// AvatarGroup generates a group of avatars based on the provided properties.
func AvatarGroup(p AvatarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar-group": true,
				"-space-x-6":   true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// AvatarOnline generates an online avatar based on the provided properties.
func AvatarOnline(p AvatarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar":        true,
				"avatar-online": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// AvatarOffline generates an offline avatar based on the provided properties.
func AvatarOffline(p AvatarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar":         true,
				"avatar-offline": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// AvtarPlaceholder generates a placeholder for an avatar based on the provided properties.
func AvatarPlaceholder(p AvatarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"avatar":             true,
			"avatar-placeholder": true,
		},
		htmx.Div(
			htmx.Merge(
				htmx.ClassNames{
					"bg-neutral":           true,
					"w-12":                 true,
					"text-neutral-content": true,
				},
				p.ClassNames,
			),
			htmx.Group(children...),
		),
	)
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
