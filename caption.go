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

func newCaption(content string) *caption {
	ret := &caption{
		object: newObject(js.Create("span")),
	}
	ret.SetCaption(content)
	return ret
}

func Caption(a ...interface{}) TCaption {
	return newCaption(fmt.Sprint(a...))
}

func Captionf(format string, a ...interface{}) TCaption {
	return newCaption(fmt.Sprintf(format, a...))
}
