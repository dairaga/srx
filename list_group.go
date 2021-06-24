// +build js,wasm

package srx

import "github.com/dairaga/srx/js"

type (
	TListGroup interface {
		TComponent
		Items() []TListItem
		AddItem(children ...TObject) TListItem
		Active(idx int)
		Actived() int

		Flushed() bool
		Flush(f bool)

		Actionabled() bool
		Actionable(a bool)
	}

	TOrderedListGroup = TListGroup

	TListItem interface {
		TComponent
		//Caption() string
		//SetCaption(caption string)
	}

	listgroup struct {
		*component
		items       []TListItem
		actived     int
		actionabled bool
	}

	listitem struct {
		*component
	}
)

var _ TListGroup = &listgroup{}
var _ TListItem = &listitem{}

func (g *listgroup) Items() []TListItem {
	return g.items
}

func (g *listgroup) AddItem(children ...TObject) TListItem {
	item := newListItem(g)
	item.Append(children...)
	g.items = append(g.items, item)
	g.Append(item)
	return item
}

func (g *listgroup) Active(idx int) {
	if idx >= 0 && idx < len(g.items) {
		if g.actived >= 0 {
			g.items[g.actived].Ref().Remove("active")
		}

		g.items[idx].Ref().Add("active")
		g.actived = idx
	}
}

func (g *listgroup) Actived() int {
	return g.actived
}

func (g *listgroup) Flushed() bool {
	return g.Ref().Contains("list-group-flush")
}

func (g *listgroup) Flush(f bool) {
	if f {
		g.Ref().Add("list-group-flush")
	} else {
		g.Ref().Remove("list-group-flush")
	}
}

func (g *listgroup) Actionabled() bool {
	return g.actionabled
}

func (g *listgroup) Actionable(a bool) {

	for i := range g.items {
		if a {
			g.items[i].Ref().Add("list-group-item-action")
		} else {
			g.items[i].Ref().Remove("list-group-item-action")
		}
	}
	g.actionabled = a
}

func (item *listitem) Disable() {
	item.Ref().Add("disabled")
	item.component.Disable()
}

func (item *listitem) Enable() {
	item.Ref().Remove("disabled")
	item.component.Enable()
}

//func (item *listitem) Caption() string {
//	return item.Ref().Text()
//}
//
//func (item *listitem) SetCaption(caption string) {
//	item.Ref().SetText(caption)
//}

func newListGroup(owner TComponent, ordered bool) (ret *listgroup) {
	if ordered {
		ret = &listgroup{
			component: newComponent(js.Create("ol").Add("list-group", "list-group-numbered")),
		}
	} else {
		ret = &listgroup{
			component: newComponent(js.Create("ul").Add("list-group")),
		}
	}

	bindOwner(owner, ret)
	return
}

func ListGroup(owner TComponent) TListGroup {
	return newListGroup(owner, false)
}

func OrderedListGroup(owner TComponent) TOrderedListGroup {
	return newListGroup(owner, true)
}

func newListItem(owner TComponent) *listitem {
	ret := &listitem{
		component: newComponent(js.Create("li").Add("list-group-item")),
	}

	bindOwner(owner, ret)
	return ret
}

func ListItem(owner TComponent) TListItem {
	return newListItem(owner)
}
