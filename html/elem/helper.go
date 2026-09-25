package elem

import "atomicptr.dev/weaver"

// Text returns HTML-escaped text
func Text(text string) weaver.Node {
	return weaver.Text(text)
}

// UnsafeRawText returns unescaped text directly into the output
func UnsafeRawText(text string) weaver.Node {
	return weaver.UnsafeRawText(text)
}

// Fragment returns multiple elements as is without a wrapper element
func Fragment(nodes ...weaver.Node) weaver.Node {
	return weaver.Fragment(nodes...)
}
