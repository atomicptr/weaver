package weaver

import "net/http"

// RenderHtmlResponse writes the Node directly into net/http ResponseWriter. Also sets Content-Type.
func RenderHtmlResponse(w http.ResponseWriter, node Node) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return node.RenderHtml(w)
}
