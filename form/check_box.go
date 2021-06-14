// +build js,wasm

package form

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TCheck interface {
		TBaseFormControl
		Checked() bool
		SetCheck(checked bool)
	}

	TRadio = TCheck

	TCheckGroup interface {
		srx.TComponent
		Name() string
		SetName(name string)
		AddCheck(value, caption string, checked bool) (TCheck, TLabel)
		AddRadio(value, caption string, checked bool) (TRadio, TLabel)
		Inline() bool
		SetInline(inline bool)
	}

	check struct {
		*input
	}

	checkgroup struct {
		*srx.Component
		name   string
		inline bool
	}
)

var _ TCheck = &check{}
var _ TCheckGroup = &checkgroup{}

// -----------------------------------------------------------------------------

func (c *check) Checked() bool {
	return c.Element.HasAttr("checked")
}

// -----------------------------------------------------------------------------

func (c *check) SetCheck(checked bool) {
	if checked {
		c.Element.SetAttr("checked", "true")
	} else {
		c.Element.RemoveAttr("checked")
	}
}

// -----------------------------------------------------------------------------

func newCheck(owner srx.TComponent) *check {
	input := newFormControl(owner, "input", "form-check-input")
	input.SetType("checkbox")
	ret := &check{
		input: input,
	}
	return ret
}

// -----------------------------------------------------------------------------

func CheckOf(owner srx.TComponent) *check {
	return newCheck(owner)
}

// -----------------------------------------------------------------------------

func newRadio(owner srx.TComponent) *check {
	input := newFormControl(owner, "input", "form-check-input")
	input.SetType("radio")
	ret := &check{
		input: input,
	}
	return ret
}

// -----------------------------------------------------------------------------

func RadioOf(owner srx.TComponent) TRadio {
	return newRadio(owner)
}

// -----------------------------------------------------------------------------

func (g *checkgroup) Name() string {
	return g.name
}

// -----------------------------------------------------------------------------

func (g *checkgroup) SetName(name string) {
	g.name = name
}

// -----------------------------------------------------------------------------

func (g *checkgroup) Inline() bool {
	return g.inline
}

// -----------------------------------------------------------------------------

func (g *checkgroup) SetInline(inline bool) {
	g.inline = inline
}

// -----------------------------------------------------------------------------

func (g *checkgroup) AddCheck(value, caption string, checked bool) (TCheck, TLabel) {
	div := js.Create("div").Add("form-check")
	if g.inline {
		div.Add("form-check-inline")
	}
	check := newCheck(g)
	check.SetName(g.name)
	check.SetID(g.name + "_" + value)
	check.SetCheck(checked)

	div.Append(check)

	lb := newLabel(g, `form-check-label`)
	lb.SetCaption(caption)
	lb.SetFor(check.ID())
	div.Append(lb)
	g.Element.Append(div)
	return check, lb
}

// -----------------------------------------------------------------------------

func (g *checkgroup) AddRadio(value, caption string, checked bool) (TRadio, TLabel) {
	div := js.Create("div").Add("form-check")
	if g.inline {
		div.Add("form-check-inline")
	}
	check := newRadio(g)
	check.SetName(g.name)
	check.SetID(g.name + "_" + value)
	check.SetCheck(checked)

	div.Append(check)

	lb := newLabel(g, `form-check-label`)
	lb.SetCaption(caption)
	lb.SetFor(check.ID())
	div.Append(lb)
	g.Element.Append(div)
	return check, lb
}

// -----------------------------------------------------------------------------

func CheckGroupOf(owner srx.TComponent) TCheckGroup {
	g := &checkgroup{
		Component: srx.NewComponent(owner, js.Create("div")),
		inline:    false,
		name:      "",
	}
	if owner != nil {
		owner.Add(g)
	}
	return g
}
