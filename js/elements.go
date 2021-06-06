// +build js,wasm

package js

type Elements []*Element

func (els Elements) Foreach(fn func(int, *Element)) {
	for i, el := range els {
		fn(i, el)
	}
}
