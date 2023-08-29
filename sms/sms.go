package sms

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/warthog618/modem/at"
	"github.com/warthog618/modem/gsm"
	"github.com/warthog618/modem/serial"
	smsSender "github.com/warthog618/sms"
	"io"
	"sms-sender/ent"
	"time"
)

type GsmConfig struct {
	Port string
	Baud int
}

type Observer func(message *gsm.Message, err error)

type Subject interface {
	Register(observer Observer)
	notifyAll(message *gsm.Message, err error)
}

type Sender interface {
	SendSms(sms *ent.Sms) error
}

type Handler interface {
	Register(observer Observer)
	notifyAll(message gsm.Message, err error)
}

type GSMSender struct {
	gsm       *gsm.GSM
	observers []Observer
}

func CreateGSMSender(gsmConfig GsmConfig) (*GSMSender, error) {
	log.Info("Connecting to GSM modem...")
	m, err := serial.New(serial.WithPort(gsmConfig.Port), serial.WithBaud(gsmConfig.Baud))

	if err != nil {
		return nil, err
	}

	var mio io.ReadWriter = m

	g := gsm.New(at.New(mio, at.WithTimeout(5*time.Second)))

	_, err = g.Command("+CREG?")
	if err != nil {
		return nil, err
	}

	log.Info("Connected to GSM modem")

	sender := GSMSender{gsm: g}

	handler := func(message gsm.Message) {
		sender.notifyAll(&message, nil)
	}

	errorHandler := func(error error) {
		sender.notifyAll(nil, error)
	}

	err = g.StartMessageRx(handler, errorHandler)

	log.Info("SMS handler ready")

	return &sender, nil
}

func (gsmSender *GSMSender) SendSms(sms *ent.Sms) error {
	pdus, err := smsSender.Encode([]byte(sms.Message), smsSender.To(sms.Recipient), smsSender.WithAllCharsets)

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()

	for _, p := range pdus {
		tp, err := p.MarshalBinary()
		if err != nil {
			return err
		}
		_, err = gsmSender.gsm.SendPDU(tp)
		if err != nil {
			// !!! check CPIN?? on failure to determine root cause??  If ERROR 302
			return err
		}
	}

	processedTime := time.Now().Sub(start).Seconds()

	log.Info(fmt.Sprintf("SMS send in %f seconds", processedTime), "smsRequest", sms.ID)

	return nil
}

func (gsmSender *GSMSender) Register(observer Observer) {
	gsmSender.observers = append(gsmSender.observers, observer)
}

func (gsmSender *GSMSender) notifyAll(message *gsm.Message, err error) {
	for _, observer := range gsmSender.observers {
		observer(message, err)
	}
}
