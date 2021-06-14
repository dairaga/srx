// +build js,wasm

package el

import (
	"github.com/dairaga/srx"
	"github.com/dairaga/srx/js"
)

type (
	TCeller interface {
		Cells() []TCell
		Cell(index int) TCell
		CellLen() int
	}

	TCell interface {
		srx.TObject
	}

	cell struct {
		*srx.Object
	}
)

var _ TCell = &cell{}

// -----------------------------------------------------------------------------

func Cell(children ...srx.TObject) TCell {
	ret := &cell{
		Object: srx.NewObject(js.Create("div")),
	}

	ret.Append(children...)
	return ret
}
