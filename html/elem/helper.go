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

// If returns the node provided by the closure when the condition is true
func If(condition bool, thenFunc func() weaver.Node) weaver.Node {
	return weaver.If(condition, thenFunc)
}

// IfElse returns the node provided by the first closure when the condition is true, otherwise it calls the else closure
func IfElse(condition bool, thenFunc func() weaver.Node, elseFunc func() weaver.Node) weaver.Node {
	return weaver.IfElse(condition, thenFunc, elseFunc)
}
