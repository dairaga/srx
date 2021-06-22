// +build js,wasm

package srx

import (
	"fmt"

	"github.com/dairaga/srx/js"
)

type (
	TCaption interface {
		TObject
		Caption() string
		SetCaption(caption string)
	}

	TStrong = TCaption
	TSmall  = TCaption

	caption struct {
		*object
	}
)

var _ TCaption = &caption{}

func (c *caption) Caption() string {
	return c.Element.Text()
}

func (c *caption) SetCaption(caption string) {
	c.Element.SetText(caption)
}

func newCaption(tag, content string) *caption {
	ret := &caption{
		object: newObject(js.Create(tag)),
	}
	ret.SetCaption(content)
	return ret
}

func Caption(a ...interface{}) TCaption {
	return newCaption("span", fmt.Sprint(a...))
}

func Captionf(format string, a ...interface{}) TCaption {
	return newCaption("span", fmt.Sprintf(format, a...))
}

func Strong(a ...interface{}) TStrong {
	return newCaption("strong", fmt.Sprint(a...))
}

func Strongf(format string, a ...interface{}) TStrong {
	return newCaption("strong", fmt.Sprintf(format, a...))
}

func Small(a ...interface{}) TSmall {
	return newCaption("small", fmt.Sprint(a...))
}

func Smallf(format string, a ...interface{}) TSmall {
	return newCaption("small", fmt.Sprintf(format, a...))
}
