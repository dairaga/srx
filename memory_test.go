// +build js,wasm

package srx

import (
	"testing"

	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {

	testCom := ComponentOf(Root(), js.From(js.HTML(`<div>TestMemory</div>`)))
	Root().Add(testCom)

	com := mem.lookup(testCom)
	tattoo := testCom.Tattoo()

	assert.NotEqual(t, "", tattoo)
	assert.Equal(t, tattoo, com.Tattoo())
	testCom.Release()
	assert.Nil(t, mem.table[tattoo])
}
