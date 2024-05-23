package tables

import (
	"fmt"

	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/utils"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

// DefaultLimits is a list of default limits
var DefaultLimits = []int{5, 10, 25, 50}

// PaginationProps is a struct that contains the properties of a pagination
type PaginationProps struct {
	ClassNames htmx.ClassNames
	Limit      int
	Offset     int
	Target     string
	Total      int
	URL        string
}

// Pagination ...
func Pagination(p PaginationProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"join": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Prev ...
func Prev(p PaginationProps) htmx.Node {
	return htmx.Form(
		htmx.Method("GET"),
		htmx.Action(p.URL),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("offset"),
			htmx.Value(fmt.Sprintf("%d", p.Offset-p.Limit)),
		),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("limit"),
			htmx.Value(fmt.Sprintf("%d", p.Limit)),
		),
		htmx.HxBoost(true),
		buttons.Button(
			buttons.ButtonProps{
				ClassNames: htmx.ClassNames{
					"join-item":      true,
					"btn":            true,
					"btn-outline":    true,
					"input-bordered": true,
				},
				Type: "submit",
			},
			htmx.If(p.Offset-p.Limit < 0, htmx.Disabled()),
			htmx.Text("Prev"),
		),
	)
}

// Next ...
func Next(p PaginationProps) htmx.Node {
	return htmx.Form(
		htmx.Method("GET"),
		htmx.Action(p.URL),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("offset"),
			htmx.Value(fmt.Sprintf("%d", p.Offset+p.Limit)),
		),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("limit"),
			htmx.Value(fmt.Sprintf("%d", p.Limit)),
		),
		htmx.HxBoost(true),
		buttons.Button(
			buttons.ButtonProps{
				ClassNames: htmx.ClassNames{
					"join-item":      true,
					"btn":            true,
					"btn-outline":    true,
					"input-bordered": true,
				},
				Type: "submit",
			},
			htmx.If(p.Offset+p.Limit > p.Total, htmx.Disabled()),
			htmx.Text("Next"),
		),
	)
}

// SelectProps ...
type SelectProps struct {
	ClassNames htmx.ClassNames
	Limit      int
	Limits     []int
	Offset     int
	Target     string
	Total      int
	URL        string
}

// Select ...
func Select(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Form(
		htmx.Method("GET"),
		htmx.Action(p.URL),
		htmx.HxTrigger("change from:#select-table-options"),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("offset"),
			htmx.Value(fmt.Sprintf("%d", p.Offset)),
		),
		htmx.HxBoost(true),
		forms.Select(
			forms.SelectProps{
				ClassNames: htmx.Merge(
					htmx.ClassNames{
						"join-item":      true,
						"input-bordered": true,
					},
					p.ClassNames,
				),
			},
			htmx.ID("select-table-options"),
			htmx.Attribute("name", "limit"),
			utils.Map(func(limit int) htmx.Node {
				return forms.Option(
					forms.OptionProps{
						Selected: limit == p.Limit,
					},
					htmx.Text(fmt.Sprintf("%d", limit)),
					htmx.Value(fmt.Sprintf("%d", limit)),
				)
			}, p.Limits...),
		),
	)
}

// TableToolbarProps is a struct that contains the properties of a table toolbar
type TableToolbarProps struct {
	ClassNames htmx.ClassNames
}

// TableToolbar is a component that renders a table toolbar
func TableToolbar(p TableToolbarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"table-toolbar": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// TablePaginationProps is a struct that contains the properties of a table pagination
type TablePaginationProps struct {
	ClassNames htmx.ClassNames
	Pagination htmx.Node
}

// TablePagination is a component that renders a table pagination
func TablePagination(p TablePaginationProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"flex":            true,
				"items-center":    true,
				"justify-between": true,
				"px-2":            true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.Merge(
				htmx.ClassNames{
					"flex":         true,
					"items-center": true,
					"space-x-6":    true,
					"lg:space-x-8": true,
				},
			),
			htmx.P(
				htmx.Merge(
					htmx.ClassNames{
						"text-sm":     true,
						"font-medium": true,
					},
				),
				htmx.Text("Rows per page"),
			),
			p.Pagination,
		),
	)
}

// Row is a struct that contains the properties of a row
type Row interface {
	comparable
}

// TableProps is a struct that contains the properties of a table
type TableProps struct {
	ClassNames htmx.ClassNames
	ID         string
	Pagination htmx.Node
	Toolbar    htmx.Node
}

// Columns returns a new column definition.
type Columns[R Row] []ColumnDef[R]

// ColumnDef returns a new column definition.
type ColumnDef[R Row] struct {
	ID              string
	AccessorKey     string
	Header          func(p TableProps) htmx.Node
	Cell            func(p TableProps, row R) htmx.Node
	EnableSorting   bool
	EnableFiltering bool
}

// Table is a struct that contains the properties of a table
func Table[S ~[]R, R Row](p TableProps, columns Columns[R], s S) htmx.Node {
	headers := []htmx.Node{}
	for _, column := range columns {
		headers = append(headers, column.Header(p))
	}

	rows := []htmx.Node{}
	for _, row := range s {
		cells := []htmx.Node{}
		for _, column := range columns {
			cells = append(cells, column.Cell(p, row))
		}
		rows = append(rows, htmx.Tr(cells...))

	}

	return htmx.Div(
		htmx.ID(p.ID),
		htmx.Merge(
			htmx.ClassNames{
				"space-y-4": true,
			},
		),
		p.Toolbar,
		htmx.Div(
			htmx.Merge(),
			htmx.Table(
				htmx.Merge(
					htmx.ClassNames{
						"table": true,
					},
					p.ClassNames,
				),
				htmx.THead(
					htmx.Tr(
						headers...,
					),
				),
				htmx.TBody(
					rows...,
				),
			),
		),
		p.Pagination,
	)
}

// PaginationFromContext returns a new Pagination object based on the provided context.
func PaginationPropsFromContext(c *fiber.Ctx) (limit int, offset int) {
	limit = c.QueryInt("limit", 10)
	offset = c.QueryInt("offset", 0)

	return
}
