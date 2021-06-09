// +build js,wasm

package std

import (
	"testing"

	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestLabel(t *testing.T) {
	f := ""
	caption := ""

	lb := LabelOf(srx.Root())
	js.Append(lb)

	/* init */
	assert.Equal(t, f, lb.For())
	assert.Equal(t, caption, lb.Caption())
	assert.Equal(t, f, lb.(*label).Attr("for"))
	assert.Equal(t, caption, lb.(*label).Text())

	f = "test_for"
	caption = "test_label"
	lb.SetFor(f)
	lb.SetCaption(caption)

	assert.Equal(t, f, lb.For())
	assert.Equal(t, caption, lb.Caption())
	assert.Equal(t, f, lb.(*label).Attr("for"))
	assert.Equal(t, caption, lb.(*label).Text())
}
