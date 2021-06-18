// +build js,wasm

package enum

import (
	"strings"
)

type Color int

const (
	None Color = iota
	Link
	Transparent
	Primary
	Secondary
	Success
	Info
	Warning
	Danger
	Light
	Dark
	White
	Body
	Black50
	White50
	Muted
)

const crName = `linktransparentprimarysecondarysuccessinfowarningdangerlightdarkwhitebodyblack-50white-50muted`

var crIndex = [...]uint{0, 0, 4, 15, 22, 31, 38, 42, 49, 55, 60, 64, 69, 73, 81, 89, 94}

func (i Color) String() string {
	if int(i) >= 0 && int(i) < len(crIndex)-1 {
		return crName[crIndex[i]:crIndex[i+1]]
	}
	return None.String()
}

func (i Color) IsTextColor() bool {
	return i >= Primary && i <= Muted
}

func (i Color) IsBGColor() bool {
	return i >= Transparent && i <= White
}

func (i Color) IsTheme() bool {
	return i >= Primary && i <= Dark
}

func (i Color) Style(s ...string) string {
	return strings.Join(append(s, i.String()), "-")
}

func (i Color) Apply(obj ObjRef, s ...string) {
	obj.Ref().Add(i.Style(s...))
}

func (i Color) Unapply(obj ObjRef, s ...string) {
	obj.Ref().Remove(i.Style(s...))
}

func (i *Color) SetString(s string) {
	for k := 0; k < len(crIndex)-1; k++ {
		if s == crName[crIndex[k]:crIndex[k+1]] {
			*i = Color(k)
			return
		}
	}
	*i = None
}

func (i Color) ApplyText(obj ObjRef) (ret bool) {
	if ret = i.IsTextColor(); ret {
		i.Apply(obj, "text")
	}
	return
}

func (i Color) UnapplyText(obj ObjRef) (ret bool) {
	if ret = i.IsTextColor(); ret {
		i.Unapply(obj, "text")
	}
	return
}

func (i Color) ApplyBackground(obj ObjRef) (ret bool) {
	if ret = i.IsBGColor(); ret {
		i.Apply(obj, "bg")
	}
	return
}

func (i Color) UnapplyBackground(obj ObjRef) (ret bool) {
	if ret = i.IsBGColor(); ret {
		i.Unapply(obj, "bg")
	}
	return
}

func (i Color) ApplyButton(obj ObjRef) (ret bool) {
	if ret = (i.IsTheme() || i == Link); ret {
		i.Apply(obj, "btn")
	}
	return
}

func (i Color) UnapplyButton(obj ObjRef) (ret bool) {
	if ret = (i.IsTheme() || i == Link); ret {
		i.Unapply(obj, "btn")
	}
	return
}

func (i Color) ApplyOutlineButton(obj ObjRef) (ret bool) {
	if ret = (i.IsTheme() || i == Link); ret {
		i.Apply(obj, "btn", "outline")
	}
	return
}

func (i Color) UnapplyOutlineButton(obj ObjRef) (ret bool) {
	if ret = (i.IsTheme() || i == Link); ret {
		i.Unapply(obj, "btn", "outline")
	}
	return
}

func (i Color) ApplyBorder(obj ObjRef) (ret bool) {
	if ret = (i >= Primary && i <= White); ret {
		obj.Ref().Add("border-" + i.String())
	}
	return
}

func (i Color) UnapplyBorder(obj ObjRef) (ret bool) {
	if ret = (i >= Primary && i <= White); ret {
		obj.Ref().Remove("border-" + i.String())
	}
	return
}
