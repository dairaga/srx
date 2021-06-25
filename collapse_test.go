// +build js,wasm

package srx

import (
	"fmt"
	"testing"

	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestCollapse(t *testing.T) {
	anchor := js.From(js.HTML(`<a class="btn btn-primary">Link with href</a>`))
	btn := js.From(js.HTML(`<button class="btn btn-primary" type="button">Button with data-bs-target</button>`))

	js.Append(anchor)
	js.Append(btn)

	c := Collapse(Root())
	id := `test_collapse`
	c.SetID(id)
	c.Append(newObject(js.From(`<div class="card card-body">Some placeholder content for the collapse component. This panel is hidden by default but revealed when the user activates the relevant trigger.</div>`)))
	c.Bind(anchor)
	c.Bind(btn)
	Append(c)

	assert.Equal(t, "#"+id, anchor.Attr("href"))
	assert.Equal(t, "collapse", anchor.Attr("data-bs-toggle"))
	assert.Equal(t, "false", anchor.Attr("aria-expanded"))
	assert.Equal(t, id, anchor.Attr("aria-controls"))

	assert.Equal(t, "#"+id, btn.Attr("data-bs-target"))
	assert.Equal(t, "collapse", btn.Attr("data-bs-toggle"))
	assert.Equal(t, "false", btn.Attr("aria-expanded"))
	assert.Equal(t, id, btn.Attr("aria-controls"))

	assert.False(t, c.Ref().Contains("show"))

	ch := make(chan int, 1)

	c.OnShowing(func(sender TCollapse, evt js.TEvent) {
		fmt.Println(evt.Type())
		assert.Equal(t, c, sender)
	})

	c.OnShown(func(sender TCollapse, evt js.TEvent) {
		fmt.Println(evt.Type())
		assert.Equal(t, c, sender)
		ch <- 1
	})

	c.OnHiding(func(sender TCollapse, evt js.TEvent) {
		fmt.Println(evt.Type())
		assert.Equal(t, c, sender)
	})

	c.OnHidden(func(sender TCollapse, evt js.TEvent) {
		fmt.Println(evt.Type())
		assert.Equal(t, c, sender)
		ch <- 1
	})

	c.Show()
	<-ch

	c.Hide()
	<-ch

}
