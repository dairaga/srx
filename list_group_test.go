// +build js,wasm

package srx

import (
	"testing"

	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestListGroup(t *testing.T) {
	lst := ListGroup(Root())
	items := []TListItem{}
	items = append(items, lst.AddItem(Caption("An item")))
	items = append(items, lst.AddItem(Caption("A second item")))
	items = append(items, lst.AddItem(Caption("A third item")))
	items = append(items, lst.AddItem(Caption("A fourth item")))
	items = append(items, lst.AddItem(Caption("And a fifth one")))

	js.Append(lst)

	assert.EqualValues(t, items, lst.Items())

	assert.True(t, lst.Ref().Contains("list-group"))
	assert.True(t, items[0].Ref().Contains("list-group-item"))
	lst.Active(3)
	assert.Equal(t, 3, lst.Actived())
	assert.True(t, items[3].Ref().Contains("active"))

	items[0].Disable()
	assert.True(t, items[0].Ref().Contains("disabled"))
	assert.True(t, items[0].Disabled())

	lst.Actionable(true)
	assert.True(t, lst.Actionabled())
	assert.True(t, items[1].Ref().Contains("list-group-item-action"))

	lst.Actionable(false)
	assert.False(t, lst.Actionabled())
	assert.False(t, items[1].Ref().Contains("list-group-item-action"))

	lst.Flush(true)
	assert.True(t, lst.Ref().Contains("list-group-flush"))
	assert.True(t, lst.Flushed())

	lst.Flush(false)
	assert.False(t, lst.Ref().Contains("list-group-flush"))
	assert.False(t, lst.Flushed())

	lst = OrderedListGroup(Root())
	js.Append(lst)

	assert.Equal(t, "ol", lst.Ref().TagName())
	assert.True(t, lst.Ref().Contains("list-group"))
	assert.True(t, lst.Ref().Contains("list-group-numbered"))
}
