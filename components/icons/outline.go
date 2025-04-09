// Package icons provides a set of icons that can be used in the application.
// The icons are implemented as functions that return htmx.Node.
//
// https://heroicons.com
package icons

import (
	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/katallaxie/fiber-htmx/components/tailwind"
)

// IconProps ...
type IconProps struct {
	ClassNames htmx.ClassNames
}

// ChevronUpDownOutline ...
func ChevronUpDownOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M8.25 15 12 18.75 15.75 15m-7.5-6L12 5.25 15.75 9"),
		),
	)
}

// ChevronUpOutline ...
func ChevronUpOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M19.5 15 12 7.5 4.5 15"),
		),
	)
}

// ChevronDownOutline ...
func ChevronDownOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M19.5 8.25 12 15.75 4.5 8.25"),
		),
	)
}

// ChevronLeftOutline ...
func ChevronLeftOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M15 19.5 7.5 12 15 4.5"),
		),
	)
}

// ChevronRightOutline ...
func ChevronRightOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M9 19.5 16.5 12 9 4.5"),
		),
	)
}

// SearchOutline ...
func SearchOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"),
		),
	)
}

// BoltOutline ...
func BoltOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m3.75 13.5 10.5-11.25L12 10.5h8.25L9.75 21.75 12 13.5H3.75Z"),
		),
	)
}

// BoltSlashOutline ...
func BoltSlashOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M11.412 15.655 9.75 21.75l3.745-4.012M9.257 13.5H3.75l2.659-2.849m2.048-2.194L14.25 2.25 12 10.5h8.25l-4.707 5.043M8.457 8.457 3 3m5.457 5.457 7.086 7.086m0 0L21 21"),
		),
	)
}

// HeartOutline ...
func HeartOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z"),
		),
	)
}

// Bars2Outline ...
func Bars2Outline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M3.75 9h16.5m-16.5 6.75h16.5"),
		),
	)
}

// Bars3Outline ...
func Bars3Outline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M3.75 5.25h16.5m-16.5 6.75h16.5m-16.5 6.75h16.5"),
		),
	)
}

// UserOutline ...
func UserOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z"),
		),
	)
}

// TrashOutline ...
func TrashOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0"),
		),
	)
}

// BriefcaseOutline ...
func BriefcaseOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M20.25 14.15v4.25c0 1.094-.787 2.036-1.872 2.18-2.087.277-4.216.42-6.378.42s-4.291-.143-6.378-.42c-1.085-.144-1.872-1.086-1.872-2.18v-4.25m16.5 0a2.18 2.18 0 0 0 .75-1.661V8.706c0-1.081-.768-2.015-1.837-2.175a48.114 48.114 0 0 0-3.413-.387m4.5 8.006c-.194.165-.42.295-.673.38A23.978 23.978 0 0 1 12 15.75c-2.648 0-5.195-.429-7.577-1.22a2.016 2.016 0 0 1-.673-.38m0 0A2.18 2.18 0 0 1 3 12.489V8.706c0-1.081.768-2.015 1.837-2.175a48.111 48.111 0 0 1 3.413-.387m7.5 0V5.25A2.25 2.25 0 0 0 13.5 3h-3a2.25 2.25 0 0 0-2.25 2.25v.894m7.5 0a48.667 48.667 0 0 0-7.5 0M12 12.75h.008v.008H12v-.008Z"),
		),
	)
}

// DocumentMagnifyingGlassOutline ...
func DocumentMagnifyingGlassOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m5.231 13.481L15 17.25m-4.5-15H5.625c-.621 0-1.125.504-1.125 1.125v16.5c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9Zm3.75 11.625a2.625 2.625 0 1 1-5.25 0 2.625 2.625 0 0 1 5.25 0Z"),
		),
	)
}

// MagnifyingGlassOutline ...
func MagnifyingGlassOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"),
		),
	)
}

// BellAlertOutline ...
func BellAlertOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0M3.124 7.5A8.969 8.969 0 0 1 5.292 3m13.416 0a8.969 8.969 0 0 1 2.168 4.5"),
		),
	)
}

