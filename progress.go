// +build js,wasm

package srx

import (
	"strconv"

	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TProgress interface {
		TComponent
		AddBar(bar TProgressBar)
		Bars() []TProgressBar

		Striped() bool
		Stripe(s bool)

		Animated() bool
		Animate(a bool)
	}

	TProgressBar interface {
		TComponent
		Value() int
		SetValue(v int)

		Caption() string
		SetCaption(caption string)

		Striped() bool
		Stripe(s bool)

		Animated() bool
		Animate(a bool)
	}

	progress struct {
		*component
		bars     []TProgressBar
		striped  bool
		animated bool
	}

	progressbar struct {
		*component
		value int
	}
)

var _ TProgress = &progress{}
var _ TProgressBar = &progressbar{}

func (p *progress) AddBar(bar TProgressBar) {
	bar.Stripe(p.striped)
	bar.Animate(p.animated)
	p.bars = append(p.bars, bar)
	p.Append(bar)
}

func (p *progress) Bars() []TProgressBar {
	return p.bars
}

func (p *progress) Striped() bool {
	return p.striped
}

func (p *progress) Stripe(s bool) {
	p.striped = s
	for i := range p.bars {
		p.bars[i].Stripe(s)
	}
}

func (p *progress) Animated() bool {
	return p.animated
}

func (p *progress) Animate(a bool) {
	p.animated = a
	for i := range p.bars {
		p.bars[i].Animate(a)
	}
}

func (p *progressbar) Value() int {
	return p.value
}

func (p *progressbar) SetValue(v int) {
	if v < 0 {
		v = 0
	}
	if v > 100 {
		v = 100
	}
	p.value = v
	p.Ref().SetStyle("width", strconv.Itoa(p.value)+"%")
}

func (p *progressbar) Striped() bool {
	return p.Ref().Contains("progress-bar-striped")
}

func (p *progressbar) Stripe(s bool) {
	if s {
		p.Ref().Add("progress-bar-striped")
	} else {
		p.Ref().Remove("progress-bar-striped")
	}
}

func (p *progressbar) Animated() bool {
	return p.Ref().Contains("progress-bar-animated")
}

func (p *progressbar) Animate(a bool) {
	if a {
		p.Ref().Add("progress-bar-animated")
	} else {
		p.Ref().Remove("progress-bar-animated")
	}
}

func (p *progressbar) Caption() string {
	return p.Ref().Text()
}

func (p *progressbar) SetCaption(c string) {
	p.Ref().SetText(c)
}

func (p *progressbar) Color() enum.Color {
	return p.Background()
}

func (p *progressbar) SetColor(c enum.Color) {
	p.SetBackground(c)
}

func newProgressBar(owner TComponent) *progressbar {

	el := js.Create("div").Add("progress-bar").SetAttr("role", "progressbar")
	ret := &progressbar{
		component: newComponent(owner, el),
		value:     0,
	}
	ret.SetColor(enum.Primary)
	if owner != nil {
		owner.Add(ret)
	}

	return ret
}

func ProgressBar(owner TComponent) TProgressBar {
	return newProgressBar(owner)
}

func newProgress(owner TComponent) *progress {
	ret := &progress{
		component: newComponent(owner, js.Create("div").Add("progress")),
	}

	if owner != nil {
		owner.Add(ret)
	}
	return ret
}

func Progress(owner TComponent) TProgress {
	return newProgress(owner)
}
