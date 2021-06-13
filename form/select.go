// +build js,wasm

package form

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TSelect interface {
		TBaseFormControl
		AddOption(caption, value string)
	}

	sel struct {
		*input
		//opts [][2]string
	}
)

var _ TSelect = &sel{}

// -----------------------------------------------------------------------------

func (s *sel) AddOption(caption, value string) {
	//s.opts = append(s.opts, [2]string{value, caption})
	opt := js.Create("option")
	opt.SetAttr("value", value)
	opt.SetText(caption)
	s.Element.Append(opt)
}

// -----------------------------------------------------------------------------

func SelectOf(owner srx.TComponent) TSelect {
	return &sel{
		input: newInput(owner, "select", "form-select"),
		//opts:  nil,
	}
}
