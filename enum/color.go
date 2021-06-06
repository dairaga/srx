// +build js,wasm

package enum

import "strings"

type Color int

const (
	Primary Color = iota
	Secondary
	Success
	Info
	Warning
	Danger
	Light
	Dark
	Link
)

const crName = `primarysecondarysuccessinfowarningdangerlightdarklink`

var crIndex = [...]uint{0, 7, 16, 23, 27, 34, 40, 45, 49, 53}

// -----------------------------------------------------------------------------

func (i Color) String() string {
	if int(i) >= 0 && int(i) < len(crIndex)-1 {
		return crName[crIndex[i]:crIndex[i+1]]
	}
	return Primary.String()
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
	*i = Primary
}
