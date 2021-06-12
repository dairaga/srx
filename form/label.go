// +build js,wasm

package form

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TLabel interface {
		srx.TComponent
		For() string
		SetFor(f string)

		Caption() string
		SetCaption(c string)

		Horizontal() bool
		SetHorizontal(h bool)
	}

	label struct {
		*srx.Component
	}
)

var _ TLabel = &label{}

// -----------------------------------------------------------------------------

func (lb *label) For() string {
	return lb.Element.Attr("for")
}

// -----------------------------------------------------------------------------

func (lb *label) SetFor(f string) {
	lb.Element.SetAttr("for", f)
}

// -----------------------------------------------------------------------------

func (lb *label) Caption() string {
	return lb.Text()
}

// -----------------------------------------------------------------------------

func (lb *label) SetCaption(caption string) {
	lb.Element.SetText((caption))
}

// -----------------------------------------------------------------------------

func (lb *label) Horizontal() bool {
	return lb.Ref().Contains("col-form-label")
}

// -----------------------------------------------------------------------------

func (lb *label) SetHorizontal(h bool) {
	lb.Ref().Remove("col-form-label", "form-label")
	if h {
		lb.Ref().Add("col-form-label")
	} else {
		lb.Ref().Add("form-label")
	}
}

// -----------------------------------------------------------------------------

func LabelOf(owner srx.TComponent) TLabel {
	ret := &label{
		Component: srx.NewComponent(owner, js.From(`<label class="form-label"></label>`)),
	}
	if owner != nil {
		owner.Add(ret)
	}
	return ret
}
