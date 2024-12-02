package listener

import (
	"slices"
	"time"

	"github.com/carabalonepaulo/origin/shared/broker"
	"github.com/carabalonepaulo/origin/shared/emitter"
	"github.com/carabalonepaulo/origin/shared/service"
	"github.com/carabalonepaulo/origin/shared/services/scheduler"
	"github.com/codecat/go-enet"
)

const (
	ClientConnected = iota
	ClientDisconnected
	PacketReceived
)

type Config struct {
	Port         uint16 `json:"port"`
	MaxClients   int    `json:"max_clients"`
	InLimit      int    `json:"in_limit"`
	OutLimit     int    `json:"out_limit"`
	Channels     uint64 `json:"channels"`
	TickInterval string `json:"tick_interval"`
}

type Service struct {
	emitter.Emitter

	config  *Config
	port    uint16
	host    enet.Host
	cmds    chan broker.Command[*Service]
	evs     chan broker.Event[*Service]
	peers   []enet.Peer
	running bool
}

func New(config *Config) func() service.Service {
	return func() service.Service {
		return &Service{
			Emitter: emitter.Init(3),
			config:  config,
			port:    config.Port,
			cmds:    make(chan broker.Command[*Service]),
			evs:     make(chan broker.Event[*Service]),
			peers:   make([]enet.Peer, config.MaxClients),
		}
	}
}

func (s *Service) Start(services service.Services, shutdown func()) error {
	interval, err := time.ParseDuration(s.config.TickInterval)
	if err != nil {
		return err
	}

	scheduler, err := service.Get[*scheduler.Service](services)
	if err != nil {
		return err
	}
	scheduler.Every(interval, s.Poll)

	enet.Initialize()
	host, err := enet.NewHost(enet.NewListenAddress(s.port), 32, s.config.Channels, 0, 0)
	if err != nil {
		enet.Deinitialize()
		return err
	}

	s.host = host
	s.running = true

	go s.loop()
	return nil
}

func (s *Service) Stop() {
	s.running = false
}

func (s *Service) Poll() {
	select {
	case ev := <-s.evs:
		ev.EmitEvent(s)
	default:
	}
}

func (s *Service) loop() {
	for s.running {
		s.handleEvents()
		s.handleCommands()
	}
	s.host.Destroy()
	enet.Deinitialize()
}

func (s *Service) handleEvents() {
	ev := s.host.Service(0)
	switch ev.GetType() {
	case enet.EventConnect:
		id := slices.Index(s.peers, nil)
		s.peers[id] = ev.GetPeer()
		s.peers[id].SetData(intToBytes(id))
		s.evs <- &EvClientConnected{Id: id}
		// TODO: send id

	case enet.EventReceive:
		packet := ev.GetPacket()
		defer packet.Destroy()

		buff := packet.GetData()
		id := bytesToInt(ev.GetPeer().GetData())
		s.evs <- &EvPacketReceived{Id: id, Buff: buff}

	case enet.EventDisconnect:
		id := bytesToInt(ev.GetPeer().GetData())
		s.evs <- &EvClientDisconnected{Id: id}

	case enet.EventNone:
		return
	}
}

func (s *Service) handleCommands() {
	select {
	case cmd := <-s.cmds:
		cmd.ProcessCommand(s)
	default:
	}
}
