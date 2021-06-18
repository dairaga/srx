// +build js,wasm

package srx

import (
	"github.com/dairaga/srx/js"
)

type (
	TCheck interface {
		TBaseFormControl
		Checked() bool
		SetCheck(checked bool)
	}

	//TRadio = TCheck

	TBaseCheckGroup interface {
		TComponent
		Name() string
		SetName(name string)
		Inline() bool
		SetInline(inline bool)
		Children() []TCheck
		Child(index int) TCheck
		ChildLen() int
	}

	TCheckGroup interface {
		TBaseCheckGroup
		Value() []string
		SetValue(values ...string)
		AddCheck(value, caption string, checked bool) (TCheck, TLabel)
	}

	check struct {
		*input
	}

	basecheckgroup struct {
		*component
		name     string
		inline   bool
		children []TCheck
	}

	checkgroup struct {
		*basecheckgroup
	}
)

var _ TCheck = &check{}
var _ TBaseCheckGroup = &basecheckgroup{}
var _ TCheckGroup = &checkgroup{}

func (c *check) Checked() bool {
	return c.Prop("checked").Bool()
}

func (c *check) SetCheck(checked bool) {
	c.SetProp("checked", checked)
}

func newCheckControl(owner TComponent, tagName string) *check {
	input := newFormControl(owner, "input", "form-check-input")
	input.SetType(tagName)
	ret := &check{
		input: input,
	}
	return ret
}

func newCheck(owner TComponent) *check {
	return newCheckControl(owner, "checkbox")
}

func CheckOf(owner TComponent) TCheck {
	return newCheck(owner)
}

//func newRadio(owner TComponent) *check {
//	input := newFormControl(owner, "input", "form-check-input")
//	input.SetType("radio")
//	ret := &check{
//		input: input,
//	}
//	return ret
//}

//func RadioOf(owner TComponent) TRadio {
//	return newRadio(owner)
//}

func (g *basecheckgroup) Name() string {
	return g.name
}

func (g *basecheckgroup) SetName(name string) {
	g.name = name
}

func (g *basecheckgroup) Inline() bool {
	return g.inline
}

func (g *basecheckgroup) SetInline(inline bool) {
	g.inline = inline
}

func (g *basecheckgroup) Children() []TCheck {
	return g.children
}

func (g *basecheckgroup) Child(index int) TCheck {
	if index >= 0 && index < len(g.children) {
		return g.children[index]
	}
	return nil
}

func (g *basecheckgroup) ChildLen() int {
	return len(g.children)
}

func (g *checkgroup) Value() []string {
	values := []string{}
	g.Element.QueryAll(`input[type="checkbox"]`).Foreach(
		func(_ int, el *js.Element) {
			if el.Prop("checked").Bool() {
				values = append(values, el.Prop("value").String())
			}
		},
	)
	return values
}

func (g *checkgroup) SetValue(values ...string) {
	size := len(values)
	if size <= 0 {
		return
	}

	values = values[:size]

	g.Element.QueryAll(`input[type="checkbox"]`).Foreach(
		func(_ int, el *js.Element) {
			found := false
			for i := 0; i < size; i++ {
				if el.Prop("value").String() == values[i] {
					found = true
					break
				}
			}
			el.SetProp("checked", found)
		},
	)
}

func (g *checkgroup) AddCheck(value, caption string, checked bool) (TCheck, TLabel) {
	div := js.Create("div").Add("form-check")
	if g.inline {
		div.Add("form-check-inline")
	}
	check := newCheck(g)
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

func CheckGroup(owner TComponent) TCheckGroup {
	g := &checkgroup{
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
