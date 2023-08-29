package handlers

import (
	"errors"
	"github.com/charmbracelet/log"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"sms-sender/ent"
	"sms-sender/ent/stoprequest"
	"sms-sender/sms"
	"strings"
)

type (
	SmsRequest struct {
		Message    string `json:"message" validate:"required"`
		Recipient  string `json:"recipient" validate:"required"`
		Commercial bool   `json:"is_commercial" validate:"required"`
	}

	SMSHandler struct {
		client    *ent.Client
		gsmSender sms.Sender
	}

	ValidationError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)

func CreateSMSHandler(client *ent.Client, sender *sms.GSMSender) SMSHandler {
	return SMSHandler{client: client, gsmSender: sender}
}

// SendSms POST - Send SMS to a recipient
func (s *SMSHandler) SendSms(c *fiber.Ctx) error {
	validate := validator.New()

	log.Info("Receive new request", "requestBody", string(c.Body()))

	smsRequest := new(SmsRequest)
	if err := c.BodyParser(smsRequest); err != nil {
		log.Error("Error on parsing sms request", "err", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on parsing request",
		})
	}

	err := validate.Struct(smsRequest)

	if err != nil {
		var validationErrors []ValidationError

		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidationError{
				Field:   err.Field(),
				Message: err.Tag(),
			})
		}

		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on field validation",
			"errors":  validationErrors,
		})
	}

	if err = validatePhoneNumber(smsRequest.Recipient); err != nil {
		log.Error("Invalid phone number for sms request", "err", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid phone number",
		})
	}

	clearedRecipient := cleanPhoneNumber(smsRequest.Recipient)

	smsRequest.Recipient = clearedRecipient

	smsEntity, err := s.client.Sms.
		Create().
		SetRecipient(smsRequest.Recipient).
		SetMessage(smsRequest.Message).
		SetCommercial(smsRequest.Commercial).
		Save(c.Context())

	if err != nil {
		log.Error("Error on save sms request in database", "err", err)
	}

	log.Info("New SMS request", "smsRequest", smsEntity.ID)

	if smsRequest.Commercial == true {
		result, err := s.client.StopRequest.
			Query().
			Where(stoprequest.RecipientEQ(smsRequest.Recipient)).
			Exist(c.Context())

		if err != nil {
			log.Error("Error on search stop request in database for sms request", "err", err, "smsRequest", smsEntity.ID)
		}

		if result {

			return c.Status(403).JSON(fiber.Map{
				"status": "stopped",
			})
		}
	}

	err = s.gsmSender.SendSms(smsEntity)

	if err != nil {
		log.Error("Error on sending sms with sms request", "err", err, "smsRequest", smsEntity.ID)

		_, err = smsEntity.Update().
			SetStatus("ERROR").
			Save(c.Context())

		if err != nil {
			log.Error("Error on save sms request fail in database", "smsRequest", smsEntity.ID)
		}

		return c.Status(500).JSON(fiber.Map{
			"status": "error on sending sms",
		})
	}

	_, err = smsEntity.Update().
		SetStatus("SUCCESS").
		Save(c.Context())

	if err != nil {
		log.Error("Error on update sms request in database", "err", err, "smsRequest", smsEntity.ID)
	}

	return c.JSON(fiber.Map{
		"status": "success",
	})
}

func validatePhoneNumber(phoneNumber string) error {
	validPrefixes := [4]string{"+3306", "+3307", "+336", "+337"}

	for _, validPrefix := range validPrefixes {
		if strings.HasPrefix(phoneNumber, validPrefix) {
			return nil
		}
	}

	return errors.New("invalid phone number")
}

func cleanPhoneNumber(phoneNumber string) string {

	phonePrefixes := map[string]string{"+3306": "+336", "+3307": "+337"}

	for dirtyPrefix, cleanPrefix := range phonePrefixes {
		if strings.HasPrefix(phoneNumber, dirtyPrefix) {
			return strings.Replace(phoneNumber, dirtyPrefix, cleanPrefix, 1)
		}
	}

	return phoneNumber
}
