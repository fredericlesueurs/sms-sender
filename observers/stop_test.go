package observers

import (
	"context"
	"github.com/warthog618/modem/gsm"
	"sms-sender/ent/stoprequest"
	"sms-sender/tests"
	"testing"
)

func TestStopObserver_Stop(t *testing.T) {
	ctx := context.Background()

	_client := tests.SetupDatabase(t)
	observer := CreateStopObserver(_client)

	message := &gsm.Message{
		Message: "STOP",
		Number:  "+33600000000",
	}

	observer.Stop(message, nil)

	exist, _ := _client.StopRequest.Query().Where(stoprequest.RecipientEQ("+33600000000")).Exist(ctx)

	if !exist {
		t.Error("Stop request need to be save in database")
	}
}

func TestStopObserver_Stop_Without_Capitalized_Message(t *testing.T) {
	ctx := context.Background()

	_client := tests.SetupDatabase(t)
	observer := CreateStopObserver(_client)

	message := &gsm.Message{
		Message: "stop",
		Number:  "+33600000000",
	}

	observer.Stop(message, nil)

	exist, _ := _client.StopRequest.Query().Where(stoprequest.RecipientEQ("+33600000000")).Exist(ctx)

	if !exist {
		t.Error("Stop request need to be save in database")
	}
}

func TestStopObserver_Stop_Ignore_Message(t *testing.T) {
	ctx := context.Background()

	_client := tests.SetupDatabase(t)
	observer := CreateStopObserver(_client)

	message := &gsm.Message{
		Message: "TEST",
		Number:  "+33600000000",
	}

	observer.Stop(message, nil)

	exist, _ := _client.StopRequest.Query().Where(stoprequest.RecipientEQ("+33600000000")).Exist(ctx)

	if exist {
		t.Error("Stop request doesnt exist in database")
	}
}
