package client

import (
	"log"

	"github.com/carabalonepaulo/origin/shared/emitter"
	"github.com/carabalonepaulo/origin/shared/service"
	"github.com/codecat/go-enet"
	"github.com/kelindar/binary"
)

const (
	Connected = iota
	Disconnected
)

type Service struct {
	emitter.Emitter

	host enet.Host
	peer enet.Peer

	cmds chan Command
	evs  chan any

	ip      string
	port    uint16
	running bool
}

func New(ip string, port uint16) func() service.Service {
	return func() service.Service {
		return &Service{
			ip:   ip,
			port: port,
		}
	}
}

func (s *Service) Start(services service.Services, shutdown func()) error {
	enet.Initialize()

	host, err := enet.NewHost(nil, 1, 1, 0, 0)
	if err != nil {
		enet.Deinitialize()
		return err
	}

	peer, err := host.Connect(enet.NewAddress("127.0.0.1", 5051), 1, 0)
	if err != nil {
		enet.Deinitialize()
		return err
	}

	s.host = host
	s.peer = peer

	return nil
}

func (s *Service) Stop() {
	s.running = false
}

func (s *Service) Update(_ float64) {
	// select {
	// case cmd := <-s.cmds:
	// 	switch cmd := cmd.(type) {
	// 	case CmdSend:

	// 	}
	// }
}

func (s *Service) Send(value any) {
	buff, err := binary.Marshal(value)
	if err != nil {
		log.Println(err)
		return
	}

	s.cmds <- CmdSend(buff)
}

func (s *Service) loop() {
	for s.running {
		ev := s.host.Service(0)
		switch ev.GetType() {
		case enet.EventConnect:
		case enet.EventDisconnect:
			// s.evs <- broker.Message{Kind: Disconnected, Content: }
			// s.evs <- EvConnected{Id}
		case enet.EventReceive:
		}

		select {
		case cmd := <-s.cmds:
			cmd.ProcessCommand(s)
		}
	}
	s.host.Destroy()
	enet.Deinitialize()
}
