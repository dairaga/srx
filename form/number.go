// +build js,wasm

package form

import "github.com/dairaga/srx"

type TNumber = TRange

// -----------------------------------------------------------------------------

func NumberOf(owner srx.TComponent) TNumber {
	ret := &rangeInput{
		input: newInput(owner),
	}
	ret.SetType("number")
	ret.SetMin(0)
	ret.SetMax(100)
	ret.SetStep(1)
	return ret
}
