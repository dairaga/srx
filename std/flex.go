// build js,wasm

package std

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
)

type (
	TFlexPanel interface {
		srx.TComponent
		Mode() enum.FlexMode
		SetMode(m enum.FlexMode)
		AlignHorizontal(al enum.Align)
		AlignVertical(al enum.Align)
	}

	flex struct {
		*srx.Component
		m enum.FlexMode
		h enum.Align
		v enum.Align
	}
)

var _ TFlexPanel = &flex{}

// -----------------------------------------------------------------------------

func (f *flex) Mode() enum.FlexMode {
	return f.m
}

// -----------------------------------------------------------------------------

func (f *flex) SetMode(m enum.FlexMode) {
	f.Element.Remove(enum.FlexModeRow.String(), enum.FlexModeColumn.String())

	if m != enum.FlexModeColumn && m != enum.FlexModeRow {
		m = enum.FlexModeRow
	}

	f.Element.Add(m.String())
	f.m = m
}

// -----------------------------------------------------------------------------

func (f *flex) AlignHorizontal(al enum.Align) {
	if al.IsHorizontal() {
		if f.h != enum.AlignNone {
			f.Element.Remove(f.h.Horizontal())
		}
		f.Element.Add(al.Horizontal())
		f.h = al
	}
}

// -----------------------------------------------------------------------------

func (f *flex) AlignVertical(al enum.Align) {
	if al.IsVertical() {
		if f.v != enum.AlignNone {
			f.Element.Remove(f.v.Vertical())
		}
		f.Element.Add(al.Vertical())
		f.v = al
	}
}

// -----------------------------------------------------------------------------

func FlexPanelOf(owner srx.TComponent) TFlexPanel {
	el := js.Create("div").Add("d-flex", "flex-row")
	ret := &flex{
		Component: srx.NewComponent(owner, el),
		m:         enum.FlexModeRow,
		h:         enum.AlignNone,
		v:         enum.AlignNone,
	}
	if owner != nil {
		owner.Add(ret)
	}

	return ret
}
