// +build js,wasm

package srx

import (
	"testing"

	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {

	testCom := Component(js.From(js.HTML(`<div>TestMemory</div>`)))
	bindOwner(Root(), testCom)

	com := mem.lookup(testCom)
	tattoo := testCom.Tattoo()

	assert.NotEqual(t, "", tattoo)
	assert.Equal(t, tattoo, com.Tattoo())
	testCom.Release()
	assert.Nil(t, mem.table[tattoo])
}
