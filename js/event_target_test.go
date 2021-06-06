// +build js,wasm

package js_test

import (
	"testing"

	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestEventTarget(t *testing.T) {
	caption := "test_event_target"
	count := 0

	btn := js.Create("button")

	fn := js.FuncOf(func(_this js.Value, _args []js.Value) interface{} {
		_this.Set("innerText", caption)
		count += 1
		return nil
	})

	btn.On("click", fn)
	btn.Call("click")
	assert.Equal(t, caption, btn.Text())

	// fn released, and count must be same as before.
	btn.Release()
	btn.Call("click")
	assert.Equal(t, 1, count)
}
