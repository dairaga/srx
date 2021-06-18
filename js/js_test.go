// +build js,wasm

package js_test

import (
	"os"
	"os/signal"
	"testing"

	gojs "syscall/js"

	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func query(selector string) gojs.Value {
	return gojs.Global().Get("document").Call("querySelector", selector)
}

func TestMain(m *testing.M) {
	headless := os.Getenv("WASM_HEADLESS")
	exitVal := m.Run()

	if headless == "off" {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		// Block until a signal is received.
		<-c
	}

	os.Exit(exitVal)
}

func TestCreate(t *testing.T) {
	div := js.Create("div")
	div.SetText("abc")
	assert.Equal(t, "abc", div.Text())
}

func TestAppend(t *testing.T) {
	div := js.Create("div")
	div.SetAttr("id", "test_append")
	js.Append(div)

	test := query("#test_append")
	assert.Equal(t, "test_append", test.Call("getAttribute", "id").String())
}

func TestIsJSObject(t *testing.T) {
	assert.False(t, js.IsJSObject(js.ValueOf(true)))
	assert.True(t, js.IsJSObject(js.Create("div")))
}

func TestIsJSFunc(t *testing.T) {
	assert.False(t, js.IsJSFunc(js.ValueOf(true)))
	fn := js.FuncOf(func(this gojs.Value, args []gojs.Value) interface{} {
		return nil
	})
	assert.True(t, js.IsJSFunc(fn))
}
