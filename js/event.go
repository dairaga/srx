// +build js,wasm

package js

type (
	TEvent interface {
		Wrapper

		Type() string
		Target() TEventTarget
	}

	Event struct {
		ref    Value
		target *EventTarget
	}
)

var _ TEvent = &Event{}

func (e *Event) JSValue() Value {
	return e.ref
}

func (e *Event) Type() string {
	return e.ref.Get("type").String()
}

func (e *Event) Target() TEventTarget {
	return e.target
}

func EventOf(v Value) *Event {
	return &Event{
		ref:    v,
		target: EventTargetOf(v.Get("target")),
	}
}
