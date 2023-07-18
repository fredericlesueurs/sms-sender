package sms

import (
	"github.com/warthog618/modem/at"
	"github.com/warthog618/modem/gsm"
	"github.com/warthog618/modem/serial"
	smsSender "github.com/warthog618/sms"
	"io"
	"log"
	"sms-sender/ent"
	"time"
)

type GsmConfig struct {
	Port string
	Baud int
}

type GSMSender struct {
	gsm *gsm.GSM
}

func (gsmSender *GSMSender) Init(gsmConfig GsmConfig) error {
	log.Println("Connecting to GSM modem...")
	m, err := serial.New(serial.WithPort(gsmConfig.Port), serial.WithBaud(gsmConfig.Baud))

	if err != nil {
		return err
	}

	var mio io.ReadWriter = m

	g := gsm.New(at.New(mio, at.WithTimeout(5*time.Second)))

	_, err = g.Command("+CREG?")
	if err != nil {
		return err
	}

	log.Println("Connected to GSM modem")

	gsmSender.gsm = g

	return nil
}

func (gsmSender *GSMSender) SendSms(sms *ent.Sms) error {
	pdus, err := smsSender.Encode([]byte(sms.Message), smsSender.To(sms.Recipient), smsSender.WithAllCharsets)

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()

	for i, p := range pdus {
		tp, err := p.MarshalBinary()
		if err != nil {
			return err
		}
		mr, err := gsmSender.gsm.SendPDU(tp)
		if err != nil {
			// !!! check CPIN?? on failure to determine root cause??  If ERROR 302
			return err
		}
		log.Printf("PDU %d: %v\n", i+1, mr)
	}

	processedTime := time.Now().Sub(start).Seconds()

	log.Printf("Sended took %f seconds\n", processedTime)

	return nil
}

func (gsmSender *GSMSender) GetGSM() *gsm.GSM {
	return gsmSender.gsm
}
