// +build js,wasm

package enum

type FlexMode int

const (
	FlexModeNone FlexMode = iota
	FlexModeRow
	FlexModeColumn
)

const fmName = `flex-rowflex-column`

var fmIndex = [...]uint8{0, 0, 8, 19}

// -----------------------------------------------------------------------------

func (i FlexMode) String() string {
	if int(i) >= 0 && int(i) < len(fmIndex)-1 {
		return fmName[fmIndex[i]:fmIndex[i+1]]
	}
	return FlexModeNone.String()
}

// -----------------------------------------------------------------------------

func (i FlexMode) Apply(obj ObjRef) bool {
	if i == FlexModeNone {
		obj.Ref().Remove(
			"d-flex",
			FlexModeRow.String(),
			FlexModeColumn.String(),
		)
	} else {
		obj.Ref().Add("d-flex", i.String())
	}
	return true
}

// -----------------------------------------------------------------------------

func (i FlexMode) Unapply(obj ObjRef) bool {
	if i == FlexModeNone {
		return false
	}

	obj.Ref().Remove(i.String())
	return true
}
