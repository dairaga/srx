// +build js,wasm

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFontWeight(t *testing.T) {
	data := map[FontWeight]string{
		FontWeightNone:    "",
		FontWeightBold:    "fw-bold",
		FontWeightBolder:  "fw-bolder",
		FontWeightNormal:  "fw-normal",
		FontWeightLight:   "fw-light",
		FontWeightLighter: "fw-lighter",
		FontWeight(1000):  "",
	}

	for k, v := range data {
		assert.Equal(t, k.String(), v)
	}
}
