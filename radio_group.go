// +build js,wasm

package srx

import (
	"github.com/dairaga/srx/js"
)

type (
	TRadio      = TCheck
	TRadioGroup interface {
		TBaseCheckGroup
		Value() string
		SetValue(value string)
		AddRadio(value, caption string, checked bool) (TRadio, TLabel)
	}

	radiogroup struct {
		*basecheckgroup
	}
)

var _ TRadioGroup = &radiogroup{}

func newRadio(owner TComponent) *check {
	return newCheckControl(owner, "radio")
}

func RadioOf(owner TComponent) TRadio {
	return newRadio(owner)
}

func (g *radiogroup) Value() (ret string) {
	g.Element.QueryAll(`input[type="radio"]`).Foreach(
		func(_ int, el *js.Element) {
			if el.Prop("checked").Bool() {
				ret = el.Prop("value").String()
				return
			}
		},
	)
	return
}

func (g *radiogroup) SetValue(value string) {
	g.Element.QueryAll(`input[type="radio"]`).Foreach(
		func(_ int, el *js.Element) {
			el.SetProp("checked", value == el.Prop("value").String())
		},
	)
}

func (g *radiogroup) AddRadio(value, caption string, checked bool) (TRadio, TLabel) {
	div := js.Create("div").Add("form-check")
	if g.inline {
		div.Add("form-check-inline")
	}
	check := newRadio(g)
	check.SetName(g.name)
	check.SetID(g.name + "_" + value)
	check.SetCheck(checked)
	check.SetValue(value)
	g.children = append(g.children, check)

	div.Append(check)

	lb := newLabel(g, `form-check-label`)
	lb.SetCaption(caption)
	lb.SetFor(check.ID())
	div.Append(lb)
	g.Element.Append(div)
	return check, lb
}

func RadioGroup(owner TComponent) TRadioGroup {
	g := &radiogroup{
		basecheckgroup: &basecheckgroup{
			component: newComponent(owner, js.Create("div")),
			inline:    false,
			name:      "",
		},
	}

	if owner != nil {
		owner.Add(g)
	}
	return g
}
