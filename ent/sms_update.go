// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sms-sender/ent/predicate"
	"sms-sender/ent/sms"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SmsUpdate is the builder for updating Sms entities.
type SmsUpdate struct {
	config
	hooks    []Hook
	mutation *SmsMutation
}

// Where appends a list predicates to the SmsUpdate builder.
func (su *SmsUpdate) Where(ps ...predicate.Sms) *SmsUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetRecipient sets the "recipient" field.
func (su *SmsUpdate) SetRecipient(s string) *SmsUpdate {
	su.mutation.SetRecipient(s)
	return su
}

// SetMessage sets the "message" field.
func (su *SmsUpdate) SetMessage(s string) *SmsUpdate {
	su.mutation.SetMessage(s)
	return su
}

// SetCommercial sets the "commercial" field.
func (su *SmsUpdate) SetCommercial(b bool) *SmsUpdate {
	su.mutation.SetCommercial(b)
	return su
}

// SetSentAt sets the "sentAt" field.
func (su *SmsUpdate) SetSentAt(t time.Time) *SmsUpdate {
	su.mutation.SetSentAt(t)
	return su
}

// SetNillableSentAt sets the "sentAt" field if the given value is not nil.
func (su *SmsUpdate) SetNillableSentAt(t *time.Time) *SmsUpdate {
	if t != nil {
		su.SetSentAt(*t)
	}
	return su
}

// SetStatus sets the "status" field.
func (su *SmsUpdate) SetStatus(s string) *SmsUpdate {
	su.mutation.SetStatus(s)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *SmsUpdate) SetNillableStatus(s *string) *SmsUpdate {
	if s != nil {
		su.SetStatus(*s)
	}
	return su
}

// ClearStatus clears the value of the "status" field.
func (su *SmsUpdate) ClearStatus() *SmsUpdate {
	su.mutation.ClearStatus()
	return su
}

// Mutation returns the SmsMutation object of the builder.
func (su *SmsUpdate) Mutation() *SmsMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SmsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SmsUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SmsUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SmsUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SmsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sms.Table, sms.Columns, sqlgraph.NewFieldSpec(sms.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Recipient(); ok {
		_spec.SetField(sms.FieldRecipient, field.TypeString, value)
	}
	if value, ok := su.mutation.Message(); ok {
		_spec.SetField(sms.FieldMessage, field.TypeString, value)
	}
	if value, ok := su.mutation.Commercial(); ok {
		_spec.SetField(sms.FieldCommercial, field.TypeBool, value)
	}
	if value, ok := su.mutation.SentAt(); ok {
		_spec.SetField(sms.FieldSentAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(sms.FieldStatus, field.TypeString, value)
	}
	if su.mutation.StatusCleared() {
		_spec.ClearField(sms.FieldStatus, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sms.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SmsUpdateOne is the builder for updating a single Sms entity.
type SmsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SmsMutation
}

// SetRecipient sets the "recipient" field.
func (suo *SmsUpdateOne) SetRecipient(s string) *SmsUpdateOne {
	suo.mutation.SetRecipient(s)
	return suo
}

// SetMessage sets the "message" field.
func (suo *SmsUpdateOne) SetMessage(s string) *SmsUpdateOne {
	suo.mutation.SetMessage(s)
	return suo
}

// SetCommercial sets the "commercial" field.
func (suo *SmsUpdateOne) SetCommercial(b bool) *SmsUpdateOne {
	suo.mutation.SetCommercial(b)
	return suo
}

// SetSentAt sets the "sentAt" field.
func (suo *SmsUpdateOne) SetSentAt(t time.Time) *SmsUpdateOne {
	suo.mutation.SetSentAt(t)
	return suo
}

// SetNillableSentAt sets the "sentAt" field if the given value is not nil.
func (suo *SmsUpdateOne) SetNillableSentAt(t *time.Time) *SmsUpdateOne {
	if t != nil {
		suo.SetSentAt(*t)
	}
	return suo
}

// SetStatus sets the "status" field.
func (suo *SmsUpdateOne) SetStatus(s string) *SmsUpdateOne {
	suo.mutation.SetStatus(s)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *SmsUpdateOne) SetNillableStatus(s *string) *SmsUpdateOne {
	if s != nil {
		suo.SetStatus(*s)
	}
	return suo
}

// ClearStatus clears the value of the "status" field.
func (suo *SmsUpdateOne) ClearStatus() *SmsUpdateOne {
	suo.mutation.ClearStatus()
	return suo
}

// Mutation returns the SmsMutation object of the builder.
func (suo *SmsUpdateOne) Mutation() *SmsMutation {
	return suo.mutation
}

// Where appends a list predicates to the SmsUpdate builder.
func (suo *SmsUpdateOne) Where(ps ...predicate.Sms) *SmsUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SmsUpdateOne) Select(field string, fields ...string) *SmsUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Sms entity.
func (suo *SmsUpdateOne) Save(ctx context.Context) (*Sms, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SmsUpdateOne) SaveX(ctx context.Context) *Sms {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SmsUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SmsUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SmsUpdateOne) sqlSave(ctx context.Context) (_node *Sms, err error) {
	_spec := sqlgraph.NewUpdateSpec(sms.Table, sms.Columns, sqlgraph.NewFieldSpec(sms.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Sms.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sms.FieldID)
		for _, f := range fields {
			if !sms.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sms.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Recipient(); ok {
		_spec.SetField(sms.FieldRecipient, field.TypeString, value)
	}
	if value, ok := suo.mutation.Message(); ok {
		_spec.SetField(sms.FieldMessage, field.TypeString, value)
	}
	if value, ok := suo.mutation.Commercial(); ok {
		_spec.SetField(sms.FieldCommercial, field.TypeBool, value)
	}
	if value, ok := suo.mutation.SentAt(); ok {
		_spec.SetField(sms.FieldSentAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(sms.FieldStatus, field.TypeString, value)
	}
	if suo.mutation.StatusCleared() {
		_spec.ClearField(sms.FieldStatus, field.TypeString)
	}
	_node = &Sms{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sms.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}