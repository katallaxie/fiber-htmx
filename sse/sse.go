package sse

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/katallaxie/pkg/server"
	"github.com/valyala/fasthttp"
)

var _ Event = (*EventImpl)(nil)

// Event defines the interface for content that can be broadcast to clients.
type Event interface {
	String() string // Represent the envelope contents as a string for transmission.
}

// Client is the interface for sending server-sent events.
type Client interface {
	ID() string
	Events() chan Event
	Close()
}

// Manager is the interface for broadcasting messages to clients.
type Manager interface {
	Add(Client)
	Remove(Client)
	Send() chan<- Event
}

// EventImpl is the default implementation of the Event interface.
type EventImpl struct {
	Event string
	Time  time.Time
	Data  string
}

// NewEvent creates a new event.
func NewMessage(event, data string) *EventImpl {
	return &EventImpl{
		Data:  data,
		Event: event,
		Time:  time.Now(),
	}
}

// String returns the message as a string.
func (e *EventImpl) String() string {
	sb := strings.Builder{}

	if e.Event != "" {
		sb.WriteString(fmt.Sprintf("event: %s\n", e.Event))
	}
	sb.WriteString(fmt.Sprintf("data: %v\n\n", e.Data))

	return sb.String()
}

var _ Manager = (*BroadcastManagerImpl)(nil)

// BroadcastManagerImpl is the default implementation of the BroadcastManager interface.
type BroadcastManagerImpl struct {
	id        string
	broadcast chan Event
	poolSize  int
	clients   sync.Map
}

var _ server.Listener = (*BroadcastManagerImpl)(nil)

// NewBroadcastManager creates a new broadcast manager.
func NewBroadcastManager(poolSize int) *BroadcastManagerImpl {
	return &BroadcastManagerImpl{
		id:        uuid.NewString(),
		broadcast: make(chan Event),
		poolSize:  poolSize,
	}
}

// Add adds a client to the broadcast manager.
func (b *BroadcastManagerImpl) Add(client Client) {
	b.clients.Store(client.ID(), client)
}

// Remove removes a client from the broadcast manager.
func (b *BroadcastManagerImpl) Remove(client Client) {
	b.clients.Delete(client.ID())
}

// Start starts the broadcast manager.
func (b *BroadcastManagerImpl) Start(ctx context.Context, ready server.ReadyFunc, _ server.RunFunc) func() error {
	return func() error {
		if b.poolSize < 1 {
			return server.NewServerError(fmt.Errorf("pool size must be greater than 0"))
		}

		b.startWorkers(ctx)

		ready()

		<-ctx.Done()

		return nil
	}
}

// Client returns the client.
type ClientImpl struct {
	events chan Event
	id     string
}

// ID returns the client ID.
func (c *ClientImpl) ID() string {
	return c.id
}

// Events returns the client events.
func (c *ClientImpl) Events() chan Event {
	return c.events
}

// Close closes the client.
func (c *ClientImpl) Close() {
	<-c.events // Drain the channel.
}

// NewClient creates a new client.
func NewClient(id string) *ClientImpl {
	if id == "" {
		id = uuid.New().String()
	}

	return &ClientImpl{
		events: make(chan Event),
		id:     id,
	}
}

func (b *BroadcastManagerImpl) startWorkers(ctx context.Context) {
	for i := 0; i < b.poolSize; i++ {
		go func() {
			for {
				select {
				case msg := <-b.broadcast:
					b.clients.Range(func(_, value any) bool {
						client, ok := value.(Client)
						if !ok {
							return true
						}

						select {
						case client.Events() <- msg:
						default: // client not reachable
						}

						return true
					})
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}

// Send sends a message to all clients.
func (b *BroadcastManagerImpl) Send() chan<- Event {
	return b.broadcast
}

// Config is the configuration for the server-sent events server.
type Config struct{}

// NewSSEHandler creates a new server-sent events handler.
func NewSSEHandler(manager Manager, config ...Config) fiber.Handler {
	_ = configDefault(config...)

	return func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "text/event-stream")
		c.Set(fiber.HeaderCacheControl, "no-cache")
		c.Set(fiber.HeaderConnection, "keep-alive")
		c.Set(fiber.HeaderTransferEncoding, "chunked")

		client := NewClient(uuid.NewString())
		manager.Add(client)

		notify := c.Context().Done()

		c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			for {
				select {
				case <-notify:
					manager.Remove(client)
					return
				case msg := <-client.Events():
					_, err := fmt.Fprint(w, msg.String())
					if err != nil {
						return
					}

					err = w.Flush()
					if err != nil {
						// Refreshing page in web browser will establish a new
						// SSE connection, but only (the last) one is alive, so
						// dead connections must be closed here.
						return
					}
				}
			}
		}))

		return nil
	}
}

// ConfigDefault is the default configuration for the server-sent events server.
var ConfigDefault = Config{}

// Helper function to set default values.
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	return cfg
}