// BellAlertOutlineSmall ...
func BellAlertOutlineSmall(p IconProps) htmx.Node {
	return BellAlertOutline(
		IconProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					"w-6 h-6": false,
					"w-5 h-5": true,
				},
			),
		},
	)
}

// SunOutline ...
func SunOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M12 3v2.25m6.364.386-1.591 1.591M21 12h-2.25m-.386 6.364-1.591-1.591M12 18.75V21m-4.773-4.227-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0Z"),
		),
	)
}

// SunOutlineSmall ...
func SunOutlineSmall(p IconProps) htmx.Node {
	return SunOutline(
		IconProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					"w-6 h-6": false,
					"w-5 h-5": true,
				},
			),
		},
	)
}

// MoonOutline ...
func MoonOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M21.752 15.002A9.72 9.72 0 0 1 18 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 0 0 3 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 0 0 9.002-5.998Z"),
		),
	)
}

// MoonOutlineSmall ...
func MoonOutlineSmall(p IconProps) htmx.Node {
	return MoonOutline(
		IconProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					"w-6 h-6": false,
					"w-5 h-5": true,
				},
			),
		},
	)
}

// InformationCircleOutline ...
func InformationCircleOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z"),
		),
	)
}

// PlusOutline ...
func PlusOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M12 4.5v15m7.5-7.5h-15"),
		),
	)
}

// CheckCircleOutline ...
func CheckCircleOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M9 12.75 11.25 15 15 9.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"),
		),
	)
}

// CheckOutline ...
func CheckOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m4.5 12.75 6 6 9-13.5"),
		),
	)
}

// ExclamationOutline ...
func ExclamationOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z"),
		),
	)
}

// ExclamationCircleOutline ...
func ExclamationCircleOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"),
		),
	)
}

// MinusOutline ...
func MinusOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M5 12h14"),
		),
	)
}

// MinusCircleOutline ...
func MinusCircleOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M15 12H9m12 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"),
		),
	)
}

// ArrowDownOnSquareOutline ...
func ArrowDownOnSquareOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6": true,
				"h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M9 8.25H7.5a2.25 2.25 0 0 0-2.25 2.25v9a2.25 2.25 0 0 0 2.25 2.25h9a2.25 2.25 0 0 0 2.25-2.25v-9a2.25 2.25 0 0 0-2.25-2.25H15M9 12l3 3m0 0 3-3m-3 3V2.25"),
		),
	)
}

// EllipsisHorizontalOutline ...
func EllipsisHorizontalOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M6.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM12.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM18.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z"),
		),
	)
}

// EllipsisVerticalOutline ...
func EllipsisVerticalOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M12 6.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 12.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 18.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5Z"),
		),
	)
}

// DocumentDuplicateOutline ...
func DocumentDuplicateOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				tailwind.H6: true,
				tailwind.W6: true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M15.75 17.25v3.375c0 .621-.504 1.125-1.125 1.125h-9.75a1.125 1.125 0 0 1-1.125-1.125V7.875c0-.621.504-1.125 1.125-1.125H6.75a9.06 9.06 0 0 1 1.5.124m7.5 10.376h3.375c.621 0 1.125-.504 1.125-1.125V11.25c0-4.46-3.243-8.161-7.5-8.876a9.06 9.06 0 0 0-1.5-.124H9.375c-.621 0-1.125.504-1.125 1.125v3.5m7.5 10.375H9.375a1.125 1.125 0 0 1-1.125-1.125v-9.25m12 6.625v-1.875a3.375 3.375 0 0 0-3.375-3.375h-1.5a1.125 1.125 0 0 1-1.125-1.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H9.75"),
		),
	)
}

// BoldOutline ...
func BoldOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				tailwind.H6: true,
				tailwind.W6: true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M6.75 3.744h-.753v8.25h7.125a4.125 4.125 0 0 0 0-8.25H6.75Zm0 0v.38m0 16.122h6.747a4.5 4.5 0 0 0 0-9.001h-7.5v9h.753Zm0 0v-.37m0-15.751h6a3.75 3.75 0 1 1 0 7.5h-6m0-7.5v7.5m0 0v8.25m0-8.25h6.375a4.125 4.125 0 0 1 0 8.25H6.75m.747-15.38h4.875a3.375 3.375 0 0 1 0 6.75H7.497v-6.75Zm0 7.5h5.25a3.75 3.75 0 0 1 0 7.5h-5.25v-7.5Z"),
		),
	)
}

