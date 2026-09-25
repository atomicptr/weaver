package weaver

import "io"

// EscapeAttributeValue writes an escaped attribute value into the provided io.Writer
func EscapeAttributeValue(w io.Writer, value string) error {
	return escapeKeys(w, value, func(b byte) string {
		switch b {
		case '"':
			return "&quot;"
		case '&':
			return "&amp;"
		case '\'':
			return "&#39;"
		default:
			return ""
		}
	})
}

// EscapeText writes HTML-escaped text into the provided io.Writer
func EscapeText(w io.Writer, text string) error {
	return escapeKeys(w, text, func(b byte) string {
		switch b {
		case '<':
			return "&lt;"
		case '>':
			return "&gt;"
		case '&':
			return "&amp;"
		default:
			return ""
		}
	})
}

func escapeKeys(w io.Writer, s string, escapeFunc func(b byte) string) error {
	last := 0

	for i := 0; i < len(s); i++ {
		esc := escapeFunc(s[i])
		if esc == "" {
			continue
		}

		// write prior unescaped chunk
		if i > last {
			_, err := io.WriteString(w, s[last:i])
			if err != nil {
				return err
			}
		}

		// write escaped replacement
		_, err := io.WriteString(w, esc)
		if err != nil {
			return err
		}

		last = i + 1
	}

	// write remaining unescaped chunk
	if last < len(s) {
		_, err := io.WriteString(w, s[last:])
		if err != nil {
			return err
		}
	}

	return nil
}
