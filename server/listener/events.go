package listener

type EvClientConnected struct{ Id int }

func (ev *EvClientConnected) EmitEvent(s *Service) {
	s.Emitter.Emit(ClientConnected, ev.Id)
}

type EvClientDisconnected struct{ Id int }

func (ev *EvClientDisconnected) EmitEvent(s *Service) {
	s.Emitter.Emit(ClientDisconnected, ev.Id)
}

type EvPacketReceived struct {
	Id   int
	Buff []byte
}

func (ev *EvPacketReceived) EmitEvent(s *Service) {
	s.Emitter.Emit(PacketReceived, ev)
}
