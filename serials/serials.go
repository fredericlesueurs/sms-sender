package serials

import (
	"fmt"
	"github.com/charmbracelet/log"
	"go.bug.st/serial"
)

type Serials struct {
	serials []string
}

func (s *Serials) UpdateSerials() {
	ports, err := serial.GetPortsList()

	if err != nil {
		log.Fatal("Failed on collect Serial ports", "err", err)
	}

	log.Info(fmt.Sprintf("%v serials ports find", len(ports)))

	s.serials = ports
}

func (s *Serials) GetSerials() []string {
	if len(s.serials) == 0 {
		s.UpdateSerials()
	}

	return s.serials
}
