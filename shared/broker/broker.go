package broker

import "github.com/carabalonepaulo/origin/shared/service"

type Message struct {
	Kind    int
	Content any
}

type Event[S service.Service] interface {
	EmitEvent(S)
}

type Command[S service.Service] interface {
	ProcessCommand(S)
}
