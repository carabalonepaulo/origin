package client

type EvConnected struct{ Id int }

func (ev *EvConnected) EmitEvent(s *Service) {
	s.Emitter.Emit(Connected, ev.Id)
}

type EvDisconnected struct{ Id int }

func (ev *EvDisconnected) EmitEvent(s *Service) {
	s.Emitter.Emit(Disconnected, ev.Id)
}
