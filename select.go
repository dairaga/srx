// +build js,wasm

package srx

import (
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

func Select(owner TComponent) TSelect {
	return &sel{
		input: newFormControl(owner, "select", "form-select"),
		//opts:  nil,
	}
}
