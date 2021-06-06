// +build js,wasm

package enum

type Size int

const (
	None Size = iota
	Small
	Medium
	Large
	Extra
)

const szName = "smmdlgxl"

var szIndex = [...]uint8{0, 0, 2, 4, 6, 8}

// -----------------------------------------------------------------------------

func (i Size) String() string {
	if int(i) >= 0 && int(i) < len(szIndex)-1 {
		return szName[szIndex[i]:szIndex[i+1]]
	}
	return None.String()
}

// -----------------------------------------------------------------------------

func (i *Size) SetString(s string) {
	for k := 0; k < len(szIndex)-1; k++ {
		if s == szName[szIndex[k]:szIndex[k+1]] {
			*i = Size(k)
			return
		}
	}
	*i = None
}

// -----------------------------------------------------------------------------

func (i Size) Style(s string) string {
	if i == None {
		return ""
	}
	return s + "-" + i.String()
}
