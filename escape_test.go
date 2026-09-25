package weaver

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscapeAttributeValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "plain string",
			input:    "hello-world_123",
			expected: "hello-world_123",
		},
		{
			name:     "double quotes",
			input:    `class="btn"`,
			expected: `class=&quot;btn&quot;`,
		},
		{
			name:     "ampersands",
			input:    "foo&bar&baz",
			expected: "foo&amp;bar&amp;baz",
		},
		{
			name:     "single quotes",
			input:    "It's a test",
			expected: "It&#39;s a test",
		},
		{
			name:     "mixed special characters",
			input:    `"fish & chips" - 'cheap'`,
			expected: `&quot;fish &amp; chips&quot; - &#39;cheap&#39;`,
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := &bytes.Buffer{}

			err := EscapeAttributeValue(buf, test.input)

			assert.Nil(t, err)
			assert.Equal(t, test.expected, buf.String())
		})
	}
}

func TestElementWithEscapedAttribute(t *testing.T) {
	buf := &bytes.Buffer{}

	err := El("div", Attr("data-info", `<script>alert("xss")&'test'</script>`)).RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, `<div data-info="<script>alert(&quot;xss&quot;)&amp;&#39;test&#39;</script>"></div>`, buf.String())
}

func TestVoidElementWithAttributeAndEscaping(t *testing.T) {
	buf := &bytes.Buffer{}

	err := VoidEl("img", Attr("alt", "A & B's \"photo\"")).RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, `<img alt="A &amp; B&#39;s &quot;photo&quot;" />`, buf.String())
}

func TestEscapeText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "plain string",
			input:    "Hello World!",
			expected: "Hello World!",
		},
		{
			name:     "less than and greater than",
			input:    "5 < 10 && 10 > 5",
			expected: "5 &lt; 10 &amp;&amp; 10 &gt; 5",
		},
		{
			name:     "html script tags",
			input:    "<script>alert('xss')</script>",
			expected: "&lt;script&gt;alert('xss')&lt;/script&gt;",
		},
		{
			name:     "quotes remain unescaped in text body",
			input:    `"hello" & 'world'`,
			expected: `"hello" &amp; 'world'`,
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := &bytes.Buffer{}

			err := EscapeText(buf, test.input)

			assert.Nil(t, err)
			assert.Equal(t, test.expected, buf.String())
		})
	}
}

func TestElementWithEscapedText(t *testing.T) {
	buf := &bytes.Buffer{}

	err := El("p", Text("1 < 2 & 3 > 2")).RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, "<p>1 &lt; 2 &amp; 3 &gt; 2</p>", buf.String())
}

func TestElementWithUnescapedText(t *testing.T) {
	buf := &bytes.Buffer{}

	err := El("div", UnsafeRawText("<h1>Hello</h1>")).RenderHtml(buf)

	assert.Nil(t, err)
	assert.Equal(t, "<div><h1>Hello</h1></div>", buf.String())
}
