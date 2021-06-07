// +build js,wasm

package enum

import "github.com/dairaga/srx/js"

type RoundType int

const (
	RoundNone RoundType = iota
	RoundPill
	RoundCircle
)

const rdName = "rounded-pillrounded-circle"

var rdIndex = [...]uint8{0, 0, 12, 26}

// -----------------------------------------------------------------------------

func (i RoundType) String() string {
	if int(i) >= 0 && int(i) < len(rdIndex)-1 {
		return rdName[rdIndex[i]:rdIndex[i+1]]
	}
	return ""
}

// -----------------------------------------------------------------------------

func (i RoundType) Replace(el *js.Element, old RoundType) {
	if old == RoundPill || old == RoundCircle {
		el.Remove(old.String())
	}
	if i == RoundPill || i == RoundCircle {
		el.Add(i.String())
	}
}
