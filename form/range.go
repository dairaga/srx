// +build js,wasm

package form

import (
	"strconv"

	"github.com/dairaga/srx"
)

type (
	TRange interface {
		TBaseFormControl
		Max() int
		SetMax(v int)

		Min() int
		SetMin(v int)

		Step() int
		SetStep(v int)
	}

	rangeInput struct {
		*input
	}
)

var _ TRange = &rangeInput{}

// -----------------------------------------------------------------------------

func (r *rangeInput) rangeValue(a string) int {
	v, err := strconv.Atoi(r.Attr(a))
	if err != nil {
		return 0
	}

	return v
}

// -----------------------------------------------------------------------------

func (r *rangeInput) setRangeValue(a string, v int) {
	r.SetAttr(a, strconv.Itoa(v))
}

// -----------------------------------------------------------------------------

func (r *rangeInput) Max() int {
	return r.rangeValue("max")
}

// -----------------------------------------------------------------------------

func (r *rangeInput) SetMax(v int) {
	r.setRangeValue("max", v)
}

// -----------------------------------------------------------------------------

func (r *rangeInput) Min() int {
	return r.rangeValue("min")
}

// -----------------------------------------------------------------------------

func (r *rangeInput) SetMin(v int) {
	r.setRangeValue("min", v)
}

// -----------------------------------------------------------------------------

func (r *rangeInput) Step() int {
	return r.rangeValue("step")
}

// -----------------------------------------------------------------------------

func (r *rangeInput) SetStep(v int) {
	r.setRangeValue("step", v)
}

// -----------------------------------------------------------------------------

func RangeOf(owner srx.TComponent) TRange {
	ret := &rangeInput{
		input: newFormControl(owner, "input", "form-range"),
	}
	ret.SetType("range")
	ret.SetMin(0)
	ret.SetMax(100)
	ret.SetStep(1)
	ret.SetValue("0")
	return ret
}
