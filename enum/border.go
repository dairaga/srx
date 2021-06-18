// +build js,wasm

package enum

type BorderType int

const (
	BorderNone BorderType = iota
	Border
	BorderTop
	BorderEnd
	BorderBottom
	BorderStart
	Border0
	BorderTop0
	BorderEnd0
	BorderBottom0
	BorderStart0
)

const brName = `borderborder-topborder-endborder-bottomborder-startborder-0border-top-0border-end-0border-bottom-0border-start-0`

var brIndex = [...]uint8{0, 0, 6, 16, 26, 39, 51, 59, 71, 83, 98, 112}

func (i BorderType) String() string {
	if int(i) >= 0 && int(i) < len(brIndex)-1 {
		return brName[brIndex[i]:brIndex[i+1]]
	}
	return ""
}

func (i BorderType) Apply(obj ObjRef) (ret bool) {
	if ret = (i >= Border && i <= BorderStart0); ret {
		obj.Ref().Add(i.String())
	}
	return
}

func (i BorderType) Unapply(obj ObjRef) (ret bool) {
	if ret = (i >= Border && i <= BorderStart0); ret {
		obj.Ref().Remove(i.String())
	}
	return
}
