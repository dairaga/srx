// +build js,wasm

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlexMode(t *testing.T) {

	data := map[FlexMode]string{
		FlexModeRow:    "flex-row",
		FlexModeColumn: "flex-column",
		FlexMode(1000): "",
	}

	for k, v := range data {
		assert.Equal(t, v, k.String())
	}

}
