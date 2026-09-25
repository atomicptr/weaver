package elem

import "atomicptr.dev/weaver"

// A creates the HTML <a> element (or anchor element), which with its href attribute creates a hyperlink to web pages,
// files, email addresses, locations in the same page, or anything else a URL can address.
func A(nodes ...weaver.Node) weaver.Node {
	return weaver.El("a", nodes...)
}

// Abbr creates the HTML <abbr> element representing an abbreviation or acronym.
func Abbr(nodes ...weaver.Node) weaver.Node {
	return weaver.El("abbr", nodes...)
}

// Address creates the HTML <address> element indicating that the enclosed HTML provides contact information for a
// person or people, or for an organization.
func Address(nodes ...weaver.Node) weaver.Node {
	return weaver.El("address", nodes...)
}

// Article creates the HTML <article> element representing a self-contained composition in a document, page,
// application, or site, which is intended to be independently distributable or reusable.
func Article(nodes ...weaver.Node) weaver.Node {
	return weaver.El("article", nodes...)
}

// Aside creates the HTML <aside> element representing a portion of a document whose content is only indirectly
// related to the document's main content.
func Aside(nodes ...weaver.Node) weaver.Node {
	return weaver.El("aside", nodes...)
}

// Audio creates the HTML <audio> element used to embed sound content in documents.
func Audio(nodes ...weaver.Node) weaver.Node {
	return weaver.El("audio", nodes...)
}

// B creates the HTML <b> element used to draw the reader's attention to the element's contents, which are not
// otherwise granted special importance.
func B(nodes ...weaver.Node) weaver.Node {
	return weaver.El("b", nodes...)
}

// Base creates the HTML <base> element specifying the base URL to use for all relative URLs in a document.
func Base(nodes ...weaver.Node) weaver.Node {
	return weaver.VoidEl("base", nodes...)
}

// Blockquote creates the HTML <blockquote> element indicating that the enclosed text is an extended quotation.
// Usually, this is rendered visually by indentation.
func Blockquote(nodes ...weaver.Node) weaver.Node {
	return weaver.El("blockquote", nodes...)
}

// Body creates the HTML <body> element representing the content of an HTML document. There can be only one <body>
// element in a document.
func Body(nodes ...weaver.Node) weaver.Node {
	return weaver.El("body", nodes...)
}

// Br creates the HTML <br> element representing a line break. It is a void element.
func Br(nodes ...weaver.Node) weaver.Node {
	return weaver.VoidEl("br", nodes...)
}

// Button creates the HTML <button> element as an interactive element activated by a user with a mouse, keyboard,
// finger, voice command, or other assistive technology. Once activated, it then performs an action, such as submitting
// a form or opening a dialog.
func Button(nodes ...weaver.Node) weaver.Node {
	return weaver.El("button", nodes...)
}

// Canvas creates the HTML <canvas> element used to draw graphics, on the fly, via JavaScript.
func Canvas(nodes ...weaver.Node) weaver.Node {
	return weaver.El("canvas", nodes...)
}

// Code creates the HTML <code> element displaying its contents styled in a fashion intended to indicate that the
// text is a short fragment of computer code.
func Code(nodes ...weaver.Node) weaver.Node {
	return weaver.El("code", nodes...)
}

// Datalist creates the HTML <datalist> element containing a set of <option> elements that represent the permissible
// or recommended options available to choose from within other controls.
func Datalist(nodes ...weaver.Node) weaver.Node {
	return weaver.El("datalist", nodes...)
}

// Dd creates the HTML <dd> element providing the description, definition, or value for the preceding term (<dt>) in
// a description list (<dl>).
func Dd(nodes ...weaver.Node) weaver.Node {
	return weaver.El("dd", nodes...)
}

// Del creates the HTML <del> element representing a range of text that has been deleted from a document.
func Del(nodes ...weaver.Node) weaver.Node {
	return weaver.El("del", nodes...)
}

// Details creates the HTML <details> element creating a disclosure widget in which information is visible only when
// the widget is toggled into an "open" state.
func Details(nodes ...weaver.Node) weaver.Node {
	return weaver.El("details", nodes...)
}

// Dialog creates the HTML <dialog> element representing a dialog box or other interactive component, such as a
// dismissible alert, inspector, or subwindow.
func Dialog(nodes ...weaver.Node) weaver.Node {
	return weaver.El("dialog", nodes...)
}

