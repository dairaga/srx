// +build js,wasm

package js_test

import (
	"fmt"
	"testing"

	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestHTML(t *testing.T) {
	caption := "hello"
	tmpl := js.HTML(fmt.Sprintf(`<div>%s</div>`, caption))
	elem := tmpl.Element()

	assert.Equal(t, caption, elem.Text())
}
