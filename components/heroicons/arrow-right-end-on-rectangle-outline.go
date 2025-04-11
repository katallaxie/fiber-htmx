// Code generated by herogen; DO NOT EDIT.
package heroicons

import "github.com/katallaxie/fiber-htmx"

func ArrowRightEndOnRectangleOutline(p IconProps) htmx.Node {
    return htmx.SVG(
        htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
        htmx.Attribute("viewBox", "0 0 24 24"),
        htmx.Merge(
            htmx.ClassNames{
                "h-5": true,
                "h-6": false,
                "w-5": true,
                "w-6": false,
            },
            p.ClassNames,
        ),
        htmx.Path(
            htmx.Attribute("stroke-linecap", "round"),
            htmx.Attribute("stroke-linejoin", "round"),
            htmx.Attribute("d", "M8.25 9V5.25C8.25 4.00736 9.25736 3 10.5 3L16.5 3C17.7426 3 18.75 4.00736 18.75 5.25L18.75 18.75C18.75 19.9926 17.7426 21 16.5 21H10.5C9.25736 21 8.25 19.9926 8.25 18.75V15M12 9L15 12M15 12L12 15M15 12L2.25 12"),
        ),
    )
}