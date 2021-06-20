// +build js,wasm

package enum

type FontWeight int

const (
	FontWeightNone FontWeight = iota
	FontWeightBold
	FontWeightBolder
	FontWeightNormal
	FontWeightLight
	FontWeightLighter
)

const fwName = `fw-boldfw-bolderfw-normalfw-lightfw-lighter`

var fwIndex = [...]uint8{0, 0, 7, 16, 25, 33, 43}

func (i FontWeight) IsFontWeight() bool {
	return i >= FontWeightBold && i <= FontWeightLighter
}

func (i FontWeight) String() string {
	if i.IsFontWeight() {
		return fwName[fwIndex[i]:fwIndex[i+1]]
	}
	return ""
}

func (i FontWeight) Apply(obj ObjRef) (ret bool) {
	if ret = i.IsFontWeight(); ret {
		obj.Ref().Add(i.String())
	}
	return
}

func (i FontWeight) Unapply(obj ObjRef) (ret bool) {
	if ret = i.IsFontWeight(); ret {
		obj.Ref().Remove(i.String())
	}
	return
}
