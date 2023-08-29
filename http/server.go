package http

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"os"
	"sms-sender/ent"
	"sms-sender/http/handlers"
	"sms-sender/sms"
)

type Server struct {
	Client    *ent.Client
	GSMSender *sms.GSMSender
}

func (s *Server) Start() {
	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("BASIC_AUTH_USER"): os.Getenv("BASICAUTH_PASSWORD"),
		},
	}))

	smsHandler := handlers.CreateSMSHandler(s.Client, s.GSMSender)

	app.Post("/sms", smsHandler.SendSms)

	err := app.Listen(":3000")

	if err != nil {
		log.Fatal("Error on start server", "err", err)
	}
}
