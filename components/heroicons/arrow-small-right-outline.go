// Code generated by herogen; DO NOT EDIT.
package heroicons

import "github.com/katallaxie/fiber-htmx"

func ArrowSmallRightOutline(p IconProps) htmx.Node {
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
            htmx.Attribute("d", "M4.5 12L19.5 12M19.5 12L12.75 5.25M19.5 12L12.75 18.75"),
        ),
    )
}
