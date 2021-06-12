// +build js,wasm

package enum

type FlexMode int

const (
	FlexModeRow FlexMode = iota
	FlexModeColumn
)

const fmName = `flex-rowflex-column`

var fmIndex = [...]uint8{0, 8, 19}

func (i FlexMode) String() string {
	if int(i) >= 0 && int(i) < len(fmIndex)-1 {
		return fmName[fmIndex[i]:fmIndex[i+1]]
	}
	return FlexModeRow.String()
}
