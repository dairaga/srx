// +build js,wasm

package enum

type SpinnerType int

const (
	SpinnerBorder SpinnerType = iota
	SpinnerGrow
)

const spName = `spinner-borderspinner-grow`

var spIndex = [...]uint8{0, 14, 26}

func (i SpinnerType) String() string {
	if int(i) >= 0 && int(i) < len(spIndex)-1 {
		return spName[spIndex[i]:spIndex[i+1]]
	}
	return SpinnerBorder.String()
}
