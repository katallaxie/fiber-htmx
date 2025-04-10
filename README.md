# 🔨 HTMX

[![Test & Build](https://github.com/katallaxie/fiber-htmx/actions/workflows/main.yml/badge.svg)](https://github.com/katallaxie/fiber-htmx/actions/workflows/main.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/katallaxie/fiber-htmx.svg)](https://pkg.go.dev/github.com/katallaxie/fiber-htmx)
[![Go Report Card](https://goreportcard.com/badge/github.com/katallaxie/fiber-htmx)](https://goreportcard.com/report/github.com/katallaxie/fiber-htmx)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)

A Go package to write HTML5 and HTMX components in Go. The package is designed to work with [fiber](http://gofiber.io) and [htmx](https://htmx.org/).

## Features

- Write declartive HTML5 components in Go without using templates and with the full-power of a type-safe language, auto-completion, and refactoring.
- Full support for HTMX components.
- No dependencies on JavaScript frameworks.
- Fast rendering of HTML5 and HTMX components.
- Easy to use and learn.
- Easy to extend and customize.

## Installation

```bash
go get github.com/katallaxie/fiber-htmx
```

The available web components are published in the `fiber-htmx` package.

```go
htmx.HTML5(
    htmx.HTML5Props{
        Head: []htmx.Node{
            htmx.Script(
                htmx.Src("https://unpkg.com/fiber-htmx@1.3.32"),
            ),
            htmx.Link(
                htmx.Rel("stylesheet"),
                htmx.Href("https://unpkg.com/fiber-htmx@1.3.32/dist/out.css"),
            ),
        },
    },
    htmx.Body(
        htmx.ClassNames{},
        htmx.Toasts(),
    ),
)
```

### Example

Creating a button leveraging htmx is as easy as this.

```go
htmx.Button(
    htmx.Attribute("type", "submit")
    htmx.Text("Button"),
    htmx.HxPost("/api/respond")
)
```

There are additional complex components that help to write HTML5 and HTMX components in Go.

- [x] [htmx](https://htmx.org/)
- [x] [HTML5](https://www.w3.org/TR/2011/WD-html5-20110405/)
- [x] [TailwindCSS](https://tailwindcss.com/)
- [x] [Alpine.js](https://alpinejs.dev/)
- [ ] [DaisyUI](https://daisyui.com/) (WIP)
- [ ] [Heroicons](https://heroicons.com/) (WIP)

## Elements

HTML and HTMX elements are represented as functions in Go. The functions are used to create the elements.

```go
htmx.Div(
    htmx.ClassNames{
        tailwind.FontSemibold: true,
    },
    htmx.Text("Hello World"),
)
```

This will create the following HTML element.

```html
<div class="font-semibold">Hello World</div>
```

There is support for all HTML5 elements and Tailwind classes. Use `import "github.com/katallaxie/fiber-htmx/tailwind"` to include Tailwind classes.

## Components

Write HTML5 and HTMX components in Go.

```go
func HelloWorld() htmx.Node {
    return htmx.Div(
        htmx.ClassNames{
            "font-semibold",
        },
        htmx.Text("Hello World"),
    )
}
```

There are different types of composition. For example, passing children to a component.

```go
func HelloWorld(children ...htmx.Node) htmx.Node {
    return htmx.Div(
        htmx.ClassNames{
            "font-semibold",
        },
        htmx.Text("Hello World"),
        htmx.Div(
            htmx.ClassNames{
                "text-red-500",
            },
            htmx.Group(children...),
        ),
    )
}
```

Styling of components is done with the `htmx.ClassNames` type.

```go
func HelloWorld() htmx.Node {
    return htmx.Div(
        htmx.ClassNames{
            tailwind.FontSemibold: true,
            "text-red-500": true,
        },
        htmx.Text("Hello World"),
    )
}
```

There are also helpers to make the life with styling easier by merging classes.

```go
func HelloWorld(classes htmx.ClassNames) htmx.Node {
    return htmx.Div(
        htmx.Merge(
            htmx.ClassNames{
                "font-semibold",
                "text-red-500",
            },
            classes,
        )
        htmx.Text("Hello World"),
    )
}
```

There is alos another pattern to create a component. This enables you to track state and to use the component in a more declarative way.

```go
type Page struct {
	htmx.Node
}

func NewPage(title, body string) *Page {
	return &Page{
		Node: htmx.HTML5(
			htmx.HTML5Props{
				Title: title,
			},
			htmx.Body(
				htmx.Div(
					htmx.Text(body),
				),
			),
		),
	}
}
```

There is also the option to use `htmx.Controller` to encapsulate the logic of the components.

```go

func NewHelloWorldController() htmx.ControllerFactory {
  return func() htmx.Controller {
    return &NewHelloWorldController{}
  }
}

type HelloWorldController struct {
    htmx.DefaultController
}

func (c *HelloWorldController) Get() error {
    return c.Render(
      htmx.HTML5(
            htmx.HTML5Props{
                Title:    "index",
                Language: "en",
                Head: []htmx.Node{},
            },
            htmx.Div(
                htmx.ClassNames{},
                htmx.Text("Hello World"),
            ),
        ),
    )    
}

app := fiber.New()
app.Get("/", htmx.NewHxControllerHandler(NewHelloWorldController()))

app.Listen(":3000")
```

## Server-side events (SSE)

The package supports server-side events (SSE) to update the components on the client-side.

```go
manager := sse.NewBroadcastManager(5)
app.Get("/sse", sse.NewSSEHandler(manager))
```
## Error Handling

There are components that enable handling fallbacks in case of errors and to recover from panics in rendering components.

```go
htmx.Fallback(
  htmx.ErrorBoundary(
    func() htmx.Node {
      return utils.Panic(errors.New("panic"))
    },
  ),
  htmx.Text("Fallback"),
),
```

## Examples

See [examples](https://github.com/katallaxie/fiber-htmx/tree/master/examples) to understand the provided interfaces.

## Benchmarks

```bash
BenchmarkElement-10                     12964440                77.40 ns/op
Benchmark_AttrRender-10                 16038232                74.15 ns/op
Benchmark_HTML5_Render-10                   1392            847193 ns/op
Benchmark_ClassNames_Render-10           3166761               378.2 ns/op
```

Rendering `10.000` nodes took `>0.8ms`. The package is fast enough to render HTML5 and HTMX components. 

## License

[MIT](/LICENSE)
