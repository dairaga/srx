// +build js,wasm

package srx

import (
	"testing"

	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestBadge(t *testing.T) {
	caption := `TestBadge`
	desc := `unread messages`

	b := Badge(Root())
	js.Append(b)

	/* init */
	assert.Equal(t, "", b.Value())
	assert.Equal(t, "", b.Description())
	assert.Equal(t, enum.Secondary, b.Color())

	assert.Equal(t, "", b.(*badge).Query("span").Text())
	assert.Equal(t, "", b.(*badge).Query(".visually-hidden").Text())
	assert.True(t, b.(*badge).Contains(enum.Secondary.Style("bg")))

	/* value */
	b.SetValue(caption)
	assert.Equal(t, caption, b.Value())
	assert.Equal(t, caption, b.(*badge).Query("span").Text())

	/* description */
	b.SetDescription(desc)
	assert.Equal(t, desc, b.Description())
	assert.Equal(t, desc, b.(*badge).Query(".visually-hidden").Text())

	/* color */
	colors := []enum.Color{enum.Light, enum.Warning, enum.Info}
	var old enum.Color
	for _, color := range colors {
		old = b.Color()
		b.SetColor(color)
		assert.Equal(t, color, b.Color())
		assert.True(t, b.(*badge).Contains(color.Style("bg")))
		assert.False(t, b.(*badge).Contains(old.Style("bg")))
		assert.True(t, b.(*badge).Contains("text-dark"))
	}
	old = b.Color()
	b.SetColor(enum.Primary)
	assert.True(t, b.(*badge).Contains(enum.Primary.Style("bg")))
	assert.False(t, b.(*badge).Contains(old.Style("bg")))
	assert.False(t, b.(*badge).Contains("text-dark"))

	/* type */
	assert.Equal(t, enum.RoundNone, b.Type())

	b.SetType(enum.RoundCircle)
	assert.Equal(t, enum.RoundCircle, b.Type())
	assert.True(t, b.(*badge).Contains(enum.RoundCircle.String()))

	b.SetType(enum.RoundPill)
	assert.Equal(t, enum.RoundPill, b.Type())
	assert.True(t, b.(*badge).Contains(enum.RoundPill.String()))

	b.SetType(enum.RoundNone)
	assert.Equal(t, enum.RoundNone, b.Type())
	assert.False(t, b.(*badge).Contains(enum.RoundCircle.String()))
	assert.False(t, b.(*badge).Contains(enum.RoundPill.String()))
}
