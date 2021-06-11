// +build js,wasm

package enum

type Align int

const (
	AlignNone Align = iota
	AlignStart
	AlignEnd
	AlignCenter
	AlignBetween
	AlignAround
	AlignEvenly
	AlignBaseline
	AlignStretch
)

const alName = `startendcenterbetweenaroundevenlybaselinestretch`

var alIndex = [...]uint8{0, 0, 5, 8, 14, 21, 27, 33, 41, 48}

// -----------------------------------------------------------------------------

func (i Align) String() string {
	if int(i) >= 0 && int(i) < len(alIndex)-1 {
		return alName[alIndex[i]:alIndex[i+1]]
	}
	return ""
}

// -----------------------------------------------------------------------------

func (i Align) IsHorizontal() bool {
	return i >= AlignStart && i <= AlignEvenly
}

// -----------------------------------------------------------------------------

func (i Align) Horizontal() string {
	if i.IsHorizontal() {
		return "justify-content-" + i.String()
	}
	return ""
}

// -----------------------------------------------------------------------------

func (i Align) IsVertical() bool {
	return (i >= AlignStart && i <= AlignCenter) || (i >= AlignBaseline && i <= AlignStretch)
}

// -----------------------------------------------------------------------------

func (i Align) Vertical() string {
	if i.IsVertical() {
		return "align-items-" + i.String()
	}
	return ""
}
