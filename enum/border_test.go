// +build js,wasm

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBorder(t *testing.T) {
	data := map[BorderType]string{
		BorderNone:       "",
		Border:           "border",
		BorderTop:        "border-top",
		BorderEnd:        "border-end",
		BorderBottom:     "border-bottom",
		BorderStart:      "border-start",
		Border0:          "border-0",
		BorderTop0:       "border-top-0",
		BorderEnd0:       "border-end-0",
		BorderBottom0:    "border-bottom-0",
		BorderStart0:     "border-start-0",
		BorderType(1000): "",
	}

	for k, v := range data {
		assert.Equal(t, k.String(), v)
	}
}
