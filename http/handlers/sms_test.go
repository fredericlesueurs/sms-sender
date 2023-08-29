package handlers

import (
	"bytes"
	"context"
	"errors"
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"net/http/httptest"
	"sms-sender/ent"
	"sms-sender/ent/sms"
	"sms-sender/tests"
	"testing"
)

type SMSHandlerMock struct {
}

func (s *SMSHandlerMock) SendSms(sms *ent.Sms) error {
	log.Info("Send SMS called", "sms", sms)

	return nil
}

type SMSHandlerMockError struct {
}

func (s *SMSHandlerMockError) SendSms(sms *ent.Sms) error {
	log.Info("Send SMS called", "sms", sms)
	log.Info("Simulate an error")

	return errors.New("error")
}

func TestSMSHandler_SendSms(t *testing.T) {
	ctx := context.Background()
	_client := tests.SetupDatabase(t)
	sender := &SMSHandlerMock{}

	_handler := SMSHandler{client: _client, gsmSender: sender}

	app := fiber.New()
	app.Post("/sms", _handler.SendSms)

	jsonBody := []byte(`{"message": "Coucou tests 1", "recipient": "+33604048319", "is_commercial": true}`)
	bodyReader := bytes.NewReader(jsonBody)
	req := httptest.NewRequest("POST", "/sms", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Error("Status must be 200", "status:", resp.StatusCode)
	}

	exist, _ := _client.Sms.Query().Where(sms.MessageEQ("Coucou tests 1")).Exist(ctx)

	if !exist {
		t.Error("Message not found in database")
	}
}

func TestSMSHandler_SendSms_With_Wrong_Parameters(t *testing.T) {
	_client := tests.SetupDatabase(t)
	sender := &SMSHandlerMock{}

	_handler := SMSHandler{client: _client, gsmSender: sender}

	app := fiber.New()
	app.Post("/sms", _handler.SendSms)

	jsonBody := []byte(`{"messages": "Coucou", "recipient": "+33604048319", "is_commercial": true}`)
	bodyReader := bytes.NewReader(jsonBody)
	req := httptest.NewRequest("POST", "/sms", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	if resp.StatusCode != 400 {
		t.Error("Status must be 400", "status:", resp.StatusCode)
	}
}

func TestSMSHandler_SendSms_Error_With_SMS_Sending(t *testing.T) {
	_client := tests.SetupDatabase(t)
	sender := &SMSHandlerMockError{}

	_handler := SMSHandler{client: _client, gsmSender: sender}

	app := fiber.New()
	app.Post("/sms", _handler.SendSms)

	jsonBody := []byte(`{"message": "Coucou", "recipient": "+33604048319", "is_commercial": true}`)
	bodyReader := bytes.NewReader(jsonBody)
	req := httptest.NewRequest("POST", "/sms", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	if resp.StatusCode != 500 {
		t.Error("Status must be 500", "status:", resp.StatusCode)
	}
}
