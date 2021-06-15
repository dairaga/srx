// +build js,wasm

package enum

import "strings"

type Color int

const (
	None Color = iota
	Link
	Primary
	Secondary
	Success
	Info
	Warning
	Danger
	Light
	Dark
	Block50
	White50
	Muted
	Body
	White
	Transparent
)

const crName = `linkprimarysecondarysuccessinfowarningdangerlightdarkblack-50white-50mutedbodywhitetransparent`

var crIndex = [...]uint{0, 0, 4, 11, 20, 27, 31, 38, 44, 49, 53, 61, 69, 74, 78, 83, 94}

// -----------------------------------------------------------------------------

func (i Color) String() string {
	if int(i) >= 0 && int(i) < len(crIndex)-1 {
		return crName[crIndex[i]:crIndex[i+1]]
	}
	return None.String()
}

// -----------------------------------------------------------------------------

func (i Color) Style(s ...string) string {
	return strings.Join(append(s, i.String()), "-")
}

// -----------------------------------------------------------------------------

func (i *Color) SetString(s string) {
	for k := 0; k < len(crIndex)-1; k++ {
		if s == crName[crIndex[k]:crIndex[k+1]] {
			*i = Color(k)
			return
		}
	}
	*i = None
}

// -----------------------------------------------------------------------------

func (i Color) TextColor() string {
	if i == None {
		return Body.Style("text")
	}
	return i.Style("text")
}

// -----------------------------------------------------------------------------

func (i Color) BackgroupColor() string {
	if i == None {
		return Transparent.BackgroupColor()
	}
	return i.Style("bg")
}