// ItalicOutline ...
func ItalicOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				tailwind.H6: true,
				tailwind.W6: true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M5.248 20.246H9.05m0 0h3.696m-3.696 0 5.893-16.502m0 0h-3.697m3.697 0h3.803"),
		),
	)
}

// ItalicCodeBracketOutline ...
func ItalicCodeBracketOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				tailwind.H6: true,
				tailwind.W6: true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M17.25 6.75 22.5 12l-5.25 5.25m-10.5 0L1.5 12l5.25-5.25m7.5-3-4.5 16.5"),
		),
	)
}

// LinkOutline ...
func LinkOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				tailwind.H6: true,
				tailwind.W6: true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M13.19 8.688a4.5 4.5 0 0 1 1.242 7.244l-4.5 4.5a4.5 4.5 0 0 1-6.364-6.364l1.757-1.757m13.35-.622 1.757-1.757a4.5 4.5 0 0 0-6.364-6.364l-4.5 4.5a4.5 4.5 0 0 0 1.242 7.244"),
		),
	)
}

// ImageOutline ...
func ImageOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				tailwind.H6: true,
				tailwind.W6: true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m2.25 15.75 5.159-5.159a2.25 2.25 0 0 1 3.182 0l5.159 5.159m-1.5-1.5 1.409-1.409a2.25 2.25 0 0 1 3.182 0l2.909 2.909m-18 3.75h16.5a1.5 1.5 0 0 0 1.5-1.5V6a1.5 1.5 0 0 0-1.5-1.5H3.75A1.5 1.5 0 0 0 2.25 6v12a1.5 1.5 0 0 0 1.5 1.5Zm10.5-11.25h.008v.008h-.008V8.25Zm.375 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Z"),
		),
	)
}

// ListBulletOutline ...
func ListBulletOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				tailwind.H6: true,
				tailwind.W6: true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M8.25 6.75h12M8.25 12h12m-12 5.25h12M3.75 6.75h.007v.008H3.75V6.75Zm.375 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0ZM3.75 12h.007v.008H3.75V12Zm.375 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Zm-.375 5.25h.007v.008H3.75v-.008Zm.375 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Z"),
		),
	)
}

// NumberedListOutline ...
func NumberedListOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				tailwind.H6: true,
				tailwind.W6: true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M8.242 5.992h12m-12 6.003H20.24m-12 5.999h12M4.117 7.495v-3.75H2.99m1.125 3.75H2.99m1.125 0H5.24m-1.92 2.577a1.125 1.125 0 1 1 1.591 1.59l-1.83 1.83h2.16M2.99 15.745h1.125a1.125 1.125 0 0 1 0 2.25H3.74m0-.002h.375a1.125 1.125 0 0 1 0 2.25H2.99"),
		),
	)
}

// AtSymbolOutline ...
func AtSymbolOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				tailwind.H6: true,
				tailwind.W6: true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M16.5 12a4.5 4.5 0 1 1-9 0 4.5 4.5 0 0 1 9 0Zm0 0c0 1.657 1.007 3 2.25 3S21 13.657 21 12a9 9 0 1 0-2.636 6.364M16.5 12V8.25"),
		),
	)
}

// ArrowUturnLeftOutline ...
func ArrowUturnLeftOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				tailwind.H6: true,
				tailwind.W6: true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M9 15 3 9m0 0 6-6M3 9h12a6 6 0 0 1 0 12h-3"),
		),
	)
}

// StrikethroughOutline ...
func StrikethroughOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				tailwind.H6: true,
				tailwind.W6: true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M12 12a8.912 8.912 0 0 1-.318-.079c-1.585-.424-2.904-1.247-3.76-2.236-.873-1.009-1.265-2.19-.968-3.301.59-2.2 3.663-3.29 6.863-2.432A8.186 8.186 0 0 1 16.5 5.21M6.42 17.81c.857.99 2.176 1.812 3.761 2.237 3.2.858 6.274-.23 6.863-2.431.233-.868.044-1.779-.465-2.617M3.75 12h16.5"),
		),
	)
}
