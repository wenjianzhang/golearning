package eventcenter

import (
	"fmt"
	"testing"
)

type (
	Event struct {
		Msg string
	}

	Observcer interface {
		OnNotify(Event)
	}

	Notifier interface {
		Register(Observcer)

		UNRegister(Observcer)

		Notifier(Event)
	}

	EventCenter struct {
		observcers []Observcer
	}

	EventReceiver struct {
	}
)

func (e *EventCenter) Register(observcer Observcer) {
	e.observcers = append(e.observcers, observcer)
}

func (e *EventCenter) UNRegister(observcer Observcer) {
	for i, o := range e.observcers {
		if o == observcer {
			e.observcers = append(e.observcers[:i], e.observcers[i+1:]...)
			break
		}
	}
}

func (e *EventCenter) Notify(event Event) {
	for _, o := range e.observcers {
		o.OnNotify(event)
	}
}

func (e EventReceiver) OnNotify(event Event) {
	fmt.Printf("Receiver receive: %s\n", event.Msg)
}

func NewEventCenter() *EventCenter {
	ec := EventCenter{}
	ec.observcers = make([]Observcer, 0)
	return &ec
}

func TestEventCenter(T *testing.T) {
	eventCenter := NewEventCenter()
	recover1 := EventReceiver{}
	recover2 := EventReceiver{}

	eventCenter.Register(recover1)
	eventCenter.Register(recover2)

	eventCenter.Notify(Event{Msg: "hello go"})

	eventCenter.UNRegister(recover1)

	eventCenter.Notify(Event{"hello docker"})
}
