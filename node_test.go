package weaver

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleElement(t *testing.T) {
	buf := &bytes.Buffer{}

	err := El("div").RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, "<div></div>", buf.String())
}

func TestSimpleWithChild(t *testing.T) {
	buf := &bytes.Buffer{}

	err := El("hello", El("world")).RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, "<hello><world></world></hello>", buf.String())
}

func TestSimpleElementWithText(t *testing.T) {
	buf := &bytes.Buffer{}

	err := El("h1", Text("Hello, World!")).RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, "<h1>Hello, World!</h1>", buf.String())
}

func TestElementWithChildrenAndText(t *testing.T) {
	buf := &bytes.Buffer{}

	err := El("div", El("h1", Text("Hello!")), Text("World")).RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, "<div><h1>Hello!</h1>World</div>", buf.String())
}

func TestSimpleVoidElement(t *testing.T) {
	buf := &bytes.Buffer{}

	err := VoidEl("img").RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, "<img />", buf.String())
}

func TestErrChildInVoidElement(t *testing.T) {
	buf := &bytes.Buffer{}

	err := VoidEl("img", El("div", Text("Hello!"))).RenderHtml(buf)

	assert.NotNil(t, err)
}

func TestErrTextInVoidElement(t *testing.T) {
	buf := &bytes.Buffer{}

	err := VoidEl("img", Text("Hello!")).RenderHtml(buf)

	assert.NotNil(t, err)
}

func TestElementWithAttribute(t *testing.T) {
	buf := &bytes.Buffer{}

	err := El("div", Attr("class", "a b c"), Text("Hello!")).RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, `<div class="a b c">Hello!</div>`, buf.String())
}

func TestVoidElementWithMultipleAttributes(t *testing.T) {
	buf := &bytes.Buffer{}

	err := VoidEl("img", Attr("id", "hello-world"), Attr("class", "img rounded-full")).RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, `<img id="hello-world" class="img rounded-full" />`, buf.String())
}

func TestErrWhenAddingAttributeAfterChild(t *testing.T) {
	buf := &bytes.Buffer{}

	err := El("div", Attr("id", "test"), Text("Hello!"), Attr("class", "a b c")).RenderHtml(buf)

	assert.NotNil(t, err)
}

func TestFragment(t *testing.T) {
	buf := &bytes.Buffer{}

	err := Fragment(El("h1", Text("Hello World!")), El("h2", Text("Subtitle"))).RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, "<h1>Hello World!</h1><h2>Subtitle</h2>", buf.String())
}

func TestIf(t *testing.T) {
	buf := &bytes.Buffer{}

	err := El("h1", Text("Hello: "), If(true, func() Node {
		return Text("Yes")
	})).RenderHtml(buf)
	assert.Nil(t, err)
	assert.Equal(t, "<h1>Hello: Yes</h1>", buf.String())

	buf = &bytes.Buffer{}

	err = El("h1", Text("Hello: "), If(false, func() Node {
		return Text("No")
	})).RenderHtml(buf)
	assert.Nil(t, err)
	assert.Equal(t, "<h1>Hello: </h1>", buf.String())
}

func TestIfElse(t *testing.T) {
	buf := &bytes.Buffer{}

	err := El("h1", Text("Hello: "), IfElse(true, func() Node {
		return Text("Yes")
	}, func() Node {
		return Text("Nope")
	})).RenderHtml(buf)
	assert.Nil(t, err)
	assert.Equal(t, "<h1>Hello: Yes</h1>", buf.String())

	buf = &bytes.Buffer{}

	err = El("h1", Text("Hello: "), IfElse(false, func() Node {
		return Text("Yes")
	}, func() Node {
		return Text("Nope")
	})).RenderHtml(buf)
	assert.Equal(t, "<h1>Hello: Nope</h1>", buf.String())
}
