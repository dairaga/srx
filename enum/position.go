// +build js,wasm

package enum

type Pos int

const (
	All Pos = iota
	Top
	Bottom
	Start
	End
	X
	Y
)

const (
	posName = `tbsexy`
)

var (
	posIndex = [...]uint8{0, 0, 1, 2, 3, 4, 5, 6, 10}
)

func (i Pos) String() string {
	if int(i) >= 0 && int(i) < len(posIndex)-1 {
		return posName[posIndex[i]:posIndex[i+1]]
	}
	return ""
}

func (i Pos) Margin(s Size) string {
	if s >= N0 && s <= Auto {
		return s.Style("m" + i.String())
	}
	return ""
}

func (i Pos) Padding(s Size) string {
	if s >= N0 && s <= Auto {
		return s.Style("p" + i.String())
	}
	return ""
}
