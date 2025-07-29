package htmx

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	goth "github.com/katallaxie/fiber-goth"
	"github.com/katallaxie/fiber-goth/adapters"
	reload "github.com/katallaxie/fiber-reload"
	"github.com/katallaxie/htmx"
	"github.com/katallaxie/pkg/conv"
	"gorm.io/gorm"
)

// Controller is the interface for the htmx controller.
type Controller interface {
	// Init is called when the controller is initialized.
	Init(ctx *fiber.Ctx) error
	// Prepare is called before the controller is executed.
	Prepare() error
	// Finalize is called after the controller is executed.
	Finalize() error
	// Get is called when the controller is executed with the GET method.
	Get() error
	// Post is called when the controller is executed with the POST method.
	Post() error
	// Put is called when the controller is executed with the PUT method.
	Put() error
	// Patch is called when the controller is executed with the PATCH method.
	Patch() error
	// Delete is called when the controller is executed with the DELETE method.
	Delete() error
	// Options is called when the controller is executed with the OPTIONS method.
	Options() error
	// Trace is called when the controller is executed with the TRACE method.
	Trace() error
	// Head  is called when the controller is executed with the HEAD method.
	Head() error
	// Error is called when an error occurs.
	Error(err error) error
	// Reset resets the controller.
	Reset()
	// Ctx returns the fiber.Ctx.
	Ctx() *fiber.Ctx
}

// TransactionController is the interface for a controller that also does database transactions.
type TransactionController interface {
	// Returns the transaction.
	Tx() *gorm.DB
	// Begin begins a transaction.
	Begin(*gorm.DB) error
	// Commit commits the transaction.
	Commmit() error
	// Rollback rolls back the transaction.
	Rollback() error
}

var _ Controller = (*DefaultController)(nil)

// UnimplementedController is not to be used anymore.
// Deprecated: Use DefaultController instead.
type UnimplementedController = DefaultController

// NewDefaultController returns a new default controller.
func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

// UnimplementedController is the default controller implementation.
type DefaultController struct {
	ctx *fiber.Ctx
}

// Init is called when the controller is initialized.
func (c *DefaultController) Init(ctx *fiber.Ctx) error {
	c.Reset()
	c.ctx = ctx

	return nil
}

// Prepare is called before the controller is executed.
func (c *DefaultController) Prepare() error {
	return nil
}

// Finalize is called after the controller is executed.
func (c *DefaultController) Finalize() error {
	return nil
}

// Get is called when the controller is executed with the GET method.
func (c *DefaultController) Get() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Post is called when the controller is executed with the POST method.
func (c *DefaultController) Post() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Put is called when the controller is executed with the PUT method.
func (c *DefaultController) Put() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Patch is called when the controller is executed with the PATCH method.
func (c *DefaultController) Patch() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Delete is called when the controller is executed with the DELETE method.
func (c *DefaultController) Delete() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Options is called when the controller is executed with the OPTIONS method.
func (c *DefaultController) Options() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Error is called when an error occurs.
func (c *DefaultController) Error(_ error) error {
	return fiber.NewError(fiber.StatusInternalServerError)
}

// Head is called when the controller is executed with the HEAD method.
func (c *DefaultController) Head() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Trace is called when the controller is executed with the TRACE method.
func (c *DefaultController) Trace() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// BindForm binds the form to the given struct.
func (c *DefaultController) BindBody(obj interface{}) error {
	return c.ctx.BodyParser(obj)
}

// BindParams binds the params to the given struct.
func (c *DefaultController) BindParams(obj interface{}) error {
	return c.ctx.ParamsParser(obj)
}

// BindQuery binds the query to the given struct.
func (c *DefaultController) BindQuery(obj interface{}) error {
	return c.ctx.QueryParser(obj)
}

// Values is a helper function to get the values from the context.
func (c *DefaultController) Values(key any, value ...any) (val any) {
	return c.ctx.Locals(key, value...)
}

// ValuesString is a helper function to get the values as a string from the context.
func (c *DefaultController) ValuesString(key any, value ...any) (val string) {
	return conv.String(c.ctx.Locals(key, value...))
}

// ValuesInt is a helper function to get the values as an int from the context.
func (c *DefaultController) ValuesInt(key any, value ...any) (val int) {
	return conv.Int(c.ctx.Locals(key, value...))
}

// ValuesBool is a helper function to get the values as a bool from the context.
func (c *DefaultController) ValuesBool(key any, value ...any) (val bool) {
	return conv.Bool(c.ctx.Locals(key, value...))
}

// Session is a helper function to get the session from the context.
func (c *DefaultController) Session() adapters.GothSession {
	session, err := goth.SessionFromContext(c.Ctx())
	if err != nil {
		return adapters.GothSession{}
	}

	return session
}

// Messages is a helper function to get the message from the context.
func (c *DefaultController) Messages(msgs ...Message) {
	m := MessagesFromContext(c.Ctx())
	m.Add(msgs...)
}

// Context returns the context.
func (c *DefaultController) Context() context.Context {
	return c.ctx.Context()
}

// Ctx returns the fiber.Ctx.
func (c *DefaultController) Ctx() *fiber.Ctx {
	return c.ctx
}

// IsDevelopment returns true if the environment is development.
func (c *DefaultController) IsDevelopment() bool {
	return reload.IsDevelopment(c.ctx.UserContext())
}

// IsProduction returns true if the environment is production.
func (c *DefaultController) IsProduction() bool {
	return reload.IsProduction(c.ctx.UserContext())
}

// Redirect redirects to the given path.
func (c *DefaultController) Redirect(path string) error {
	Redirect(c.ctx, path)

	return nil
}

// Render renders a component.
func (c *DefaultController) Render(node htmx.Node, opt ...RenderOpt) error {
	return RenderComp(c.ctx, node, opt...)
}

// Path returns the path.
func (c *DefaultController) Path() string {
	return c.ctx.Path()
}

// OriginalURL returns the original URL.
func (c *DefaultController) OriginalURL() string {
	return c.ctx.OriginalURL()
}

// Reset resets the controller.
func (c *DefaultController) Reset() {
	c.ctx = nil
}

var _ TransactionController = (*DefaultTransactionController)(nil)

// NewTransactionController returns a new transaction controller.
func NewTransactionController() *DefaultTransactionController {
	return &DefaultTransactionController{
		DefaultController: NewDefaultController(),
	}
}

// DefaultTransactionController is the interface for the htmx transaction controller.
type DefaultTransactionController struct {
	*DefaultController
	tx *gorm.DB
}

// Begin begins a transaction.
func (c *DefaultTransactionController) Begin(conn *gorm.DB) error {
	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	c.tx = tx

	return nil
}

// Commit commits the transaction.
func (c *DefaultTransactionController) Commmit() error {
	if c.tx == nil {
		return nil
	}

	c.tx.Commit()
	if err := c.tx.Error; err != nil {
		return err
	}

	c.tx = nil

	return nil
}

// Rollback rolls back the transaction.
func (c *DefaultTransactionController) Rollback() error {
	if c.tx == nil {
		return nil
	}

	c.tx.Rollback()
	if err := c.tx.Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}
	c.tx = nil

	return nil
}

// Txn returns the transaction.
func (c *DefaultTransactionController) Tx() *gorm.DB {
	if c.tx != nil {
		return c.tx
	}

	return nil
}
