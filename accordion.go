// +build js,wasm

package srx

type (
	Accordion interface {
		TComponent
	}

	AccordionItem interface {
		TComponent
	}

	accordion struct {
		*component
	}

	accordionItem struct {
		*component
	}
)

func newAccordionItem(owner TComponent) *accordionItem {
	return nil
}
