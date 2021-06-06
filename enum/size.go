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
	Auto
	N50
	N100
	Small
	Medium
	Large
	Extra
)

const szName = "012345auto50100smmdlgxl"

var szIndex = [...]uint8{0, 1, 2, 3, 4, 5, 6, 10, 12, 15, 17, 19, 21, 23}

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
