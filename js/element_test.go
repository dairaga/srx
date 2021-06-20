// +build js,wasm

package js_test

import (
	"fmt"
	"testing"

	"github.com/dairaga/srx/js"
	"github.com/stretchr/testify/assert"
)

func TestElementOK(t *testing.T) {
	el := js.HTML(`<div></div>`).Element()
	assert.True(t, el.OK())

	el = js.ElementOf(js.ValueOf(""))
	assert.False(t, el.OK())

}

func TestElementProp(t *testing.T) {
	text := `abc`
	el := js.HTML(`<div>` + text + `</div>`).Element()
	assert.Equal(t, text, el.Prop("innerText").String())
	assert.NotPanics(t, func() { _ = el.Prop("inna") })
	el.SetProp("innerText", text+text)
	assert.Equal(t, text+text, el.Prop("innerText").String())
}

func TestElementAttr(t *testing.T) {
	v := `test`
	el := js.HTML(`<div data-test="` + v + `"></div>`).Element()
	assert.True(t, el.HasAttr("data-test"))
	assert.Equal(t, v, el.Attr("data-test"))

	el.SetAttr("data-test", v+v)
	assert.Equal(t, v+v, el.Attr("data-test"))
	el.RemoveAttr("data-test")
	assert.Equal(t, "", el.Attr("data-test"))
	assert.False(t, el.HasAttr("data-test"))

	attrs := map[string]string{
		"data-a": "a",
		"data-b": "b",
		"data-c": "c",
	}

	el.AddAttrs(attrs)

	for _, testv := range []string{"a", "b", "c"} {
		assert.Equal(t, testv, el.Attr("data-"+testv))
	}
}

func TestElementTagName(t *testing.T) {
	assert.Equal(t, "div", js.HTML(`<DIV></div>`).Element().TagName())
}

func TestElementID(t *testing.T) {
	el := js.HTML(`<div id="test"></div>`).Element()
	assert.Equal(t, "test", el.ID())

	el.SetID("test2")
	assert.Equal(t, "test2", el.ID())
}

func TestElementText(t *testing.T) {
	text := `test`
	el := js.HTML(`<div>` + text + `</div>`).Element()
	assert.Equal(t, text, el.Text())
	el.SetText(text + text)
	assert.Equal(t, text+text, el.Text())
}

func TestElementHTML(t *testing.T) {
	id := `test`
	con := `Hello`
	html := `<span id="` + id + `">` + con + `</span>`

	el := js.HTML(`<div>` + html + `</div>`).Element()
	assert.Equal(t, con, el.Text())
	assert.True(t, el.Query("#"+id).OK())
	assert.Equal(t, js.HTML(html), el.HTML())

	html = `<span id="` + id + `">` + con + con + `</span>`
	el.SetHTML(js.HTML(html))
	assert.Equal(t, con+con, el.Text())
	assert.True(t, el.Query("#"+id).OK())
	assert.Equal(t, js.HTML(html), el.HTML())
}

func TestElementAppendPrepend(t *testing.T) {
	el := js.HTML(`<div><span id='middle'>middle</span></div>`).Element()
	el.Append(js.HTML(`<span id='right'>right</span>`))
	els := el.QueryAll("span")
	assert.Equal(t, "right", els[1].Text())

	el.Prepend(js.HTML(`<span id='left'>left</span>`))
	els = el.QueryAll("span")
	assert.Equal(t, "left", els[0].Text())
	assert.Equal(t, "middle", els[1].Text())
	assert.Equal(t, "right", els[2].Text())
}

func TestElementClass(t *testing.T) {
	el := js.HTML(`<div></div>`).Element()
	el.Add("class_a", "class_b")
	assert.True(t, el.Contains("class_a"))
	assert.True(t, el.Contains("class_b"))
	assert.False(t, el.Contains("class_c"))

	el.Remove("class_b")
	assert.True(t, el.Contains("class_a"))
	assert.False(t, el.Contains("class_b"))

	el.Toggle("class_a")
	el.Toggle("class_b")
	assert.False(t, el.Contains("class_a"))
	assert.True(t, el.Contains("class_b"))

	el.Replace("class_b", "class_c")
	assert.True(t, el.Contains("class_c"))
	assert.False(t, el.Contains("class_b"))
}

func TestStyle(t *testing.T) {
	el := js.HTML(`<div style="color: red">ABC</div>`).Element()
	js.Append(el)
	assert.Equal(t, "red", el.Style("color"))

	el.SetStyle("color", "blue")
	assert.Equal(t, "blue", el.Style("color"))

	el.RemoveStyle("color")
	assert.Equal(t, "", el.Style("color"))
}

func TestHidden(t *testing.T) {
	elA := js.HTML(`<div>H1</div>`).Element()
	js.Append(elA)

	el := js.HTML(`<div style="color: red">ABC</div>`).Element()
	js.Append(el)

	elB := js.HTML(`<div>XYZ</div>`).Element()
	js.Append(elB)

	el.Hide()
	assert.True(t, el.Prop("hidden").Bool())
	assert.Equal(t, 0, el.Prop("offsetWidth").Int())
	assert.Equal(t, 0, el.Prop("offsetHeight").Int())

	el.Show()
	assert.False(t, el.Prop("hidden").Bool())
	assert.Equal(t, elB.Prop("offsetTop").Int(), elA.Prop("offsetTop").Int()+elA.Prop("offsetHeight").Int()+el.Prop("offsetHeight").Int())
}

func TestDisabled(t *testing.T) {
	el := js.HTML(`<div></div>`).Element()
	js.Append(el)

	assert.False(t, el.Disabled())

	el.Disable()
	assert.True(t, el.Disabled())

	el.Enable()
	assert.False(t, el.Disabled())
}

func TestRelease(t *testing.T) {
	p1 := js.Create("div")
	p1.SetID("test_release_1")
	p2 := js.Create("div")
	p2.SetID("test_release_2")
	c1 := js.Create("div").SetText("child 1")
	c2 := js.Create("div").SetText("child 2")

	p1.Append(c1)
	p2.Append(c2)
	js.Append(p1)
	js.Append(p2)

	c1.Release()
	fmt.Println(c1.OK())
	p2.Append(c1)
}
