// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sms-sender/ent/predicate"
	"sms-sender/ent/sms"
	"sms-sender/ent/stoprequest"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeSms         = "Sms"
	TypeStopRequest = "StopRequest"
)

// SmsMutation represents an operation that mutates the Sms nodes in the graph.
type SmsMutation struct {
	config
	op            Op
	typ           string
	id            *int
	recipient     *string
	message       *string
	commercial    *bool
	sentAt        *time.Time
	status        *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Sms, error)
	predicates    []predicate.Sms
}

var _ ent.Mutation = (*SmsMutation)(nil)

// smsOption allows management of the mutation configuration using functional options.
type smsOption func(*SmsMutation)

// newSmsMutation creates new mutation for the Sms entity.
func newSmsMutation(c config, op Op, opts ...smsOption) *SmsMutation {
	m := &SmsMutation{
		config:        c,
		op:            op,
		typ:           TypeSms,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSmsID sets the ID field of the mutation.
func withSmsID(id int) smsOption {
	return func(m *SmsMutation) {
		var (
			err   error
			once  sync.Once
			value *Sms
		)
		m.oldValue = func(ctx context.Context) (*Sms, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Sms.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSms sets the old Sms of the mutation.
func withSms(node *Sms) smsOption {
	return func(m *SmsMutation) {
		m.oldValue = func(context.Context) (*Sms, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SmsMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SmsMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SmsMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SmsMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Sms.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetRecipient sets the "recipient" field.
func (m *SmsMutation) SetRecipient(s string) {
	m.recipient = &s
}

// Recipient returns the value of the "recipient" field in the mutation.
func (m *SmsMutation) Recipient() (r string, exists bool) {
	v := m.recipient
	if v == nil {
		return
	}
	return *v, true
}

// OldRecipient returns the old "recipient" field's value of the Sms entity.
// If the Sms object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SmsMutation) OldRecipient(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRecipient is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRecipient requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRecipient: %w", err)
	}
	return oldValue.Recipient, nil
}

// ResetRecipient resets all changes to the "recipient" field.
func (m *SmsMutation) ResetRecipient() {
	m.recipient = nil
}

// SetMessage sets the "message" field.
func (m *SmsMutation) SetMessage(s string) {
	m.message = &s
}

// Message returns the value of the "message" field in the mutation.
func (m *SmsMutation) Message() (r string, exists bool) {
	v := m.message
	if v == nil {
		return
	}
	return *v, true
}

// OldMessage returns the old "message" field's value of the Sms entity.
// If the Sms object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SmsMutation) OldMessage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMessage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMessage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMessage: %w", err)
	}
	return oldValue.Message, nil
}

// ResetMessage resets all changes to the "message" field.
func (m *SmsMutation) ResetMessage() {
	m.message = nil
}

// SetCommercial sets the "commercial" field.
func (m *SmsMutation) SetCommercial(b bool) {
	m.commercial = &b
}

// Commercial returns the value of the "commercial" field in the mutation.
func (m *SmsMutation) Commercial() (r bool, exists bool) {
	v := m.commercial
	if v == nil {
		return
	}
	return *v, true
}

// OldCommercial returns the old "commercial" field's value of the Sms entity.
// If the Sms object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SmsMutation) OldCommercial(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCommercial is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCommercial requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCommercial: %w", err)
	}
	return oldValue.Commercial, nil
}

// ResetCommercial resets all changes to the "commercial" field.
func (m *SmsMutation) ResetCommercial() {
	m.commercial = nil
}

// SetSentAt sets the "sentAt" field.
func (m *SmsMutation) SetSentAt(t time.Time) {
	m.sentAt = &t
}

// SentAt returns the value of the "sentAt" field in the mutation.
func (m *SmsMutation) SentAt() (r time.Time, exists bool) {
	v := m.sentAt
	if v == nil {
		return
	}
	return *v, true
}

// OldSentAt returns the old "sentAt" field's value of the Sms entity.
// If the Sms object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SmsMutation) OldSentAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSentAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSentAt: %w", err)
	}
	return oldValue.SentAt, nil
}

// ResetSentAt resets all changes to the "sentAt" field.
func (m *SmsMutation) ResetSentAt() {
	m.sentAt = nil
}

