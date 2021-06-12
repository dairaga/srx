// +build js,wasm

package enum

import "fmt"

type ItemsPerRow int

const (
	N1PerRow  ItemsPerRow = 12
	N2PerRow  ItemsPerRow = 6
	N3PerRow  ItemsPerRow = 4
	N4PerRow  ItemsPerRow = 3
	N6PerRow  ItemsPerRow = 2
	N12PerRow ItemsPerRow = 1
)

func (i ItemsPerRow) String() string {
	return fmt.Sprintf("col-%d", i)
}
