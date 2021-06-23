// +build js,wasm

package srx

import (
	"testing"
	"time"

	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestAlert(t *testing.T) {
	a := Alert(Root())
	a.Append(Caption("test alert 1"), Caption("test alert 2"))
	a.SetColor(enum.Primary)
	js.Append(a)

	ch := make(chan int)
	defer close(ch)

	a.OnClosing(func(sender TAlert, _ js.TEvent) {
		assert.Equal(t, a, sender)
	})

	a.OnClosed(func(sender TAlert, _ js.TEvent) {
		assert.Equal(t, sender, a)
		ch <- 1
	})
	a.Close()
	select {
	case <-ch:
	case <-time.After(5 * time.Second):
		t.Error("test alert failure")
	}
	assert.Nil(t, Lookup(a))
}
