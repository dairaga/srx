// +build js,wasm

package srx

type TNumber = TRange

// -----------------------------------------------------------------------------

func Number(owner TComponent) TNumber {
	ret := &rangeInput{
		input: newInput(owner),
	}
	ret.SetType("number")
	ret.SetMin(0)
	ret.SetMax(100)
	ret.SetStep(1)
	return ret
}
