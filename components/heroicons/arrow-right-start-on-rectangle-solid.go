// Code generated by herogen; DO NOT EDIT.
package heroicons

import htmx "github.com/katallaxie/fiber-htmx"

func ArrowRightStartOnRectangleSolid(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Merge(
			htmx.ClassNames{
				"h-5": false,
				"h-6": true,
				"w-5": false,
				"w-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("d", "M7.5 3.75C6.67157 3.75 6 4.42157 6 5.25L6 18.75C6 19.5784 6.67157 20.25 7.5 20.25H13.5C14.3284 20.25 15 19.5784 15 18.75V15C15 14.5858 15.3358 14.25 15.75 14.25C16.1642 14.25 16.5 14.5858 16.5 15V18.75C16.5 20.4069 15.1569 21.75 13.5 21.75H7.5C5.84315 21.75 4.5 20.4069 4.5 18.75L4.5 5.25C4.5 3.59315 5.84315 2.25 7.5 2.25L13.5 2.25C15.1569 2.25 16.5 3.59315 16.5 5.25V9C16.5 9.41421 16.1642 9.75 15.75 9.75C15.3358 9.75 15 9.41421 15 9V5.25C15 4.42157 14.3284 3.75 13.5 3.75L7.5 3.75ZM18.2197 8.46967C18.5126 8.17678 18.9874 8.17678 19.2803 8.46967L22.2803 11.4697C22.5732 11.7626 22.5732 12.2374 22.2803 12.5303L19.2803 15.5303C18.9874 15.8232 18.5126 15.8232 18.2197 15.5303C17.9268 15.2374 17.9268 14.7626 18.2197 14.4697L19.9393 12.75L9 12.75C8.58579 12.75 8.25 12.4142 8.25 12C8.25 11.5858 8.58579 11.25 9 11.25L19.9393 11.25L18.2197 9.53033C17.9268 9.23744 17.9268 8.76256 18.2197 8.46967Z"),
		),
	)
}
