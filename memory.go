// +build js,wasm

package srx

import (
	"sync"

	"github.com/dairaga/srx/js"
)

type memory struct {
	mux   *sync.Mutex
	table map[string]TComponent
}

var mem = &memory{
	mux:   &sync.Mutex{},
	table: make(map[string]TComponent),
}

func (m *memory) alloc(com TComponent) {
	defer m.mux.Unlock()
	m.mux.Lock()
	m.table[com.Tattoo()] = com
}

func (m *memory) free(com TComponent) {
	defer m.mux.Unlock()
	m.mux.Lock()
	delete(m.table, com.Tattoo())
}

func (m *memory) lookup(v js.Wrapper) TComponent {
	if tattoo := v.JSValue().Call("getAttribute", srxTattoo); tattoo.Truthy() {
		return m.table[tattoo.String()]
	}

	return nil
}

func Lookup(v js.Wrapper) TComponent {
	return mem.lookup(v)
}
