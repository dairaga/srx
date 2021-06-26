// +build js,wasm

package srx

import (
	"fmt"

	"github.com/dairaga/srx/js"
)

type (
	TLink interface {
		TObject
		Caption() string
		SetCaption(caption string)

		HRef() string
		SetHRef(href string)
	}

	link struct {
		*object
	}
)

var _ TLink = &link{}

func (a *link) Caption() string {
	return a.Text()
}

func (a *link) SetCaption(caption string) {
	a.SetText(caption)
}

func (a *link) HRef() string {
	return a.Attr("href")
}

func (a *link) SetHRef(href string) {
	a.SetAttr("href", href)
}

func newLink(caption string, href string) *link {
	ret := &link{
		object: newObject(js.Create("a")),
	}
	ret.SetCaption(caption)
	if href != "" {
		ret.SetHRef(href)
	}
	return ret
}

func Link(href string, a ...interface{}) TLink {
	caption := fmt.Sprint(a...)
	return newLink(caption, href)
}

func Linkf(href, format string, a ...interface{}) TLink {
	caption := fmt.Sprintf(format, a...)
	return newLink(caption, href)
}
