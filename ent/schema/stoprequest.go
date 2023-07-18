package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// StopRequest holds the schema definition for the StopRequest entity.
type StopRequest struct {
	ent.Schema
}

// Fields of the StopRequest.
func (StopRequest) Fields() []ent.Field {
	return []ent.Field{
		field.String("recipient"),
		field.Time("stop_date").Default(time.Now()),
	}
}

// Edges of the StopRequest.
func (StopRequest) Edges() []ent.Edge {
	return nil
}
