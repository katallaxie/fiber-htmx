// Code generated by herogen; DO NOT EDIT.
package heroicons

import "github.com/katallaxie/fiber-htmx"

func EllipsisHorizontalCircleSolid(p IconProps) htmx.Node {
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
            htmx.Attribute("d", "M12 2.25C6.61522 2.25 2.25 6.61522 2.25 12C2.25 17.3848 6.61522 21.75 12 21.75C17.3848 21.75 21.75 17.3848 21.75 12C21.75 6.61522 17.3848 2.25 12 2.25ZM12 10.875C11.3787 10.875 10.875 11.3787 10.875 12C10.875 12.6213 11.3787 13.125 12 13.125C12.6213 13.125 13.125 12.6213 13.125 12C13.125 11.3787 12.6213 10.875 12 10.875ZM15.375 12C15.375 11.3787 15.8787 10.875 16.5 10.875C17.1213 10.875 17.625 11.3787 17.625 12C17.625 12.6213 17.1213 13.125 16.5 13.125C15.8787 13.125 15.375 12.6213 15.375 12ZM7.5 10.875C6.87868 10.875 6.375 11.3787 6.375 12C6.375 12.6213 6.87868 13.125 7.5 13.125C8.12132 13.125 8.625 12.6213 8.625 12C8.625 11.3787 8.12132 10.875 7.5 10.875Z"),
        ),
    )
}
