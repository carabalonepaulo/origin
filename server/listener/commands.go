package listener

type CmdSend struct {
	Id   int
	Buff []byte
}

func (cmd *CmdSend) ProcessCommand(s *Service) {}

type CmdKick struct {
	Id int
}

func (cmd *CmdKick) ProcessCommand(s *Service) {}

type CmdShutdown struct{}

func (cmd *CmdShutdown) ProcessCommand(s *Service) {
	s.running = false
}
