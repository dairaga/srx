// +build js,wasm

package enum

import "github.com/dairaga/srx/js"

type ObjRef interface {
	Ref() *js.Element
}
