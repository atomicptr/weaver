package weaver

import (
	"errors"
	"io"
)

var (
	// ErrAttributeAfterChild Added an attribute after children have been added
	ErrAttributeAfterChild = errors.New("html: attribute passed after child node")

	// ErrChildInVoidElement Added a child to a void element
	ErrChildInVoidElement = errors.New("html: void elements can't have children")
)

type RenderContext struct {
	Writer io.Writer
	IsOpen bool
}

type Node func(ctx RenderContext) (RenderContext, error)

// RenderHtml renders the node tree into the provided io.Writer
func (n Node) RenderHtml(w io.Writer) error {
	ctx := RenderContext{
		Writer: w,
		IsOpen: false,
	}

	_, err := n(ctx)
	return err
}

// El returns a custom element node opening/closing tags
func El(tag string, nodes ...Node) Node {
	return func(ctx RenderContext) (RenderContext, error) {
		var err error

		ctx, err = ensureParentIsClosed(ctx)
		if err != nil {
			return ctx, err
		}

		_, err = io.WriteString(ctx.Writer, "<"+tag)
		if err != nil {
			return ctx, err
		}

		elementContext := RenderContext{Writer: ctx.Writer, IsOpen: true}

		for _, n := range nodes {
			if n == nil {
				continue
			}

			elementContext, err = n(elementContext)
			if err != nil {
				return ctx, err
			}
		}

		if elementContext.IsOpen {
			_, err := io.WriteString(ctx.Writer, ">")
			if err != nil {
				return ctx, err
			}
		}

		_, err = io.WriteString(ctx.Writer, "</"+tag+">")
		return ctx, err
	}
}

// VoidEl returns a custom element node that is self-closing
func VoidEl(tag string, nodes ...Node) Node {
	return func(ctx RenderContext) (RenderContext, error) {
		var err error

		ctx, err = ensureParentIsClosed(ctx)
		if err != nil {
			return ctx, err
		}

		_, err = io.WriteString(ctx.Writer, "<"+tag)
		if err != nil {
			return ctx, err
		}

		elementContext := RenderContext{Writer: ctx.Writer, IsOpen: true}

		for _, n := range nodes {
			if n == nil {
				continue
			}

			elementContext, err = n(elementContext)
			if err != nil {
				return ctx, err
			}

			// if any node flipped `IsInsideTag` to false, a child attempted to render
			if !elementContext.IsOpen {
				return ctx, ErrChildInVoidElement
			}
		}

		_, err = io.WriteString(ctx.Writer, " />")
		return ctx, err
	}
}

// Attr returns an attribute with escaped value
func Attr(key, value string) Node {
	return func(ctx RenderContext) (RenderContext, error) {
		if !ctx.IsOpen {
			return ctx, ErrAttributeAfterChild
		}

		_, err := io.WriteString(ctx.Writer, " "+key+`="`)
		if err != nil {
			return ctx, err
		}

		err = EscapeAttributeValue(ctx.Writer, value)
		if err != nil {
			return ctx, err
		}

		_, err = io.WriteString(ctx.Writer, `"`)
		return ctx, err
	}
}

// UnsafeRawAttr returns an attribute with unescaped value
func UnsafeRawAttr(key, value string) Node {
	return func(ctx RenderContext) (RenderContext, error) {
		if !ctx.IsOpen {
			return ctx, ErrAttributeAfterChild
		}

		_, err := io.WriteString(ctx.Writer, " "+key+`="`+value+`"`)
		return ctx, err
	}
}

// Flag returns an attribute without value
func Flag(key string, enabled bool) Node {
	return func(ctx RenderContext) (RenderContext, error) {
		if !ctx.IsOpen {
			return ctx, ErrAttributeAfterChild
		}

		if !enabled {
			return ctx, nil
		}

		_, err := io.WriteString(ctx.Writer, " "+key)
		return ctx, err
	}
}

// Text returns HTML-escaped text
func Text(text string) Node {
	return func(ctx RenderContext) (RenderContext, error) {
		var err error

		ctx, err = ensureParentIsClosed(ctx)
		if err != nil {
			return ctx, err
		}

		return ctx, EscapeText(ctx.Writer, text)
	}
}

// UnsafeRawText returns unescaped text directly into the output
func UnsafeRawText(text string) Node {
	return func(ctx RenderContext) (RenderContext, error) {
		var err error

		ctx, err = ensureParentIsClosed(ctx)
		if err != nil {
			return ctx, err
		}

		_, err = io.WriteString(ctx.Writer, text)
		return ctx, err
	}
}

// Fragment returns multiple elements as is without a wrapper element
func Fragment(nodes ...Node) Node {
	return func(ctx RenderContext) (RenderContext, error) {
		var err error

		for _, n := range nodes {
			if n == nil {
				continue
			}

			ctx, err = n(ctx)
			if err != nil {
				return ctx, err
			}
		}

		return ctx, nil
	}
}

// If returns the node provided by the closure when the condition is true
func If(condition bool, thenFunc func() Node) Node {
	return func(ctx RenderContext) (RenderContext, error) {
		if !condition {
			return ctx, nil
		}

		return thenFunc()(ctx)
	}
}

// IfElse returns the node provided by the first closure when the condition is true, otherwise it calls the else closure
func IfElse(condition bool, thenFunc func() Node, elseFunc func() Node) Node {
	return func(ctx RenderContext) (RenderContext, error) {
		if condition {
			return thenFunc()(ctx)
		}

		return elseFunc()(ctx)
	}
}

// Each creates a node for each element in a slice
func Each[T any](items []T, fn func(T) Node) Node {
	return func(ctx RenderContext) (RenderContext, error) {
		nodes := make([]Node, 0, len(items))

		for _, item := range items {
			nodes = append(nodes, fn(item))
		}

		return Fragment(nodes...)(ctx)
	}
}

func ensureParentIsClosed(ctx RenderContext) (RenderContext, error) {
	if ctx.IsOpen {
		ctx.IsOpen = false
		_, err := io.WriteString(ctx.Writer, ">")
		return ctx, err
	}
	return ctx, nil
}
