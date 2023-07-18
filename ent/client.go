// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"sms-sender/ent/migrate"

	"sms-sender/ent/sms"
	"sms-sender/ent/stoprequest"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Sms is the client for interacting with the Sms builders.
	Sms *SmsClient
	// StopRequest is the client for interacting with the StopRequest builders.
	StopRequest *StopRequestClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Sms = NewSmsClient(c.config)
	c.StopRequest = NewStopRequestClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Sms:         NewSmsClient(cfg),
		StopRequest: NewStopRequestClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Sms:         NewSmsClient(cfg),
		StopRequest: NewStopRequestClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Sms.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Sms.Use(hooks...)
	c.StopRequest.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Sms.Intercept(interceptors...)
	c.StopRequest.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *SmsMutation:
		return c.Sms.mutate(ctx, m)
	case *StopRequestMutation:
		return c.StopRequest.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// SmsClient is a client for the Sms schema.
type SmsClient struct {
	config
}

// NewSmsClient returns a client for the Sms from the given config.
func NewSmsClient(c config) *SmsClient {
	return &SmsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sms.Hooks(f(g(h())))`.
func (c *SmsClient) Use(hooks ...Hook) {
	c.hooks.Sms = append(c.hooks.Sms, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `sms.Intercept(f(g(h())))`.
func (c *SmsClient) Intercept(interceptors ...Interceptor) {
	c.inters.Sms = append(c.inters.Sms, interceptors...)
}

// Create returns a builder for creating a Sms entity.
func (c *SmsClient) Create() *SmsCreate {
	mutation := newSmsMutation(c.config, OpCreate)
	return &SmsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Sms entities.
func (c *SmsClient) CreateBulk(builders ...*SmsCreate) *SmsCreateBulk {
	return &SmsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Sms.
func (c *SmsClient) Update() *SmsUpdate {
	mutation := newSmsMutation(c.config, OpUpdate)
	return &SmsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SmsClient) UpdateOne(s *Sms) *SmsUpdateOne {
	mutation := newSmsMutation(c.config, OpUpdateOne, withSms(s))
	return &SmsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SmsClient) UpdateOneID(id int) *SmsUpdateOne {
	mutation := newSmsMutation(c.config, OpUpdateOne, withSmsID(id))
	return &SmsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Sms.
func (c *SmsClient) Delete() *SmsDelete {
	mutation := newSmsMutation(c.config, OpDelete)
	return &SmsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SmsClient) DeleteOne(s *Sms) *SmsDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SmsClient) DeleteOneID(id int) *SmsDeleteOne {
	builder := c.Delete().Where(sms.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SmsDeleteOne{builder}
}

// Query returns a query builder for Sms.
func (c *SmsClient) Query() *SmsQuery {
	return &SmsQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSms},
		inters: c.Interceptors(),
	}
}

// Get returns a Sms entity by its id.
func (c *SmsClient) Get(ctx context.Context, id int) (*Sms, error) {
	return c.Query().Where(sms.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SmsClient) GetX(ctx context.Context, id int) *Sms {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SmsClient) Hooks() []Hook {
	return c.hooks.Sms
}

// Interceptors returns the client interceptors.
func (c *SmsClient) Interceptors() []Interceptor {
	return c.inters.Sms
}

func (c *SmsClient) mutate(ctx context.Context, m *SmsMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SmsCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SmsUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SmsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SmsDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Sms mutation op: %q", m.Op())
	}
}

// StopRequestClient is a client for the StopRequest schema.
type StopRequestClient struct {
	config
}

// NewStopRequestClient returns a client for the StopRequest from the given config.
func NewStopRequestClient(c config) *StopRequestClient {
	return &StopRequestClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `stoprequest.Hooks(f(g(h())))`.
func (c *StopRequestClient) Use(hooks ...Hook) {
	c.hooks.StopRequest = append(c.hooks.StopRequest, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `stoprequest.Intercept(f(g(h())))`.
func (c *StopRequestClient) Intercept(interceptors ...Interceptor) {
	c.inters.StopRequest = append(c.inters.StopRequest, interceptors...)
}

// Create returns a builder for creating a StopRequest entity.
func (c *StopRequestClient) Create() *StopRequestCreate {
	mutation := newStopRequestMutation(c.config, OpCreate)
	return &StopRequestCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of StopRequest entities.
func (c *StopRequestClient) CreateBulk(builders ...*StopRequestCreate) *StopRequestCreateBulk {
	return &StopRequestCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for StopRequest.
func (c *StopRequestClient) Update() *StopRequestUpdate {
	mutation := newStopRequestMutation(c.config, OpUpdate)
	return &StopRequestUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StopRequestClient) UpdateOne(sr *StopRequest) *StopRequestUpdateOne {
	mutation := newStopRequestMutation(c.config, OpUpdateOne, withStopRequest(sr))
	return &StopRequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StopRequestClient) UpdateOneID(id int) *StopRequestUpdateOne {
	mutation := newStopRequestMutation(c.config, OpUpdateOne, withStopRequestID(id))
	return &StopRequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for StopRequest.
func (c *StopRequestClient) Delete() *StopRequestDelete {
	mutation := newStopRequestMutation(c.config, OpDelete)
	return &StopRequestDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StopRequestClient) DeleteOne(sr *StopRequest) *StopRequestDeleteOne {
	return c.DeleteOneID(sr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StopRequestClient) DeleteOneID(id int) *StopRequestDeleteOne {
	builder := c.Delete().Where(stoprequest.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StopRequestDeleteOne{builder}
}

// Query returns a query builder for StopRequest.
func (c *StopRequestClient) Query() *StopRequestQuery {
	return &StopRequestQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStopRequest},
		inters: c.Interceptors(),
	}
}

// Get returns a StopRequest entity by its id.
func (c *StopRequestClient) Get(ctx context.Context, id int) (*StopRequest, error) {
	return c.Query().Where(stoprequest.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StopRequestClient) GetX(ctx context.Context, id int) *StopRequest {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *StopRequestClient) Hooks() []Hook {
	return c.hooks.StopRequest
}

// Interceptors returns the client interceptors.
func (c *StopRequestClient) Interceptors() []Interceptor {
	return c.inters.StopRequest
}

func (c *StopRequestClient) mutate(ctx context.Context, m *StopRequestMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StopRequestCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StopRequestUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StopRequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StopRequestDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown StopRequest mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Sms, StopRequest []ent.Hook
	}
	inters struct {
		Sms, StopRequest []ent.Interceptor
	}
)
