package observers

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/warthog618/modem/gsm"
	"sms-sender/ent"
	"sms-sender/ent/stoprequest"
	"strings"
)

type StopObserver struct {
	client *ent.Client
}

func CreateStopObserver(client *ent.Client) StopObserver {
	return StopObserver{client: client}
}

func (so *StopObserver) Stop(message *gsm.Message, _ error) {
	c := context.Background()

	if strings.ToUpper(message.Message) != "STOP" {
		return
	}

	log.Info("New stop request received")

	result, err := so.client.StopRequest.
		Query().
		Where(stoprequest.RecipientEQ(message.Number)).
		Exist(c)

	if err != nil {
		log.Error("Error on search stop request in database", "err", err)
	}

	if !result {
		_, err = so.client.StopRequest.
			Create().
			SetRecipient(message.Number).
			Save(c)

		if err != nil {
			log.Error("Error on save stop request in database", "err", err)
		}
	}
}
