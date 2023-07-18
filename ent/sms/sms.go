// Code generated by ent, DO NOT EDIT.

package sms

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the sms type in the database.
	Label = "sms"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRecipient holds the string denoting the recipient field in the database.
	FieldRecipient = "recipient"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldCommercial holds the string denoting the commercial field in the database.
	FieldCommercial = "commercial"
	// FieldSentAt holds the string denoting the sentat field in the database.
	FieldSentAt = "sent_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the sms in the database.
	Table = "sms"
)

// Columns holds all SQL columns for sms fields.
var Columns = []string{
	FieldID,
	FieldRecipient,
	FieldMessage,
	FieldCommercial,
	FieldSentAt,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSentAt holds the default value on creation for the "sentAt" field.
	DefaultSentAt time.Time
)

// OrderOption defines the ordering options for the Sms queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRecipient orders the results by the recipient field.
func ByRecipient(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecipient, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByCommercial orders the results by the commercial field.
func ByCommercial(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommercial, opts...).ToFunc()
}

// BySentAt orders the results by the sentAt field.
func BySentAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSentAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}