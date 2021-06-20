// +build js,wasm

package enum

type Decoration int

const (
	DecorationReset Decoration = iota
	DecorationNone
	DecorationUnderline
	DecorationLineThrough
)

const dcName = `text-decoration-nonetext-decoration-underlinetext-decoration-line-through`

var dcIndex = [...]uint8{0, 0, 20, 45, 73}

func (i Decoration) IsDecoration() bool {
	return i >= DecorationNone && i <= DecorationLineThrough
}

func (i Decoration) String() string {
	if i.IsDecoration() {
		return dcName[dcIndex[i]:dcIndex[i+1]]
	}
	return ""
}

func (i Decoration) Apply(obj ObjRef) (ret bool) {
	if ret = i.IsDecoration(); ret {
		obj.Ref().Add(i.String())
	}
	return
}

func (i Decoration) Unapply(obj ObjRef) (ret bool) {
	if ret = i.IsDecoration(); ret {
		obj.Ref().Remove(i.String())
	}
	return
}
