package htmx

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/katallaxie/pkg/errorx"
	"github.com/yuin/goldmark"
)

// Node is a node in the HTML tree.
type Node interface {
	Render(w io.Writer) error
}

// Nodes is a slice of nodes.
type Nodes []Node

// NodeType is the type of a node.
type NodeType int

const (
	ElementType = NodeType(iota)
	AttributeType
)

// NodeTypeDescriptor is a node that has a type.
type NodeTypeDescriptor interface {
	Type() NodeType
}

// NodeFunc is a function that renders a node.
type NodeFunc func(io.Writer) error

// Render renders the node.
func (n NodeFunc) Render(w io.Writer) error {
	return n(w)
}

// Type returns the node type.
func (n NodeFunc) Type() NodeType {
	return ElementType
}

// String returns the node as a string.
func (n NodeFunc) String() string {
	var b strings.Builder

	_ = n.Render(&b)

	return b.String()
}

// Element is a node that renders an HTML element.
func Element(name string, children ...Node) NodeFunc {
	return NodeFunc(func(w2 io.Writer) error {
		w := &statefulWriter{w: w2}

		w.Write([]byte("<" + name))

		for _, c := range children {
			renderChild(w, c, AttributeType)
		}

		w.Write([]byte(">"))

		if isVoidElement(name) {
			return w.err
		}

		for _, c := range children {
			renderChild(w, c, ElementType)
		}

		w.Write([]byte("</" + name + ">"))

		return w.err
	})
}

//nolint:gocyclo
func renderChild(w *statefulWriter, c Node, t NodeType) {
	if w.err != nil || c == nil {
		return
	}

	if g, ok := c.(group); ok {
		for _, groupC := range g.children {
			renderChild(w, groupC, t)
		}
		return
	}

	switch t {
	case ElementType:
		if p, ok := c.(NodeTypeDescriptor); !ok || p.Type() == ElementType {
			w.err = c.Render(w.w)
		}
	case AttributeType:
		if p, ok := c.(NodeTypeDescriptor); ok && p.Type() == AttributeType {
			w.err = c.Render(w.w)
		}
	}
}

type statefulWriter struct {
	w   io.Writer
	err error
}

// Write is a node that writes to the stateful writer.
func (w *statefulWriter) Write(p []byte) {
	if w.err != nil {
		return
	}
	_, w.err = w.w.Write(p)
}

var voidElements = map[string]struct{}{
	"area":    {},
	"base":    {},
	"br":      {},
	"col":     {},
	"command": {},
	"embed":   {},
	"hr":      {},
	"img":     {},
	"input":   {},
	"keygen":  {},
	"link":    {},
	"meta":    {},
	"param":   {},
	"source":  {},
	"track":   {},
	"wbr":     {},
}

func isVoidElement(name string) bool {
	_, ok := voidElements[name]
	return ok
}

// Attribute is a node that renders an HTML attribute.
func Attribute(name string, value ...string) Node {
	switch len(value) {
	case 0:
		return &attr{name: name}
	case 1:
		return &attr{name: name, value: &value[0]}
	default:
		panic("attribute must be just name or name and value pair")
	}
}

type attr struct {
	name  string
	value *string
}

// Render is a node that renders an attribute.
func (a *attr) Render(w io.Writer) error {
	if a.value == nil {
		_, err := w.Write([]byte(" " + a.name))
		return err
	}

	_, err := w.Write([]byte(" " + a.name + `="` + template.HTMLEscapeString(*a.value) + `"`))

	return err
}

// Type is a node that returns the type of an attribute.
func (a *attr) Type() NodeType {
	return AttributeType
}

// String is a node that returns the attribute as a string.
func (a *attr) String() string {
	var b strings.Builder

	_ = a.Render(&b)

	return b.String()
}

// Text is a node that renders a text.
func Text(t string) Node {
	return NodeFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(template.HTMLEscapeString(t)))

		return err
	})
}

// Textf is a node that renders a formatted text.
func Textf(format string, a ...interface{}) Node {
	return NodeFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(template.HTMLEscapeString(fmt.Sprintf(format, a...))))

		return err
	})
}

// Raw is a node that renders raw HTML.
func Raw(t string) Node {
	return NodeFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(t))

		return err
	})
}

// UnsafeRaw is a node that renders an unsafe raw HTML.
func UnsafeRaw(t string) Node {
	return Raw(t)
}

