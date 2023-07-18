package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Sms holds the schema definition for the Sms entity.
type Sms struct {
	ent.Schema
}

// Fields of the Sms.
func (Sms) Fields() []ent.Field {
	return []ent.Field{
		field.String("recipient"),
		field.String("message"),
		field.Bool("commercial"),
		field.Time("sentAt").Default(time.Now()),
		field.String("status").Optional().Nillable(),
	}
}

// Edges of the Sms.
func (Sms) Edges() []ent.Edge {
	return nil
}
