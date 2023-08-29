package cmd

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"sms-sender/ent"
	"sms-sender/http"
	"sms-sender/observers"
	"sms-sender/serials"
	"sms-sender/sms"
)

var UserProvidedPort string

var rootCmd = &cobra.Command{
	Use:   "sms-sender",
	Short: "Serve sms api",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		p := viper.GetString("port")

		if p == "" {
			s := serials.Serials{}

			prompt := promptui.Select{
				Label: "Select an port",
				Items: s.GetSerials(),
			}

			_, p, err = prompt.Run()
		}

		client := connectToDatabase()

		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatal("failed creating schema resources", "err", err)
		}

		config := sms.GsmConfig{
			Port: p,
			Baud: 115200,
		}

		gsmSender, err := sms.CreateGSMSender(config)

		if err != nil {
			log.Fatal("Error on create gsm sender", "err", err)
		}

		stopObserver := observers.CreateStopObserver(client)
		gsmSender.Register(stopObserver.Stop)

		server := http.Server{Client: client, GSMSender: gsmSender}

		server.Start()
	},
}

func connectToDatabase() *ent.Client {
	client, err := ent.Open("sqlite3", "file:sms_sender.db?_fk=1")

	if err != nil {
		log.Fatal("Failed to opening connection to sqlite", "err", err)
	}

	log.Info("Connected to the database")

	return client
}

func init() {
	rootCmd.Flags().StringVarP(&UserProvidedPort, "port", "p", "", "Set the port of Modem")
	viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
