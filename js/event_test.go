// +build js,wasm

package js_test

import (
	"testing"

	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {

	btn := js.Create("button")

	fn := js.FuncOf(func(_this js.Value, _args []js.Value) interface{} {
		evt := js.EventOf(_args[0])
		assert.Equal(t, "click", evt.Type())
		assert.Equal(t, true, evt.Target().JSValue().Equal(btn.JSValue()))
		return nil
	})

	btn.On("click", fn)
	btn.Call("click")
}
