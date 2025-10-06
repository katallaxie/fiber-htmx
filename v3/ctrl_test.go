package htmx_test

import (
	"testing"

	htmx "github.com/katallaxie/fiber-htmx/v3"

	"github.com/stretchr/testify/require"
)

func TestUnimplementedTransactionController(t *testing.T) {
	ctrl := htmx.UnimplementedTransactionController{}
	require.NotNil(t, ctrl)
	require.Implements(t, (*htmx.TransactionController)(nil), &ctrl)
}
