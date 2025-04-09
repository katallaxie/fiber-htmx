package dates

import (
	"time"

	htmx "github.com/katallaxie/fiber-htmx"
)

// DateTextProps is a struct that holds the properties of a date
type DateTextProps struct {
	// Format is a string that holds the format of a date
	Format string
}

// DateText ...
func DateText(props DateTextProps, date time.Time) htmx.Node {
	if props.Format == "" {
		props.Format = time.RFC822
	}

	return htmx.Text(date.Format(props.Format))
}
