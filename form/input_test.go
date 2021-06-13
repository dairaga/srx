// +build js,wasm

package form

import (
	"testing"

	"github.com/dairaga/srx"
	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/dairaga/srx/std"
	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	typ := ""
	name := ""
	value := ""
	readonly := false
	required := false

	input := InputOf(srx.Root())
	js.Append(input)

	/* init */
	assert.Equal(t, typ, input.Type())
	assert.Equal(t, name, input.Name())
	assert.Equal(t, value, input.Value())
	assert.Equal(t, readonly, input.ReadOnly())
	assert.Equal(t, required, input.Required())

	typ = "text"
	name = "test_input"
	value = "test_input_value"
	readonly = true
	required = true

	input.SetType(typ)
	input.SetName(name)
	input.SetValue(value)
	input.SetReadOnly(readonly)
	input.SetRequired(required)

	assert.Equal(t, typ, input.Type())
	assert.Equal(t, name, input.Name())
	assert.Equal(t, value, input.Value())
	assert.Equal(t, readonly, input.ReadOnly())
	assert.Equal(t, required, input.Required())

	/* View All */

	container := srx.NewComponent(
		srx.Root(),
		js.From(`<div class="container"></div>`),
	)

	form := srx.NewComponent(
		container,
		js.Create("form"),
	)
	container.Append(form)
	js.Append(container)

	gutter := std.GutterPanelOf(form)
	gutter.SetGutterSize(enum.N3, enum.N3)
	gutter.SetItemsPerRow(enum.N2PerRow)
	form.Append(gutter)

	p := std.PanelOf(gutter)
	email := EmailOf(p)
	email.SetID("test_email")
	lb := LabelOf(p)
	lb.SetCaption("email")
	lb.SetFor(email.ID())
	p.Append(lb)
	p.Append(email)
	gutter.Append(p)

	p = std.PanelOf(gutter)
	password := PasswordOf(p)
	password.SetID("test_password")
	lb = LabelOf(p)
	lb.SetCaption("password")
	lb.SetFor(password.ID())
	p.Append(lb)
	p.Append(password)
	gutter.Append(p)

	p = std.PanelOf(gutter)
	name2 := TextOf(p)
	name2.SetID("test_name")
	lb = LabelOf(p)
	lb.SetCaption("name")
	lb.SetFor(name2.ID())
	p.Append(lb)
	p.Append(name2)
	gutter.Append(p)

	p = std.PanelOf(gutter)
	color := ColorPickerOf(p)
	color.SetID("test_color")
	lb = LabelOf(p)
	lb.SetCaption("color")
	lb.SetFor(color.ID())
	p.Append(lb)
	p.Append(color)
	gutter.Append(p)

	p = std.PanelOf(gutter)
	file := FileOf(p)
	file.SetID("test_file")
	lb = LabelOf(p)
	lb.SetCaption("file")
	lb.SetFor(file.ID())
	p.Append(lb)
	p.Append(file)
	gutter.Append(p)

	p = std.PanelOf(gutter)
	rng := RangeOf(p)
	rng.SetID("test_range")
	lb = LabelOf(p)
	lb.SetCaption("range")
	lb.SetFor(rng.ID())
	p.Append(lb)
	p.Append(rng)
	gutter.Append(p)

	p = std.PanelOf(gutter)
	num := NumberOf(p)
	num.SetID("test_number")
	lb = LabelOf(p)
	lb.SetCaption("number")
	lb.SetFor(num.ID())
	p.Append(lb)
	p.Append(num)
	gutter.Append(p)

	p = std.PanelOf(gutter)
	sel := SelectOf(p)
	sel.AddOption("test1", "1")
	sel.AddOption("test2", "2")
	sel.AddOption("test3", "3")
	sel.SetValue("2")
	sel.SetID("test_sel")
	lb = LabelOf(p)
	lb.SetCaption("select")
	lb.SetFor(sel.ID())
	p.Append(lb)
	p.Append(sel)
	gutter.Append(p)

	p = std.PanelOf(gutter)
	ckgroup := CheckGroupOf(p)
	ckgroup.SetName("test_check")
	ckgroup.SetInline(true)
	ckgroup.AddCheck("A", "AAA", true)
	ckgroup.AddCheck("B", "BBB", false)
	lb = LabelOf(p)
	lb.SetCaption("checkbox")
	p.Append(lb)
	p.Append(ckgroup)
	gutter.Append(p)

	p = std.PanelOf(gutter)
	ckgroup = CheckGroupOf(p)
	ckgroup.SetName("test_Radio")
	ckgroup.SetInline(true)
	ckgroup.AddRadio("Y", "Yes", true)
	ckgroup.AddRadio("N", "No", false)
	lb = LabelOf(p)
	lb.SetCaption("checkbox")
	p.Append(lb)
	p.Append(ckgroup)
	gutter.Append(p)
}
