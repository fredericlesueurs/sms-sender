package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.bug.st/serial"
	"log"
)

func init() {
	rootCmd.AddCommand(listSerialCmd)
}

var listSerialCmd = &cobra.Command{
	Use:   "list-serial",
	Short: "List all serial connected",
	Run: func(cmd *cobra.Command, args []string) {
		ports, err := serial.GetPortsList()

		if err != nil {
			log.Fatal(err)
		}
		if len(ports) == 0 {
			log.Fatal("No serial ports found!")
		}

		for _, port := range ports {
			fmt.Printf("Found port: %v\n", port)
		}
	},
}
