// +build js,wasm

package el

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TCell interface {
		srx.TObject
	}

	cell struct {
		*srx.Object
	}
)

var _ TCell = &cell{}

// -----------------------------------------------------------------------------

func Cell() TCell {
	return &caption{
		Object: srx.NewObject(js.Create("div")),
	}
}
