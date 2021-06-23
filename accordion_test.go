// +build js,wasm

package srx

import (
	"testing"

	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestAccordion(t *testing.T) {
	a := Accordion(Root())
	var items []TAccordionItem

	js.Append(a)
	items = append(items,
		a.AddItem(
			"Accordion Item #1",
			Strong("This is the first item's accordion body."),
			Caption("It is shown by default")))

	items = append(items,
		a.AddItem(
			"Accordion Item #2",
			Strong("This is the second item's accordion body."),
		))

	items = append(items,
		a.AddItem(
			"Accordion Item #3",
			Strong("This is the third item's accordion body."),
		))
	items[0].Active(true)
	assert.EqualValues(t, items, a.Items())

	parentID := a.ID()

	for i := range items {
		headerID := items[i].Ref().Query(".accordion-header").ID()
		collapseID := items[i].Ref().Query(".accordion-collapse").ID()

		assert.Equal(t,
			items[i].Ref().Query(".accordion-collapse").Attr("data-bs-parent"),
			"#"+parentID)

		assert.Equal(t,
			items[i].Ref().Query(".accordion-collapse").Attr("aria-labelledby"),
			headerID)

		assert.Equal(t,
			items[i].Ref().Query(".accordion-button").Attr("data-bs-target"),
			"#"+collapseID)
		assert.Equal(t,
			items[i].Ref().Query(".accordion-button").Attr("aria-controls"),
			collapseID,
		)
	}

	assert.True(t, items[0].Ref().Query(".accordion-collapse").Contains("show"))

	a.SetID("test_accordion")
	parentID = a.ID()
	for i := range items {
		assert.Equal(t,
			items[i].Ref().Query(".accordion-collapse").Attr("data-bs-parent"),
			"#"+parentID)
	}
}
