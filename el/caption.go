// +build js,wasm

package el

import (
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

func Caption() TCaption {
	return &caption{
		Object: srx.NewObject(js.Create("span")),
	}
}