// Div creates the HTML <div> element as the generic container for flow content. It has no effect on the content or
// layout until styled in some way using CSS.
func Div(nodes ...weaver.Node) weaver.Node {
	return weaver.El("div", nodes...)
}

// Dl creates the HTML <dl> element representing a description list. The element encloses a list of groups of terms
// (specified using the <dt> element) and descriptions (provided by <dd> elements).
func Dl(nodes ...weaver.Node) weaver.Node {
	return weaver.El("dl", nodes...)
}

// Dt creates the HTML <dt> element specifying a term in a description or definition list, and as such must be used
// inside a <dl> element.
func Dt(nodes ...weaver.Node) weaver.Node {
	return weaver.El("dt", nodes...)
}

// Em creates the HTML <em> element marking text that has stress emphasis. The <em> element can be nested, with each
// level of nesting indicating a greater degree of emphasis.
func Em(nodes ...weaver.Node) weaver.Node {
	return weaver.El("em", nodes...)
}

// Fieldset creates the HTML <fieldset> element used to group several controls as well as labels (<label>)
// within a web form.
func Fieldset(nodes ...weaver.Node) weaver.Node {
	return weaver.El("fieldset", nodes...)
}

// Figcaption creates the HTML <figcaption> element representing a caption or legend describing the rest of the
// contents of its parent <figure> element, providing the <figure> an accessible name.
func Figcaption(nodes ...weaver.Node) weaver.Node {
	return weaver.El("figcaption", nodes...)
}

// Figure creates the HTML <figure> element representing self-contained content, potentially with an optional caption,
// which is specified using the <figcaption> element. The figure, its caption, and its contents are referenced as a
// single unit.
func Figure(nodes ...weaver.Node) weaver.Node {
	return weaver.El("figure", nodes...)
}

// Footer creates the HTML <footer> element representing a footer for its nearest ancestor sectioning content or
// sectioning root element.
func Footer(nodes ...weaver.Node) weaver.Node {
	return weaver.El("footer", nodes...)
}

// Form creates the HTML <form> element representing a document section containing interactive controls for
// submitting information.
func Form(nodes ...weaver.Node) weaver.Node {
	return weaver.El("form", nodes...)
}

// H1 creates the HTML <h1> element representing the highest section level heading.
func H1(nodes ...weaver.Node) weaver.Node {
	return weaver.El("h1", nodes...)
}

// H2 creates the HTML <h2> element representing a second-level section heading.
func H2(nodes ...weaver.Node) weaver.Node {
	return weaver.El("h2", nodes...)
}

// H3 creates the HTML <h3> element representing a third-level section heading.
func H3(nodes ...weaver.Node) weaver.Node {
	return weaver.El("h3", nodes...)
}

// H4 creates the HTML <h4> element representing a fourth-level section heading.
func H4(nodes ...weaver.Node) weaver.Node {
	return weaver.El("h4", nodes...)
}

// H5 creates the HTML <h5> element representing a fifth-level section heading.
func H5(nodes ...weaver.Node) weaver.Node {
	return weaver.El("h5", nodes...)
}

// H6 creates the HTML <h6> element representing the lowest section level heading.
func H6(nodes ...weaver.Node) weaver.Node {
	return weaver.El("h6", nodes...)
}

// Head creates the HTML <head> element containing machine-readable information (metadata) about the document, like
// its title, scripts, and style sheets.
func Head(nodes ...weaver.Node) weaver.Node {
	return weaver.El("head", nodes...)
}

// Header creates the HTML <header> element representing introductory content, typically a group of introductory or
// navigational aids.
func Header(nodes ...weaver.Node) weaver.Node {
	return weaver.El("header", nodes...)
}

// Hr creates the HTML <hr> element representing a thematic break between paragraph-level elements.
func Hr(nodes ...weaver.Node) weaver.Node {
	return weaver.VoidEl("hr", nodes...)
}

// Html creates the HTML <html> element representing the root (top-level element) of an HTML document. All other
// elements must be descendants of this element.
func Html(nodes ...weaver.Node) weaver.Node {
	return weaver.Fragment(
		weaver.UnsafeRawText("<!DOCTYPE html>"),
		weaver.El("html", nodes...),
	)
}

