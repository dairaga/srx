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
	typ := "text"
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

	typ = "password"
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

	email := EmailOf(gutter)
	email.SetID("test_email")
	lb := LabelOf(gutter)
	lb.SetCaption("email")
	lb.SetFor(email.ID())
	gutter.Append(lb, email)

	password := PasswordOf(gutter)
	password.SetID("test_password")
	lb = LabelOf(gutter)
	lb.SetCaption("password")
	lb.SetFor(password.ID())
	gutter.Append(lb, password)

	name2 := TextOf(gutter)
	name2.SetID("test_name")
	lb = LabelOf(gutter)
	lb.SetCaption("name")
	lb.SetFor(name2.ID())
	gutter.Append(lb, name2)

	color := ColorPickerOf(gutter)
	color.SetID("test_color")
	lb = LabelOf(gutter)
	lb.SetCaption("color")
	lb.SetFor(color.ID())
	gutter.Append(lb, color)

	file := FileOf(gutter)
	file.SetID("test_file")
	lb = LabelOf(gutter)
	lb.SetCaption("file")
	lb.SetFor(file.ID())
	gutter.Append(lb, file)

	rng := RangeOf(gutter)
	rng.SetID("test_range")
	lb = LabelOf(gutter)
	lb.SetCaption("range")
	lb.SetFor(rng.ID())
	gutter.Append(lb, rng)

	num := NumberOf(gutter)
	num.SetID("test_number")
	lb = LabelOf(gutter)
	lb.SetCaption("number")
	lb.SetFor(num.ID())
	gutter.Append(lb, num)

	sel := SelectOf(gutter)
	sel.AddOption("test1", "1")
	sel.AddOption("test2", "2")
	sel.AddOption("test3", "3")
	sel.SetValue("2")
	sel.SetID("test_sel")
	lb = LabelOf(gutter)
	lb.SetCaption("select")
	lb.SetFor(sel.ID())
	gutter.Append(lb, sel)

	ckgroup := CheckGroupOf(gutter)
	ckgroup.SetName("test_check")
	ckgroup.SetInline(true)
	ckgroup.AddCheck("A", "AAA", true)
	ckgroup.AddCheck("B", "BBB", false)
	lb = LabelOf(gutter)
	lb.SetCaption("checkbox")
	gutter.Append(lb, ckgroup)
	assert.EqualValues(t, []string{"A"}, ckgroup.Value())
	assert.True(t, ckgroup.Child(0).Checked())
	ckgroup.SetValue("A", "B")
	assert.True(t, ckgroup.Child(0).Checked())
	assert.True(t, ckgroup.Child(1).Checked())

	ragroup := RadioGroupOf(gutter)
	ragroup.SetName("test_Radio")
	ragroup.SetInline(true)
	ragroup.AddRadio("Y", "Yes", true)
	ragroup.AddRadio("N", "No", false)
	assert.Equal(t, "Y", ragroup.Value())
	assert.True(t, ragroup.Child(0).Checked())
	assert.False(t, ragroup.Child(1).Checked())

	ragroup.SetValue("N")
	assert.Equal(t, "N", ragroup.Value())
	assert.False(t, ragroup.Child(0).Checked())
	assert.True(t, ragroup.Child(1).Checked())

	lb = LabelOf(gutter)
	lb.SetCaption("radio")
	gutter.Append(lb, ragroup)
}
