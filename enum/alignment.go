// +build js,wasm

package enum

type Align int

const (
	AlignStart Align = iota
	AlignEnd
	AlignCenter
	AlignBetween
	AlignAround
	AlignEvenly
	AlignBaseline
	AlignStretch
)

const alName = `startendcenterbetweenaroundevenlybaselinestretch`

var alIndex = [...]uint8{0, 5, 8, 14, 21, 27, 33, 41, 48}

// -----------------------------------------------------------------------------

func (i Align) String() string {
	if int(i) >= 0 && int(i) < len(alIndex)-1 {
		return alName[alIndex[i]:alIndex[i+1]]
	}
	return AlignStart.String()
}

// -----------------------------------------------------------------------------

func (i Align) Horizontal() string {
	if i >= AlignStart && i <= AlignEvenly {
		return "justify-content-" + i.String()
	}
	return AlignStart.Horizontal()
}

// -----------------------------------------------------------------------------

func (i Align) Vertical() string {
	if (i >= AlignStart && i <= AlignCenter) || (i >= AlignBaseline && i <= AlignStretch) {
		return "align-items-" + i.String()
	}
	return AlignStart.Vertical()
}
