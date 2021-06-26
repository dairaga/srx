// +build js,wasm

package srx

import (
	"fmt"

	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TDropdown interface {
		TComponent
		Caption() string
		SetCaption(caption string)
		Color() enum.Color
		SetColor(c enum.Color)
	}

	TDropdownMenu interface {
		TComponent
		AddItem(item TDropdownItem)
	}

	TDropdownItem        = TObject
	TDropdownItemHeader  = TObject
	TDropdownItemDivider = TObject
	TDropdownItemLink    = TLink
	TDropdownItemButton  = TButton

	dropdown struct {
		*component
		title *button
	}

	dropdownMenu struct {
		*component
		items []TDropdownItem
	}
)

var _ TDropdown = &dropdown{}
var _ TDropdownMenu = &dropdownMenu{}

func (d *dropdown) Caption() string {
	return d.title.Caption()
}

func (d *dropdown) SetCaption(caption string) {
	d.title.SetCaption(caption)
}

func (d *dropdown) Color() enum.Color {
	return d.title.Color()
}

func (d *dropdown) SetColor(c enum.Color) {
	d.title.SetColor(c)
}

func newDropdown(owner TComponent) *dropdown {
	el := js.Create("div").Add("dropdown")

	ret := &dropdown{
		component: newComponent(el),
	}
	btn := newButton(ret)
	btn.Ref().
		Add("dropdown-toggle").
		SetAttr("data-bs-toggle", "dropdown").
		SetAttr("aria-expanded", "false")

	ret.title = btn
	el.Append(btn)

	bindOwner(owner, ret)
	return ret
}

func Dropdown(owner TComponent) TDropdown {
	return newDropdown(owner)
}

func (d *dropdownMenu) AddItem(item TDropdownItem) {
	d.items = append(d.items, item)
	el := newObject(js.Create("li"))
	el.Append(item)
	d.Append(el)
}

func newDropdownMenu(owner TComponent) *dropdownMenu {
	el := js.Create("li").Add("dropdown-menu")
	ret := &dropdownMenu{
		component: newComponent(el),
		items:     nil,
	}

	bindOwner(owner, ret)
	return ret
}

func DropdownMenu(owner TComponent) TDropdownMenu {
	return newDropdownMenu(owner)
}

func newDropdownItemHeader(caption string) *object {
	el := js.Create("h6").Add("dropdown-header")
	el.SetText(caption)
	ret := newObject(el)
	return ret
}

func DropdownItemHeader(a ...interface{}) TDropdownItemHeader {
	return newDropdownItemHeader(fmt.Sprint(a...))
}

func DropdownItemHeaderf(format string, a ...interface{}) TDropdownItemHeader {
	return newDropdownItemHeader(fmt.Sprintf(format, a...))
}

func newDropdownItemDivider() *object {
	el := js.Create("hr").Add("dropdown-divider")
	return newObject(el)
}

func DropdownItemDivider() TDropdownItemDivider {
	return newDropdownItemDivider()
}

func newDropdownItemLink(caption string, href string) *link {
	ret := newLink(caption, href)
	ret.Ref().Add("dropdown-item")
	return ret
}

func DropdownItemLink(href string, a ...interface{}) TDropdownItemLink {
	caption := fmt.Sprint(a...)
	return newDropdownItemLink(caption, href)
}

func DropdownItemLinkf(href, format string, a ...interface{}) TDropdownItemLink {
	caption := fmt.Sprintf(format, a...)
	return newDropdownItemLink(caption, href)
}

func newDropdownItemButton(owner TComponent) *button {
	ret := newButton(owner)
	ret.Ref().Add("dropdown-item").Remove("btn")
	return ret
}

func DropdownItemButton(owner TComponent) TButton {
	return newDropdownItemButton(owner)
}
