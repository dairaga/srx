// +build js,wasm

package js

type (
	TEventTarget interface {
		Wrapper

		On(event string, fn Func)

		Release()
	}

	EventTarget struct {
		ref   Value
		funcs map[string]Func
	}
)

var _ TEventTarget = &EventTarget{}

func (t *EventTarget) JSValue() Value {
	return t.ref
}

func (t *EventTarget) On(event string, fn Func) {
	if t.funcs == nil {
		t.funcs = make(map[string]Func)
	}

	old, ok := t.funcs[event]
	if ok && old.Truthy() {
		old.Release()
	}
	t.funcs[event] = fn
	t.ref.Call("addEventListener", event, fn)
}

func (t *EventTarget) Release() {
	for k, v := range t.funcs {
		if v.Truthy() {
			v.Release()
		}
		delete(t.funcs, k)
	}
}

func EventTargetOf(v Value) *EventTarget {
	return &EventTarget{
		ref:   v,
		funcs: nil,
	}
}
