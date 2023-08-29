package tests

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"sms-sender/ent"
	"sms-sender/ent/enttest"
	"testing"
)

func SetupDatabase(t *testing.T) (client *ent.Client) {
	client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}

	return client
}
