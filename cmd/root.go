package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/warthog618/modem/gsm"
	"log"
	"os"
	"sms-sender/ent"
	"sms-sender/ent/stoprequest"
	"sms-sender/sms"
	"strings"
)

type SmsRequest struct {
	Message    string `json:"message"`
	Recipient  string `json:"recipient"`
	Commercial bool   `json:"is_commercial"`
}

var rootCmd = &cobra.Command{
	Use:   "sms-sender",
	Short: "Serve sms api",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		bCtx := context.Background()

		client, err := ent.Open("sqlite3", "file:sms_sender.db?_fk=1")
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}

		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		config := sms.GsmConfig{
			Port: os.Getenv("SERIAL_PORT"),
			Baud: 115200,
		}

		gsmSender := sms.GSMSender{}

		err = gsmSender.Init(config)

		if err != nil {
			log.Fatalln(err)
		}

		handler := func(msg gsm.Message) {
			if strings.ToUpper(msg.Message) != "STOP" {
				return
			}

			result, err := client.StopRequest.
				Query().
				Where(stoprequest.RecipientEQ(msg.Number)).
				Exist(bCtx)

			if err == nil {
				log.Println(err)
			}

			if !result {
				_, err = client.StopRequest.
					Create().
					SetRecipient(msg.Number).
					Save(bCtx)

				if err != nil {
					log.Println(err)
				}
			}
		}

		errorHandler := func(err error) {
			log.Println(err)
		}

		err = gsmSender.GetGSM().StartMessageRx(handler, errorHandler)

		app := fiber.New()

		app.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				os.Getenv("BASIC_AUTH_USER"): os.Getenv("BASICAUTH_PASSWORD"),
			},
		}))

		app.Post("/sms", func(c *fiber.Ctx) error {
			smsRequest := new(SmsRequest)

			if err := c.BodyParser(smsRequest); err != nil {
				log.Println(err)
				return c.Status(500).JSON(fiber.Map{
					"status":  "error",
					"message": "Error on parsing request",
				})
			}

			if err = validatePhoneNumber(smsRequest.Recipient); err != nil {
				log.Println(err)
				return c.Status(400).JSON(fiber.Map{
					"status":  "error",
					"message": "Invalid phone number",
				})
			}

			clearedRecipient := cleanPhoneNumber(smsRequest.Recipient)

			smsRequest.Recipient = clearedRecipient

			if smsRequest.Commercial == true {
				log.Println(smsRequest.Recipient)
				result, err := client.StopRequest.
					Query().
					Where(stoprequest.RecipientEQ(smsRequest.Recipient)).
					Exist(bCtx)

				if err != nil {
					log.Println(err)
				}

				if result {
					return c.Status(403).JSON(fiber.Map{
						"status": "stopped",
					})
				}
			}

			sms, err := client.Sms.
				Create().
				SetRecipient(smsRequest.Recipient).
				SetMessage(smsRequest.Message).
				SetCommercial(smsRequest.Commercial).
				Save(bCtx)

			if err != nil {
				log.Println(err)
			}

			err = gsmSender.SendSms(sms)

			if err != nil {
				log.Println(err)

				_, err = sms.Update().
					SetStatus("ERROR").
					Save(bCtx)

				if err != nil {
					log.Println(err)
				}
			}

			_, err = sms.Update().
				SetStatus("SUCCESS").
				Save(bCtx)

			if err != nil {
				log.Println(err)
			}

			return c.JSON(fiber.Map{
				"status": "success",
			})
		})

		err = app.Listen(":3000")

		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
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
