// +build js,wasm

package srx

import "github.com/dairaga/srx/js"

type (
	TAccordion interface {
		TComponent
		AddItem(capiton string, children ...TObject) TAccordionItem
		Items() []TAccordionItem
	}

	TAccordionItem interface {
		TComponent
		Caption() string
		SetCaption(s string)
		Append(children ...TObject)
		Prepend(children ...TObject)
		Actived() bool
		Active(active bool)
	}

	accordion struct {
		*component
		items []TAccordionItem
	}

	accordionItem struct {
		*component
	}
)

var _ TAccordion = &accordion{}
var _ TAccordionItem = &accordionItem{}

func (a *accordion) AddItem(capiton string, children ...TObject) TAccordionItem {
	item := newAccordionItem(a)
	item.SetCaption(capiton)
	item.Append(children...)
	a.Append(item)
	return item
}

func (a *accordion) Items() []TAccordionItem {
	return a.items
}

func (a *accordion) Append(children ...TObject) {
	size := len(children)
	if size <= 0 {
		return
	}
	items := make([]TAccordionItem, size)
	var ok bool
	for i := range children {
		items[i], ok = children[i].(TAccordionItem)
		if !ok || items[i] == nil {
			return
		}
	}
	a.items = append(a.items, items...)
	a.component.Append(children...)
}

func (a *accordion) Prepend(children ...TObject) {
	a.Append(children...)
}

func newAccordion(owner TComponent) *accordion {
	el := js.Create("div").Add("accordion")
	el.SetID(rand(10))

	ret := &accordion{
		component: newComponent(owner, el),
	}
	if owner != nil {
		owner.Add(ret)
	}
	return ret
}

func (a *accordion) SetID(id string) {
	a.Ref().QueryAll(".accordion-collapse").Foreach(func(_ int, el *js.Element) {
		el.SetAttr("data-bs-parent", "#"+id)
	})
	a.component.SetID(id)
}

func Accordion(owner TComponent) TAccordion {
	return newAccordion(owner)
}

func (a *accordionItem) Caption() string {
	return a.Ref().Query(".accordion-button").Text()
}

func (a *accordionItem) SetCaption(s string) {
	a.Ref().Query(".accordion-button").SetText(s)
}

func (a *accordionItem) Append(children ...TObject) {
	body := a.Query(".accordion-body")
	for i := range children {
		body.Append(children[i])
	}
}

func (a *accordionItem) Prepend(children ...TObject) {
	body := a.Query(".accordion-body")
	for i := range children {
		body.Prepend(children[i])
	}
}

func (a *accordionItem) Actived() bool {
	return !a.Query(".accordion-collapse").Contains("show")
}

func (a *accordionItem) Active(active bool) {
	if active {
		a.Query(".accordion-collapse").Add("show")
	} else {
		a.Query(".accordion-collapse").Remove("show")
	}

}

func newAccordionItem(owner TComponent) *accordionItem {

	parentID := owner.ID()
	headerID := rand(10)
	collapseID := rand(10)

	el := js.From(js.HTML(`<div class="accordion-item"><h2 class="accordion-header"><button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="" aria-expanded="true" aria-controls=""></button></h2><div id="" class="accordion-collapse collapse" aria-labelledby="" data-bs-parent=""><div class="accordion-body"></div></div></div>`))

	el.Query(".accordion-header").SetID(headerID)
	el.Query(".accordion-button").
		SetAttr("data-bs-target", "#"+collapseID).
		SetAttr("aria-controls", collapseID)

	el.Query(".accordion-collapse").
		SetAttr("aria-labelledby", headerID).
		SetAttr("data-bs-parent", "#"+parentID).
		SetID(collapseID)

	ret := &accordionItem{
		component: newComponent(owner, el),
	}

	if owner != nil {
		owner.Add(ret)
	}

	return ret
}

func AccordionItem(owner TComponent) TAccordionItem {
	return newAccordionItem(owner)
}
