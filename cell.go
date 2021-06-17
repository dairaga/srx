// +build js,wasm

package srx

import (
	"github.com/dairaga/srx/js"
)

type (
	TCeller interface {
		Cells() []TCell
		Cell(index int) TCell
		CellLen() int
	}

	TCell interface {
		TObject
	}

	cell struct {
		*object
	}
)

var _ TCell = &cell{}

// -----------------------------------------------------------------------------

func Cell(children ...TObject) TCell {
	ret := &cell{
		object: newObject(js.Create("div")),
	}

	ret.Append(children...)
	return ret
}
