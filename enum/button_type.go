// +build js,wasm

package enum

type ButtonType int

const (
	Button ButtonType = iota
	Submit
	Reset
)

const btnTypName = `buttonsubmitreset`

var btnTypIndex = [4]uint8{0, 6, 12, 17}

func (i ButtonType) String() string {
	if int(i) >= 0 && int(i) < len(btnTypIndex)-1 {
		return btnTypName[btnTypIndex[i]:btnTypIndex[i+1]]
	}
	return Button.String()
}

func (i *ButtonType) SetString(s string) {
	for k := 0; k < len(btnTypIndex)-1; k++ {
		if s == btnTypName[btnTypIndex[k]:btnTypIndex[k+1]] {
			*i = ButtonType(k)
			return
		}
	}
	*i = Button
}
