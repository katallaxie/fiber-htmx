// Code generated by herogen; DO NOT EDIT.
package heroicons

import "github.com/katallaxie/fiber-htmx"

func ArrowDownTrayOutline(p IconProps) htmx.Node {
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
            htmx.Attribute("d", "M3 16.5V18.75C3 19.9926 4.00736 21 5.25 21H18.75C19.9926 21 21 19.9926 21 18.75V16.5M16.5 12L12 16.5M12 16.5L7.5 12M12 16.5V3"),
        ),
    )
}