// SetStatus sets the "status" field.
func (m *SmsMutation) SetStatus(s string) {
	m.status = &s
}

// Status returns the value of the "status" field in the mutation.
func (m *SmsMutation) Status() (r string, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Sms entity.
// If the Sms object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SmsMutation) OldStatus(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ClearStatus clears the value of the "status" field.
func (m *SmsMutation) ClearStatus() {
	m.status = nil
	m.clearedFields[sms.FieldStatus] = struct{}{}
}

// StatusCleared returns if the "status" field was cleared in this mutation.
func (m *SmsMutation) StatusCleared() bool {
	_, ok := m.clearedFields[sms.FieldStatus]
	return ok
}

// ResetStatus resets all changes to the "status" field.
func (m *SmsMutation) ResetStatus() {
	m.status = nil
	delete(m.clearedFields, sms.FieldStatus)
}

// Where appends a list predicates to the SmsMutation builder.
func (m *SmsMutation) Where(ps ...predicate.Sms) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the SmsMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *SmsMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Sms, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *SmsMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *SmsMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Sms).
func (m *SmsMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SmsMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.recipient != nil {
		fields = append(fields, sms.FieldRecipient)
	}
	if m.message != nil {
		fields = append(fields, sms.FieldMessage)
	}
	if m.commercial != nil {
		fields = append(fields, sms.FieldCommercial)
	}
	if m.sentAt != nil {
		fields = append(fields, sms.FieldSentAt)
	}
	if m.status != nil {
		fields = append(fields, sms.FieldStatus)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SmsMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case sms.FieldRecipient:
		return m.Recipient()
	case sms.FieldMessage:
		return m.Message()
	case sms.FieldCommercial:
		return m.Commercial()
	case sms.FieldSentAt:
		return m.SentAt()
	case sms.FieldStatus:
		return m.Status()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SmsMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case sms.FieldRecipient:
		return m.OldRecipient(ctx)
	case sms.FieldMessage:
		return m.OldMessage(ctx)
	case sms.FieldCommercial:
		return m.OldCommercial(ctx)
	case sms.FieldSentAt:
		return m.OldSentAt(ctx)
	case sms.FieldStatus:
		return m.OldStatus(ctx)
	}
	return nil, fmt.Errorf("unknown Sms field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SmsMutation) SetField(name string, value ent.Value) error {
	switch name {
	case sms.FieldRecipient:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRecipient(v)
		return nil
	case sms.FieldMessage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMessage(v)
		return nil
	case sms.FieldCommercial:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCommercial(v)
		return nil
	case sms.FieldSentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSentAt(v)
		return nil
	case sms.FieldStatus:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	}
	return fmt.Errorf("unknown Sms field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SmsMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SmsMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SmsMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Sms numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SmsMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(sms.FieldStatus) {
		fields = append(fields, sms.FieldStatus)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SmsMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SmsMutation) ClearField(name string) error {
	switch name {
	case sms.FieldStatus:
		m.ClearStatus()
		return nil
	}
	return fmt.Errorf("unknown Sms nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SmsMutation) ResetField(name string) error {
	switch name {
	case sms.FieldRecipient:
		m.ResetRecipient()
		return nil
	case sms.FieldMessage:
		m.ResetMessage()
		return nil
	case sms.FieldCommercial:
		m.ResetCommercial()
		return nil
	case sms.FieldSentAt:
		m.ResetSentAt()
		return nil
	case sms.FieldStatus:
		m.ResetStatus()
		return nil
	}
	return fmt.Errorf("unknown Sms field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SmsMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SmsMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SmsMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SmsMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SmsMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SmsMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SmsMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Sms unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SmsMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Sms edge %s", name)
}

// StopRequestMutation represents an operation that mutates the StopRequest nodes in the graph.
type StopRequestMutation struct {
	config
	op            Op
	typ           string
	id            *int
	recipient     *string
	stop_date     *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*StopRequest, error)
	predicates    []predicate.StopRequest
}

var _ ent.Mutation = (*StopRequestMutation)(nil)

// stoprequestOption allows management of the mutation configuration using functional options.
type stoprequestOption func(*StopRequestMutation)

// newStopRequestMutation creates new mutation for the StopRequest entity.
func newStopRequestMutation(c config, op Op, opts ...stoprequestOption) *StopRequestMutation {
	m := &StopRequestMutation{
		config:        c,
		op:            op,
		typ:           TypeStopRequest,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withStopRequestID sets the ID field of the mutation.
func withStopRequestID(id int) stoprequestOption {
	return func(m *StopRequestMutation) {
		var (
			err   error
			once  sync.Once
			value *StopRequest
		)
		m.oldValue = func(ctx context.Context) (*StopRequest, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().StopRequest.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withStopRequest sets the old StopRequest of the mutation.
func withStopRequest(node *StopRequest) stoprequestOption {
	return func(m *StopRequestMutation) {
		m.oldValue = func(context.Context) (*StopRequest, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m StopRequestMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m StopRequestMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *StopRequestMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *StopRequestMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().StopRequest.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetRecipient sets the "recipient" field.
func (m *StopRequestMutation) SetRecipient(s string) {
	m.recipient = &s
}

// Recipient returns the value of the "recipient" field in the mutation.
func (m *StopRequestMutation) Recipient() (r string, exists bool) {
	v := m.recipient
	if v == nil {
		return
	}
	return *v, true
}

// OldRecipient returns the old "recipient" field's value of the StopRequest entity.
// If the StopRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StopRequestMutation) OldRecipient(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRecipient is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRecipient requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRecipient: %w", err)
	}
	return oldValue.Recipient, nil
}

// ResetRecipient resets all changes to the "recipient" field.
func (m *StopRequestMutation) ResetRecipient() {
	m.recipient = nil
}

// SetStopDate sets the "stop_date" field.
func (m *StopRequestMutation) SetStopDate(t time.Time) {
	m.stop_date = &t
}

// StopDate returns the value of the "stop_date" field in the mutation.
func (m *StopRequestMutation) StopDate() (r time.Time, exists bool) {
	v := m.stop_date
	if v == nil {
		return
	}
	return *v, true
}

// OldStopDate returns the old "stop_date" field's value of the StopRequest entity.
// If the StopRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StopRequestMutation) OldStopDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStopDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStopDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStopDate: %w", err)
	}
	return oldValue.StopDate, nil
}

// ResetStopDate resets all changes to the "stop_date" field.
func (m *StopRequestMutation) ResetStopDate() {
	m.stop_date = nil
}

// Where appends a list predicates to the StopRequestMutation builder.
func (m *StopRequestMutation) Where(ps ...predicate.StopRequest) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the StopRequestMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *StopRequestMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.StopRequest, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *StopRequestMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *StopRequestMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (StopRequest).
func (m *StopRequestMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *StopRequestMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.recipient != nil {
		fields = append(fields, stoprequest.FieldRecipient)
	}
	if m.stop_date != nil {
		fields = append(fields, stoprequest.FieldStopDate)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *StopRequestMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case stoprequest.FieldRecipient:
		return m.Recipient()
	case stoprequest.FieldStopDate:
		return m.StopDate()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *StopRequestMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case stoprequest.FieldRecipient:
		return m.OldRecipient(ctx)
	case stoprequest.FieldStopDate:
		return m.OldStopDate(ctx)
	}
	return nil, fmt.Errorf("unknown StopRequest field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StopRequestMutation) SetField(name string, value ent.Value) error {
	switch name {
	case stoprequest.FieldRecipient:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRecipient(v)
		return nil
	case stoprequest.FieldStopDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStopDate(v)
		return nil
	}
	return fmt.Errorf("unknown StopRequest field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *StopRequestMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *StopRequestMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StopRequestMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown StopRequest numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *StopRequestMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *StopRequestMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *StopRequestMutation) ClearField(name string) error {
	return fmt.Errorf("unknown StopRequest nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *StopRequestMutation) ResetField(name string) error {
	switch name {
	case stoprequest.FieldRecipient:
		m.ResetRecipient()
		return nil
	case stoprequest.FieldStopDate:
		m.ResetStopDate()
		return nil
	}
	return fmt.Errorf("unknown StopRequest field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *StopRequestMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *StopRequestMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *StopRequestMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *StopRequestMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *StopRequestMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *StopRequestMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *StopRequestMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown StopRequest unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *StopRequestMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown StopRequest edge %s", name)
}