// Rawf is a node that renders a formatted raw HTML.
func Rawf(format string, a ...interface{}) Node {
	return NodeFunc(func(w io.Writer) error {
		_, err := fmt.Fprintf(w, format, a...)

		return err
	})
}

// UnsafeRawf is a node that renders an unsafe formatted raw HTML.
func UnsafeRawf(format string, a ...interface{}) Node {
	return Rawf(format, a...)
}

// RawScript is a node that renders a raw script.
func RawScript(t string) Node {
	return NodeFunc(func(w io.Writer) error {
		_, err := w.Write([]byte("<script>" + t + "</script>"))

		return err
	})
}

// UnsafeRawScript is a node that renders an unsafe raw script.
func UnsafeRawScript(t string) Node {
	return RawScript(t)
}

// Empty is a node that renders an empty string.
func Empty() Node {
	return Text("")
}

type group struct {
	children []Node
}

// String is a node that renders a group of nodes.
func (c group) String() string {
	panic("cannot render group directly")
}

// Render is a node that renders a group of nodes.
func (c group) Render(io.Writer) error {
	panic("cannot render children directly")
}

// Group is a node that groups children nodes.
func Group(children ...Node) Node {
	return group{children: children}
}

type fragment struct {
	children []Node
}

// Fragment is a node that renders a fragment of nodes.
func Fragment(children ...Node) Node {
	return fragment{children: children}
}

// String is a node that renders a fragment of nodes.
func (c fragment) String() string {
	var b strings.Builder

	for _, child := range c.children {
		_ = child.Render(&b)
	}

	return b.String()
}

// Render is a node that renders a fragment of nodes.
func (c fragment) Render(w io.Writer) error {
	for _, child := range c.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}

	return nil
}

type errorBoundary struct {
	n ErrBoundaryFunc
}

// ErrBoundaryFunc is a function that returns a node.
type ErrBoundaryFunc func() Node

// ErrorBoundary is a node that catches panics and renders an error.
func ErrorBoundary(n ErrBoundaryFunc) Node {
	return errorBoundary{n: n}
}

// Render is a node that renders an error boundary.
func (c errorBoundary) Render(w io.Writer) error {
	n := c.n()

	return n.Render(w)
}

type fallback struct {
	n Node
	f FallbackFunc
}

// FallbackFunc is a function that returns a node.
type FallbackFunc func(error) Node

// Fallback is a node that renders a fallback node if a condition is false.
func Fallback(n Node, f FallbackFunc) Node {
	return fallback{n: n, f: f}
}

// Render is a node that renders a fallback node.
func (c fallback) Render(w io.Writer) (err error) {
	if c.n == nil {
		n := c.f(nil)
		return n.Render(w)
	}

	defer func() {
		if r := recover(); r != nil {
			n := c.f(errorx.RecoverError(r))
			err = n.Render(w)
		}
	}()

	var b bytes.Buffer

	if err := c.n.Render(&b); err != nil {
		return c.f(err).Render(w)
	}

	_, err = io.Copy(w, &b)

	return err
}

// If is a node that renders a child node if a condition is true.
func If(condition bool, n Node) Node {
	if condition {
		return n
	}

	return nil
}

// IfElse is a node that renders a child node if a condition is true, otherwise it renders another child node.
func IfElse(condition bool, n, elseN Node) Node {
	if condition {
		return n
	}

	return elseN
}

// Errors is a slice of errors.
type Errors[K comparable] map[K]error

// Error is a node that renders an error.
func FromValidationErrors[K comparable](errr []validator.FieldError) Errors[K] {
	ee := make(Errors[K])

	for _, err := range errr {
		if k, ok := any(err).(K); ok {
			ee[k] = err
		}
	}

	return ee
}

// KeyExists is a node that renders a child node if a key exists in a map.
func KeyExists[K comparable, V any](m map[K]V, key K, fn func(k K, v V) Node) Node {
	if v, ok := m[key]; ok {
		return fn(key, v)
	}

	return nil
}

// ErrorExists is a node that renders a child node if an error exists.
func ErrorExists[K comparable](e Errors[K], key K, fn func(k K, v error) Node) Node {
	return KeyExists(e, key, fn)
}

type md struct {
	source []byte
	opts   []goldmark.Option
}

// Markdown is a node that renders a markdown.
func Markdown(source []byte, opts ...goldmark.Option) Node {
	return md{source: source, opts: opts}
}

// Render is a node that renders a markdown.
func (m md) Render(w io.Writer) error {
	md := goldmark.New(m.opts...)

	return md.Convert(m.source, w)
}
