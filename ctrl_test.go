package htmx_test

import (
	"testing"

	htmx "github.com/katallaxie/fiber-htmx"
	"github.com/stretchr/testify/require"
)

func TestNewTransactionControl(t *testing.T) {
	ctrl := htmx.NewTransactionController()
	require.NotNil(t, ctrl)
	require.Implements(t, (*htmx.TransactionController)(nil), ctrl)
}
