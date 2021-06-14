// +build js,wasm

package el

import (
	"fmt"

	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TCaption interface {
		srx.TObject
		Caption() string
		SetCaption(caption string)
	}

	caption struct {
		*srx.Object
	}
)

var _ TCaption = &caption{}

// -----------------------------------------------------------------------------

func (c *caption) Caption() string {
	return c.Element.Text()
}

// -----------------------------------------------------------------------------

func (c *caption) SetCaption(caption string) {
	c.Element.SetText(caption)
}

// -----------------------------------------------------------------------------

func newCaption(content string) *caption {
	ret := &caption{
		Object: srx.NewObject(js.Create("span")),
	}
	ret.SetCaption(content)
	return ret
}

// -----------------------------------------------------------------------------

func Caption(a ...interface{}) TCaption {
	return newCaption(fmt.Sprint(a...))
}

// -----------------------------------------------------------------------------

func Captionf(format string, a ...interface{}) TCaption {
	return newCaption(fmt.Sprintf(format, a...))
}
