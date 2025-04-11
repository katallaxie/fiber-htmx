// Code generated by herogen; DO NOT EDIT.
package heroicons

import "github.com/katallaxie/fiber-htmx"

func CubetransparentOutline(p IconProps) htmx.Node {
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
            htmx.Attribute("d", "M21 7.5L18.75 6.1875M21 7.5V9.75M21 7.5L18.75 8.8125M3 7.5L5.25 6.1875M3 7.5L5.25 8.8125M3 7.5V9.75M12 12.75L14.25 11.4375M12 12.75L9.75 11.4375M12 12.75V15M12 21.75L14.25 20.4375M12 21.75V19.5M12 21.75L9.75 20.4375M9.75 3.5625L12 2.25L14.25 3.5625M21 14.25V16.5L18.75 17.8125M5.25 17.8125L3 16.5V14.25"),
        ),
    )
}