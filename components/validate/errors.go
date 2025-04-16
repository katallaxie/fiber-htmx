package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/katallaxie/pkg/slices"
	"github.com/katallaxie/pkg/utilx"
)

// Errors is a map of field names to error messages.
type Errors validator.ValidationErrors

// Field returns the error message for the provided field.
func (e Errors) Field(name string) error {
	for _, err := range e {
		if err.Field() == name {
			return err
		}
	}

	return nil
}

// HasError returns true if the field has an error.
func (e Errors) HasError(name string) bool {
	return utilx.NotEmpty(e.Field(name))
}

// TagNameFunc returns the tag name for the provided field.
func TagNameFunc(fld reflect.StructField) string {
	name := slices.First(strings.SplitN(fld.Tag.Get("json"), ",", 2)...) //nolint:mnd

	if name == "-" {
		return ""
	}

	return name
}