// I creates the HTML <i> element representing a range of text that is set off from the normal text for some reason,
// such as idiomatic text, technical terms, taxonomic designations, among others.
func I(nodes ...weaver.Node) weaver.Node {
	return weaver.El("i", nodes...)
}

// Iframe creates the HTML <iframe> element representing a nested browsing context, embedding another HTML page into
// the current one.
func Iframe(nodes ...weaver.Node) weaver.Node {
	return weaver.El("iframe", nodes...)
}

// Img creates the HTML <img> element embedding an image into the document.
func Img(nodes ...weaver.Node) weaver.Node {
	return weaver.VoidEl("img", nodes...)
}

// Input creates the HTML <input> element used to create interactive controls for web-based forms in order to accept
// data from the user.
func Input(nodes ...weaver.Node) weaver.Node {
	return weaver.VoidEl("input", nodes...)
}

// Ins creates the HTML <ins> element representing a range of text that has been added to a document.
func Ins(nodes ...weaver.Node) weaver.Node {
	return weaver.El("ins", nodes...)
}

// Label creates the HTML <label> element representing a caption for an item in a user interface.
func Label(nodes ...weaver.Node) weaver.Node {
	return weaver.El("label", nodes...)
}

// Legend creates the HTML <legend> element representing a caption for the content of its parent <fieldset>.
func Legend(nodes ...weaver.Node) weaver.Node {
	return weaver.El("legend", nodes...)
}

// Li creates the HTML <li> element used to represent an item in a list.
func Li(nodes ...weaver.Node) weaver.Node {
	return weaver.El("li", nodes...)
}

// Link creates the HTML <link> element specifying relationships between the current document and an external resource.
func Link(nodes ...weaver.Node) weaver.Node {
	return weaver.VoidEl("link", nodes...)
}

// Main creates the HTML <main> element representing the dominant content of the <body> of a document.
func Main(nodes ...weaver.Node) weaver.Node {
	return weaver.El("main", nodes...)
}

// Mark creates the HTML <mark> element representing text which is marked or highlighted for reference or notation
// purposes.
func Mark(nodes ...weaver.Node) weaver.Node {
	return weaver.El("mark", nodes...)
}

// Meta creates the HTML <meta> element representing metadata that cannot be represented by other HTML meta-related
// elements.
func Meta(nodes ...weaver.Node) weaver.Node {
	return weaver.VoidEl("meta", nodes...)
}

// Meter creates the HTML <meter> element representing either a scalar value within a known range or a fractional
// value.
func Meter(nodes ...weaver.Node) weaver.Node {
	return weaver.El("meter", nodes...)
}

// Nav creates the HTML <nav> element representing a section of a page whose purpose is to provide navigation links.
func Nav(nodes ...weaver.Node) weaver.Node {
	return weaver.El("nav", nodes...)
}

// Ol creates the HTML <ol> element representing an ordered list of items — typically rendered as a numbered list.
func Ol(nodes ...weaver.Node) weaver.Node {
	return weaver.El("ol", nodes...)
}

// Optgroup creates the HTML <optgroup> element creating a grouping of options within a <select> element.
func Optgroup(nodes ...weaver.Node) weaver.Node {
	return weaver.El("optgroup", nodes...)
}

// Option creates the HTML <option> element used to define an item contained in a <select>, an <optgroup>, or a
// <datalist> element.
func Option(nodes ...weaver.Node) weaver.Node {
	return weaver.El("option", nodes...)
}

// P creates the HTML <p> element representing a paragraph.
func P(nodes ...weaver.Node) weaver.Node {
	return weaver.El("p", nodes...)
}

// Pre creates the HTML <pre> element representing preformatted text which is to be presented exactly as written in
// the HTML file.
func Pre(nodes ...weaver.Node) weaver.Node {
	return weaver.El("pre", nodes...)
}

// Progress creates the HTML <progress> element displaying an indicator showing the completion progress of a task,
// typically displayed as a progress bar.
func Progress(nodes ...weaver.Node) weaver.Node {
	return weaver.El("progress", nodes...)
}

// Script creates the HTML <script> element used to embed executable code or data; this is typically used to embed or
// refer to JavaScript code.
func Script(nodes ...weaver.Node) weaver.Node {
	return weaver.El("script", nodes...)
}

