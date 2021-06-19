// +build js,wasm

package srx

import (
	"testing"

	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestProgress(t *testing.T) {

	p := Progress(Root())
	p.Stripe(true)
	p.Animate(true)
	js.Append(p)

	bar1 := ProgressBar(p)
	bar1.SetColor(enum.Primary)
	bar1.SetValue(50)
	bar1.SetCaption("abc")
	p.AddBar(bar1)

	assert.Equal(t, "abc", bar1.Caption())
	assert.Equal(t, 50, bar1.Value())
	assert.True(t, bar1.Animated())
	assert.True(t, bar1.Striped())
	assert.True(t, bar1.Ref().Contains("progress-bar-striped"))
	assert.True(t, bar1.Ref().Contains("progress-bar-animated"))
	assert.True(t, bar1.Ref().Contains(enum.Primary.Style("bg")))

	bar2 := ProgressBar(p)
	bar2.SetColor(enum.Info)
	bar2.SetValue(25)
	bar2.SetCaption("def")
	p.AddBar(bar2)

	assert.Equal(t, "def", bar2.Caption())
	assert.Equal(t, 25, bar2.Value())
	assert.True(t, bar2.Animated())
	assert.True(t, bar2.Striped())
	assert.True(t, bar2.Ref().Contains("progress-bar-striped"))
	assert.True(t, bar2.Ref().Contains("progress-bar-animated"))
	assert.True(t, bar2.Ref().Contains(enum.Info.Style("bg")))

}
