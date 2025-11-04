package htmx

import (
	"context"
	"database/sql"
	"errors"

	"github.com/katallaxie/htmx"

	"github.com/gofiber/fiber/v3"
	goth "github.com/katallaxie/fiber-goth/v3"
	"github.com/katallaxie/fiber-goth/v3/adapters"
	reload "github.com/katallaxie/fiber-reload/v3"
	"github.com/katallaxie/pkg/conv"
	"gorm.io/gorm"
)

// ErrUnimplemented is returned when a method is not implemented.
var ErrUnimplemented = fiber.NewError(fiber.StatusNotImplemented, "method not implemented")

// Controller is the interface for the htmx controller.
type Controller interface {
	// Init is called when the controller is initialized.
	Init(c fiber.Ctx) error
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
	// Reset resets the controller.
	Reset()
	// Clone returns a new instance of the controller.
	Clone() Controller
}

var _ Controller = (*UnimplementedController)(nil)

// UnimplementedController is the default controller implementation.
type UnimplementedController struct {
	ctx fiber.Ctx
}

// Init is called when the controller is initialized.
func (c *UnimplementedController) Init(ctx fiber.Ctx) error {
	c.Reset()
	c.ctx = ctx

	return nil
}

// Prepare is called before the controller is executed.
func (c *UnimplementedController) Prepare() error {
	return nil
}

// Finalize is called after the controller is executed.
func (c *UnimplementedController) Finalize() error {
	return nil
}

// Get is called when the controller is executed with the GET method.
func (c *UnimplementedController) Get() error {
	return ErrUnimplemented
}

// Post is called when the controller is executed with the POST method.
func (c *UnimplementedController) Post() error {
	return ErrUnimplemented
}

// Put is called when the controller is executed with the PUT method.
func (c *UnimplementedController) Put() error {
	return ErrUnimplemented
}

// Patch is called when the controller is executed with the PATCH method.
func (c *UnimplementedController) Patch() error {
	return ErrUnimplemented
}

// Delete is called when the controller is executed with the DELETE method.
func (c *UnimplementedController) Delete() error {
	return ErrUnimplemented
}

// Options is called when the controller is executed with the OPTIONS method.
func (c *UnimplementedController) Options() error {
	return ErrUnimplemented
}

// Head is called when the controller is executed with the HEAD method.
func (c *UnimplementedController) Head() error {
	return ErrUnimplemented
}

// Trace is called when the controller is executed with the TRACE method.
func (c *UnimplementedController) Trace() error {
	return ErrUnimplemented
}

// BindForm binds the form to the given struct.
func (c *UnimplementedController) BindBody(obj interface{}) error {
	return c.ctx.Bind().Body(obj)
}

// BindQuery binds the query to the given struct.
func (c *UnimplementedController) BindQuery(obj interface{}) error {
	return c.ctx.Bind().Query(obj)
}

// BindAll binds the body and query to the given struct.
func (c *UnimplementedController) BindAll(obj interface{}) error {
	return c.ctx.Bind().All(obj)
}

// Values is a helper function to get the values from the context.
func (c *UnimplementedController) Values(key any) (val any) {
	return c.ctx.Value(key)
}

// ValuesString is a helper function to get the values as a string from the context.
func (c *UnimplementedController) ValuesString(key any) (val string) {
	return conv.String(c.ctx.Value(key))
}

// ValuesInt is a helper function to get the values as an int from the context.
func (c *UnimplementedController) ValuesInt(key any) (val int) {
	return conv.Int(c.ctx.Value(key))
}

// ValuesBool is a helper function to get the values as a bool from the context.
func (c *UnimplementedController) ValuesBool(key any) (val bool) {
	return conv.Bool(c.ctx.Value(key))
}

// Session is a helper function to get the session from the context.
func (c *UnimplementedController) Session() adapters.GothSession {
	session, err := goth.SessionFromContext(c.ctx)
	if err != nil {
		return adapters.GothSession{}
	}

	return session
}

// Messages is a helper function to get the message from the context.
func (c *UnimplementedController) Messages(msgs ...Message) {
	m := MessagesFromContext(c.ctx)
	m.Add(msgs...)
}

// Context returns the context.
func (c *UnimplementedController) Context() context.Context {
	return c.ctx
}

// Ctx returns the fiber.Ctx.
func (c *UnimplementedController) Ctx() fiber.Ctx {
	return c.ctx
}

// Render is a helper function to render a component.
func (c *UnimplementedController) Render(node htmx.Node, opt ...RenderOpt) error {
	return RenderComp(c.ctx, node, opt...)
}

// Clone returns a new instance of the controller.
func (c *UnimplementedController) Clone() Controller {
	copy := *c
	copy.Reset()

	return &copy
}

// IsDevelopment returns true if the environment is development.
func (c *UnimplementedController) IsDevelopment() bool {
	return reload.IsDevelopment(c.ctx.Context())
}

// IsProduction returns true if the environment is production.
func (c *UnimplementedController) IsProduction() bool {
	return reload.IsProduction(c.ctx.Context())
}

// Redirect redirects to the given path.
func (c *UnimplementedController) Redirect(path string) {
	Redirect(c.ctx, path)
}

// Path returns the path.
func (c *UnimplementedController) Path() string {
	return c.ctx.Path()
}

// OriginalURL returns the original URL.
func (c *UnimplementedController) OriginalURL() string {
	return c.ctx.OriginalURL()
}

// Reset resets the controller.
func (c *UnimplementedController) Reset() {
	c.ctx = nil
}

var _ TransactionController = (*UnimplementedTransactionController)(nil)

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

// UnimplementedTransactionController is the default implementation of a TransactionController.
type UnimplementedTransactionController struct {
	TransactionController
	tx *gorm.DB
}

// Begin begins a transaction.
func (c *UnimplementedTransactionController) Begin(conn *gorm.DB) error {
	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	c.tx = tx

	return nil
}

// Commit commits the transaction.
func (c *UnimplementedTransactionController) Commit() error {
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
func (c *UnimplementedTransactionController) Rollback() error {
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
func (c *UnimplementedTransactionController) Tx() *gorm.DB {
	if c.tx != nil {
		return c.tx
	}

	return nil
}
