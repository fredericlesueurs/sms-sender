// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SmsColumns holds the columns for the "sms" table.
	SmsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "recipient", Type: field.TypeString},
		{Name: "message", Type: field.TypeString},
		{Name: "commercial", Type: field.TypeBool},
		{Name: "sent_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeString, Nullable: true},
	}
	// SmsTable holds the schema information for the "sms" table.
	SmsTable = &schema.Table{
		Name:       "sms",
		Columns:    SmsColumns,
		PrimaryKey: []*schema.Column{SmsColumns[0]},
	}
	// StopRequestsColumns holds the columns for the "stop_requests" table.
	StopRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "recipient", Type: field.TypeString},
		{Name: "stop_date", Type: field.TypeTime},
	}
	// StopRequestsTable holds the schema information for the "stop_requests" table.
	StopRequestsTable = &schema.Table{
		Name:       "stop_requests",
		Columns:    StopRequestsColumns,
		PrimaryKey: []*schema.Column{StopRequestsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SmsTable,
		StopRequestsTable,
	}
)

func init() {
}
