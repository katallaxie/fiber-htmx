package alpine

import htmx "github.com/katallaxie/fiber-htmx"

// XData returns an attribute that tells Alpine.js to ignore an element.
func XData(data string) htmx.Node {
	return htmx.Attribute("x-data", data)
}

// XInit returns an attribute that tells Alpine.js to run a function when an element is initialized.
func XInit(data string) htmx.Node {
	return htmx.Attribute("x-init", data)
}

// XShow returns an attribute that tells Alpine.js to show an element.
func XShow(data string) htmx.Node {
	return htmx.Attribute("x-show", data)
}

// XBind returns an attribute that tells Alpine.js to bind a value to an element.
func XBind(selector string, data string) htmx.Node {
	return htmx.Attribute("x-bind:"+selector, data)
}

// XOn returns an attribute that tells Alpine.js to listen for an event on an element.
func XOn(selector string, data string) htmx.Node {
	return htmx.Attribute("x-on:"+selector, data)
}

// XText returns an attribute that tells Alpine.js to bind a value to an element's text content.
func XText(data string) htmx.Node {
	return htmx.Attribute("x-text", data)
}

// XHtml returns an attribute that tells Alpine.js to bind a value to an element's inner HTML.
func XHtml(data string) htmx.Node {
	return htmx.Attribute("x-html", data)
}

// XModel returns an attribute that tells Alpine.js to bind a value to an input element.
func XModel(data string) htmx.Node {
	return htmx.Attribute("x-model", data)
}

// XModelable returns an attribute that tells Alpine.js to bind a value to an input element.
func XModelable(data string) htmx.Node {
	return htmx.Attribute("x-modelable", data)
}

// XFor returns an attribute that tells Alpine.js to loop over an array.
func XFor(data string) htmx.Node {
	return htmx.Attribute("x-for", data)
}

// XTransition returns an attribute that tells Alpine.js to apply a transition to an element.
func XTransition(customize string) htmx.Node {
	if customize == "" {
		return htmx.Attribute("x-transition", "")
	}

	return htmx.Attribute("x-transition:"+customize, "")
}

// XEffect returns an attribute that tells Alpine.js to apply an effect to an element.
func XEffect(customize string) htmx.Node {
	return htmx.Attribute("x-effect", customize)
}

// XIgnore returns an attribute that tells Alpine.js to ignore an element.
func XIgnore() htmx.Node {
	return htmx.Attribute("x-ignore", "")
}

// XRef returns an attribute that tells Alpine.js to reference an element.
func XRef(name string) htmx.Node {
	return htmx.Attribute("x-ref", name)
}

// XCloak returns an attribute that tells Alpine.js to hide an element until it is initialized.
func XCloak() htmx.Node {
	return htmx.Attribute("x-cloak", "")
}

// XTeleport returns an attribute that tells Alpine.js to move an element to a different location in the DOM.
func XTeleport(data string) htmx.Node {
	return htmx.Attribute("x-teleport", data)
}

// XIf returns an attribute that tells Alpine.js to conditionally show an element.
func XIf(data string) htmx.Node {
	return htmx.Attribute("x-if", data)
}

// XId returns an attribute that tells Alpine.js to set an element's ID.
func XId(data string) htmx.Node {
	return htmx.Attribute("x-id", data)
}
