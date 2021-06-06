// +build js,wasm

package enum

type Pos int

const (
	Top Pos = iota
	Bottom
	Start
	End
	X
	Y
)

const (
	spName = `tbsexy`
)

var (
	spIndex = [...]uint8{0, 1, 2, 3, 4, 5, 6, 10}
)

// -----------------------------------------------------------------------------

func (i Pos) Spacing() string {
	if int(i) >= 0 && int(i) < len(spIndex)-1 {
		return spName[spIndex[i]:spIndex[i+1]]
	}
	return ""
}

// -----------------------------------------------------------------------------

func (i Pos) Margin(s Size) string {
	if s >= N0 && s <= Auto {
		return s.Style("m" + i.Spacing())
	}
	return ""
}

// -----------------------------------------------------------------------------

func (i Pos) Padding(s Size) string {
	if s >= N0 && s <= Auto {
		return s.Style("p" + i.Spacing())
	}
	return ""
}
