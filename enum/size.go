// +build js,wasm

package enum

type Size int

const (
	N0 Size = iota
	N1
	N2
	N3
	N4
	N5
	N6
	N7
	N8
	N9
	N10
	N11
	N12
	Auto
	N50
	N100
	Small
	Medium
	Large
	Extra
)

const szName = "0123456789101112auto50100smmdlgxl"

var szIndex = [...]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14, 16, 20, 22, 25, 27, 29, 31, 33}

// -----------------------------------------------------------------------------

func (i Size) String() string {
	if int(i) >= 0 && int(i) < len(szIndex)-1 {
		return szName[szIndex[i]:szIndex[i+1]]
	}
	return N0.String()
}

// -----------------------------------------------------------------------------

func (i *Size) SetString(s string) {
	for k := 0; k < len(szIndex)-1; k++ {
		if s == szName[szIndex[k]:szIndex[k+1]] {
			*i = Size(k)
			return
		}
	}
	*i = N0
}

// -----------------------------------------------------------------------------

func (i Size) Style(s string) string {
	return s + "-" + i.String()
}

// -----------------------------------------------------------------------------

func (i Size) Col() string {
	if i >= N1 && i <= Auto {
		return "col-" + i.String()
	}
	return "col"
}
