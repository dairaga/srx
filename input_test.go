// +build js,wasm

package srx

import (
	"testing"

	"github.com/dairaga/srx/enum"
	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	typ := "text"
	name := ""
	value := ""
	readonly := false
	required := false

	input := Input(Root())
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

	container := newComponent(
		Root(),
		js.From(`<div class="container"></div>`),
	)

	form := newComponent(
		container,
		js.Create("form"),
	)
	container.Append(form)
	js.Append(container)

	gutter := GutterPanel(form)
	gutter.SetGutterSize(enum.N3, enum.N3)
	gutter.SetItemsPerRow(enum.N2PerRow)
	form.Append(gutter)

	email := Email(gutter)
	email.SetID("test_email")
	lb := Label(gutter)
	lb.SetCaption("email")
	lb.SetFor(email.ID())
	gutter.Append(lb, email)

	password := Password(gutter)
	password.SetID("test_password")
	lb = Label(gutter)
	lb.SetCaption("password")
	lb.SetFor(password.ID())
	gutter.Append(lb, password)

	name2 := Text(gutter)
	name2.SetID("test_name")
	lb = Label(gutter)
	lb.SetCaption("name")
	lb.SetFor(name2.ID())
	gutter.Append(lb, name2)

	color := ColorPicker(gutter)
	color.SetID("test_color")
	lb = Label(gutter)
	lb.SetCaption("color")
	lb.SetFor(color.ID())
	gutter.Append(lb, color)

	file := File(gutter)
	file.SetID("test_file")
	lb = Label(gutter)
	lb.SetCaption("file")
	lb.SetFor(file.ID())
	gutter.Append(lb, file)

	rng := Range(gutter)
	rng.SetID("test_range")
	lb = Label(gutter)
	lb.SetCaption("range")
	lb.SetFor(rng.ID())
	gutter.Append(lb, rng)

	num := Number(gutter)
	num.SetID("test_number")
	lb = Label(gutter)
	lb.SetCaption("number")
	lb.SetFor(num.ID())
	gutter.Append(lb, num)

	sel := Select(gutter)
	sel.AddOption("test1", "1")
	sel.AddOption("test2", "2")
	sel.AddOption("test3", "3")
	sel.SetValue("2")
	sel.SetID("test_sel")
	lb = Label(gutter)
	lb.SetCaption("select")
	lb.SetFor(sel.ID())
	gutter.Append(lb, sel)

	ckgroup := CheckGroup(gutter)
	ckgroup.Switch(true)
	ckgroup.SetName("test_check")
	ckgroup.SetInline(true)
	ckgroup.AddCheck("A", "AAA", true)
	ckgroup.AddCheck("B", "BBB", false)
	lb = Label(gutter)
	lb.SetCaption("checkbox")
	gutter.Append(lb, ckgroup)
	assert.EqualValues(t, []string{"A"}, ckgroup.Value())
	assert.True(t, ckgroup.Child(0).Checked())
	ckgroup.SetValue("A", "B")
	assert.True(t, ckgroup.Child(0).Checked())
	assert.True(t, ckgroup.Child(1).Checked())

	ragroup := RadioGroup(gutter)
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

	lb = Label(gutter)
	lb.SetCaption("radio")
	ragroup.Switch(true)
	gutter.Append(lb, ragroup)
}
