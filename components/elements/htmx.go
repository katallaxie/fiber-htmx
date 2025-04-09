package elements

import htmx "github.com/katallaxie/fiber-htmx"

// HTMXToasts is a component that displays toasts.
func HTMXToasts() htmx.Node {
	return htmx.Element("htmx-toasts")
}

// GithubMarkdownToolbar ...
func GithubMarkdownToolbar(children ...htmx.Node) htmx.Node {
	return htmx.Element("markdown-toolbar", children...)
}

// GithubMarkdownBold ...
func GithubMarkdownBold(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-bold", children...)
}

// GithubMarkdownHeader ...
func GithubMarkdownHeader(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-header", children...)
}

// GithubMarkdownItalic ...
func GithubMarkdownItalic(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-italic", children...)
}

// GithubMarkdownQuote ...
func GithubMarkdownQuote(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-quote", children...)
}

// GithubMarkdownCode ...
func GithubMarkdownCode(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-code", children...)
}

// GithubMarkdownLink ...
func GithubMarkdownLink(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-link", children...)
}

// GithubMarkdownImage ...
func GithubMarkdownImage(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-image", children...)
}

// GithubMarkdownUnorderedList ...
func GithubMarkdownUnorderedList(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-unordered-list", children...)
}

// GithubMarkdownOrderedList ...
func GithubMarkdownOrderedList(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-ordered-list", children...)
}

// GithubMarkdownTaskList ...
func GithubMarkdownTaskList(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-task-list", children...)
}

// GithubMarkdownMention ...
func GithubMarkdownMention(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-mention", children...)
}

// GithubMarkdownRef ...
func GithubMarkdownRef(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-ref", children...)
}

// GithubMarkdownStrikethrough ...
func GithubMarkdownStrikethrough(children ...htmx.Node) htmx.Node {
	return htmx.Element("md-strikethrough", children...)
}
