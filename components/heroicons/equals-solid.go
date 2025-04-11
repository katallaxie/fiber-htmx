// Code generated by herogen; DO NOT EDIT.
package heroicons

import "github.com/katallaxie/fiber-htmx"

func EqualsSolid(p IconProps) htmx.Node {
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
            htmx.Attribute("d", "M3.74805 8.24809C3.74805 7.83385 4.08384 7.49805 4.49807 7.49805H19.4978C19.912 7.49805 20.2478 7.83385 20.2478 8.24809C20.2478 8.66232 19.912 8.99813 19.4978 8.99813H4.49807C4.08384 8.99813 3.74805 8.66232 3.74805 8.24809Z"),
        ),
        htmx.Path(
            htmx.Attribute("d", "M3.74805 15.7491C3.74805 15.3348 4.08384 14.999 4.49807 14.999H19.4978C19.912 14.999 20.2478 15.3348 20.2478 15.7491C20.2478 16.1633 19.912 16.4991 19.4978 16.4991H4.49807C4.08384 16.4991 3.74805 16.1633 3.74805 15.7491Z"),
        ),
    )
}