// Section creates the HTML <section> element representing a generic standalone section of a document.
func Section(nodes ...weaver.Node) weaver.Node {
	return weaver.El("section", nodes...)
}

// Select creates the HTML <select> element representing a control that provides a menu of options.
func Select(nodes ...weaver.Node) weaver.Node {
	return weaver.El("select", nodes...)
}

// Small creates the HTML <small> element representing side-comments and small print, like copyright and legal text,
// independent of its styled presentation.
func Small(nodes ...weaver.Node) weaver.Node {
	return weaver.El("small", nodes...)
}

// Span creates the HTML <span> element as a generic inline container for phrasing content.
func Span(nodes ...weaver.Node) weaver.Node {
	return weaver.El("span", nodes...)
}

// Strong creates the HTML <strong> element indicating that its contents have strong importance, seriousness, or
// urgency.
func Strong(nodes ...weaver.Node) weaver.Node {
	return weaver.El("strong", nodes...)
}

// Style creates the HTML <style> element containing style information for a document, or part of a document.
func Style(nodes ...weaver.Node) weaver.Node {
	return weaver.El("style", nodes...)
}

// Sub creates the HTML <sub> element specifying inline text which should be displayed as subscript for solely
// typographical reasons.
func Sub(nodes ...weaver.Node) weaver.Node {
	return weaver.El("sub", nodes...)
}

// Summary creates the HTML <summary> element specifying a summary, caption, or legend for a <details> element's
// disclosure box.
func Summary(nodes ...weaver.Node) weaver.Node {
	return weaver.El("summary", nodes...)
}

// Sup creates the HTML <sup> element specifying inline text which is to be displayed as superscript for solely
// typographical reasons.
func Sup(nodes ...weaver.Node) weaver.Node {
	return weaver.El("sup", nodes...)
}

// TBody creates the HTML <tbody> element encapsulating a set of table rows (<tr> elements), indicating that they
// comprise the body of the table (<table>).
func TBody(nodes ...weaver.Node) weaver.Node {
	return weaver.El("tbody", nodes...)
}

// TFoot creates the HTML <tfoot> element defining a set of rows summarizing the columns of the table.
func TFoot(nodes ...weaver.Node) weaver.Node {
	return weaver.El("tfoot", nodes...)
}

// THead creates the HTML <thead> element defining a set of rows defining the head of the columns of the table.
func THead(nodes ...weaver.Node) weaver.Node {
	return weaver.El("thead", nodes...)
}

// Table creates the HTML <table> element representing tabular data — that is, information presented in a
// two-dimensional table comprised of rows and columns of cells containing data.
func Table(nodes ...weaver.Node) weaver.Node {
	return weaver.El("table", nodes...)
}

// Td creates the HTML <td> element defining a cell of a table that contains data.
func Td(nodes ...weaver.Node) weaver.Node {
	return weaver.El("td", nodes...)
}

// TextArea creates the HTML <textarea> element representing a multi-line plain-text editing control.
func TextArea(nodes ...weaver.Node) weaver.Node {
	return weaver.El("textarea", nodes...)
}

// Th creates the HTML <th> element defining a cell as header of a group of table cells.
func Th(nodes ...weaver.Node) weaver.Node {
	return weaver.El("th", nodes...)
}

// Time creates the HTML <time> element representing a specific period in time.
func Time(nodes ...weaver.Node) weaver.Node {
	return weaver.El("time", nodes...)
}

// Title creates the HTML <title> element defining the document's title that is shown in a browser's title bar or a
// page's tab.
func Title(title string) weaver.Node {
	return weaver.El("title", weaver.Text(title))
}

// Tr creates the HTML <tr> element defining a row of cells in a table.
func Tr(nodes ...weaver.Node) weaver.Node {
	return weaver.El("tr", nodes...)
}

// U creates the HTML <u> element representing text that should be stylistically different from normal text, typically
// rendered as underlined.
func U(nodes ...weaver.Node) weaver.Node {
	return weaver.El("u", nodes...)
}

// Ul creates the HTML <ul> element representing an unordered list of items.
func Ul(nodes ...weaver.Node) weaver.Node {
	return weaver.El("ul", nodes...)
}

// Video creates the HTML <video> element embedding a media player which supports video playback into the document.
func Video(nodes ...weaver.Node) weaver.Node {
	return weaver.El("video", nodes...)
}
