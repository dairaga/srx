// +build js,wasm

package enum

type RoundedType int

const (
	RoundedNone RoundedType = iota
	RoundedPill
	RoundedCircle
	RoundedStart
	RoundedEnd
	RoundedTop
	RoundedBottom
	Rounded
)

const rdName = "rounded-pillrounded-circlerounded-startrounded-endrounded-toprounded-bottomrounded"

var rdIndex = [...]uint8{0, 0, 12, 26, 39, 50, 61, 75, 82}

// -----------------------------------------------------------------------------

func (i RoundedType) String() string {
	if int(i) >= 0 && int(i) < len(rdIndex)-1 {
		return rdName[rdIndex[i]:rdIndex[i+1]]
	}
	return ""
}

// -----------------------------------------------------------------------------

//func (i RoundType) Replace(el *js.Element, old RoundType) {
//	if old == RoundedPill || old == RoundedCircle {
//		el.Remove(old.String())
//	}
//	if i == RoundedPill || i == RoundedCircle {
//		el.Add(i.String())
//	}
//}

// -----------------------------------------------------------------------------

func (i RoundedType) Apply(obj ObjRef) (ret bool) {
	if ret = (i >= RoundedPill && i <= Rounded); ret {
		obj.Ref().Add(i.String())
	}
	return
}

// -----------------------------------------------------------------------------

func (i RoundedType) Unapply(obj ObjRef) (ret bool) {
	if ret = (i >= RoundedPill && i <= Rounded); ret {
		obj.Ref().Remove(i.String())
	}
	return
}
