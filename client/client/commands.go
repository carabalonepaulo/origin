package client

import "github.com/codecat/go-enet"

type Command interface {
	ProcessCommand(*Service)
}

type CmdSend []byte

func (cmd CmdSend) ProcessCommand(s *Service) {
	s.peer.SendBytes(cmd, 1, enet.PacketFlagReliable)
}
