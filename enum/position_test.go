// +build js,wasm

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPos(t *testing.T) {

	data := map[Size]string{
		N0:   "0",
		N1:   "1",
		N2:   "2",
		N3:   "3",
		N4:   "4",
		N5:   "5",
		Auto: "auto",
	}

	var typ Size
	for k, v := range data {
		assert.Equal(t, v, k.String())
		typ.SetString(v)
		assert.Equal(t, k, typ)
	}

	pos := map[Pos]string{
		Top:    "t",
		Bottom: "b",
		Start:  "s",
		End:    "e",
		X:      "x",
		Y:      "y",
	}

	for p, s := range pos {
		assert.Equal(t, s, p.Spacing())
		for k, v := range data {
			assert.Equal(t, "m"+s+"-"+v, p.Margin(k))
			assert.Equal(t, "p"+s+"-"+v, p.Padding(k))
		}
	}

}
