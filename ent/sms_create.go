// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sms-sender/ent/sms"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SmsCreate is the builder for creating a Sms entity.
type SmsCreate struct {
	config
	mutation *SmsMutation
	hooks    []Hook
}

// SetRecipient sets the "recipient" field.
func (sc *SmsCreate) SetRecipient(s string) *SmsCreate {
	sc.mutation.SetRecipient(s)
	return sc
}

// SetMessage sets the "message" field.
func (sc *SmsCreate) SetMessage(s string) *SmsCreate {
	sc.mutation.SetMessage(s)
	return sc
}

// SetCommercial sets the "commercial" field.
func (sc *SmsCreate) SetCommercial(b bool) *SmsCreate {
	sc.mutation.SetCommercial(b)
	return sc
}

// SetSentAt sets the "sentAt" field.
func (sc *SmsCreate) SetSentAt(t time.Time) *SmsCreate {
	sc.mutation.SetSentAt(t)
	return sc
}

// SetNillableSentAt sets the "sentAt" field if the given value is not nil.
func (sc *SmsCreate) SetNillableSentAt(t *time.Time) *SmsCreate {
	if t != nil {
		sc.SetSentAt(*t)
	}
	return sc
}

// SetStatus sets the "status" field.
func (sc *SmsCreate) SetStatus(s string) *SmsCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *SmsCreate) SetNillableStatus(s *string) *SmsCreate {
	if s != nil {
		sc.SetStatus(*s)
	}
	return sc
}

// Mutation returns the SmsMutation object of the builder.
func (sc *SmsCreate) Mutation() *SmsMutation {
	return sc.mutation
}

// Save creates the Sms in the database.
func (sc *SmsCreate) Save(ctx context.Context) (*Sms, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SmsCreate) SaveX(ctx context.Context) *Sms {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SmsCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SmsCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SmsCreate) defaults() {
	if _, ok := sc.mutation.SentAt(); !ok {
		v := sms.DefaultSentAt
		sc.mutation.SetSentAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SmsCreate) check() error {
	if _, ok := sc.mutation.Recipient(); !ok {
		return &ValidationError{Name: "recipient", err: errors.New(`ent: missing required field "Sms.recipient"`)}
	}
	if _, ok := sc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "Sms.message"`)}
	}
	if _, ok := sc.mutation.Commercial(); !ok {
		return &ValidationError{Name: "commercial", err: errors.New(`ent: missing required field "Sms.commercial"`)}
	}
	if _, ok := sc.mutation.SentAt(); !ok {
		return &ValidationError{Name: "sentAt", err: errors.New(`ent: missing required field "Sms.sentAt"`)}
	}
	return nil
}

func (sc *SmsCreate) sqlSave(ctx context.Context) (*Sms, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SmsCreate) createSpec() (*Sms, *sqlgraph.CreateSpec) {
	var (
		_node = &Sms{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(sms.Table, sqlgraph.NewFieldSpec(sms.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Recipient(); ok {
		_spec.SetField(sms.FieldRecipient, field.TypeString, value)
		_node.Recipient = value
	}
	if value, ok := sc.mutation.Message(); ok {
		_spec.SetField(sms.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if value, ok := sc.mutation.Commercial(); ok {
		_spec.SetField(sms.FieldCommercial, field.TypeBool, value)
		_node.Commercial = value
	}
	if value, ok := sc.mutation.SentAt(); ok {
		_spec.SetField(sms.FieldSentAt, field.TypeTime, value)
		_node.SentAt = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(sms.FieldStatus, field.TypeString, value)
		_node.Status = &value
	}
	return _node, _spec
}

// SmsCreateBulk is the builder for creating many Sms entities in bulk.
type SmsCreateBulk struct {
	config
	builders []*SmsCreate
}

// Save creates the Sms entities in the database.
func (scb *SmsCreateBulk) Save(ctx context.Context) ([]*Sms, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sms, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SmsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SmsCreateBulk) SaveX(ctx context.Context) []*Sms {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SmsCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SmsCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
