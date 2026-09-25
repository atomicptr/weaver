package attr

import "atomicptr.dev/weaver"

// AccessKey sets the accesskey attribute, providing a hint for generating a keyboard shortcut for the element.
func AccessKey(v string) weaver.Node {
	return weaver.Attr("accesskey", v)
}

// Autocapitalize sets the autocapitalize attribute, controlling whether and how text input is automatically
// capitalized.
func Autocapitalize(v string) weaver.Node {
	return weaver.Attr("autocapitalize", v)
}

// Autofocus sets the autofocus attribute, indicating that the element should be focused on page load.
func Autofocus(enabled bool) weaver.Node {
	return weaver.Flag("autofocus", enabled)
}

// Class sets the class attribute, specifying one or more class names for an element.
func Class(v string) weaver.Node {
	return weaver.Attr("class", v)
}

// ContentEditable sets the contenteditable attribute, indicating whether the element's content is editable.
func ContentEditable(v string) weaver.Node {
	return weaver.Attr("contenteditable", v)
}

// Dir sets the dir attribute, indicating the directionality of the element's text.
func Dir(v string) weaver.Node {
	return weaver.Attr("dir", v)
}

// Draggable sets the draggable attribute, indicating whether the element can be dragged.
func Draggable(v string) weaver.Node {
	return weaver.Attr("draggable", v)
}

// EnterKeyHint sets the enterkeyhint attribute, defining what action label (or icon) to present for the enter key on virtual keyboards.
func EnterKeyHint(v string) weaver.Node {
	return weaver.Attr("enterkeyhint", v)
}

// Hidden sets the hidden attribute, indicating that the element is not yet, or is no longer, relevant.
func Hidden(enabled bool) weaver.Node {
	return weaver.Flag("hidden", enabled)
}

// ID sets the id attribute, defining a unique identifier for the element.
func ID(v string) weaver.Node {
	return weaver.Attr("id", v)
}

// Inert sets the inert attribute, causing the browser to ignore input events for the element.
func Inert(enabled bool) weaver.Node {
	return weaver.Flag("inert", true)
}

// InputMode sets the inputmode attribute, providing a hint as to the type of virtual keyboard configuration to use.
func InputMode(v string) weaver.Node {
	return weaver.Attr("inputmode", v)
}

// Lang sets the lang attribute, specifying the primary language for the element's contents.
func Lang(v string) weaver.Node {
	return weaver.Attr("lang", v)
}

// Nonce sets the nonce attribute, defining a cryptographic nonce used by Content Security Policy.
func Nonce(v string) weaver.Node {
	return weaver.Attr("nonce", v)
}

// Popover sets the popover attribute, turning an element into a popover element.
func Popover(v string) weaver.Node {
	return weaver.Attr("popover", v)
}

// Spellcheck sets the spellcheck attribute, defining whether the element may be checked for spelling errors.
func Spellcheck(v string) weaver.Node {
	return weaver.Attr("spellcheck", v)
}

// Style sets the style attribute, containing CSS styling declarations to be applied to the element.
func Style(v string) weaver.Node {
	return weaver.Attr("style", v)
}

// TabIndex sets the tabindex attribute, indicating if the element can take input focus.
func TabIndex(v string) weaver.Node {
	return weaver.Attr("tabindex", v)
}

// Title sets the title attribute, containing text representing advisory information related to the element.
func Title(v string) weaver.Node {
	return weaver.Attr("title", v)
}

// Translate sets the translate attribute, specifying whether an element's attribute values and direct text child nodes should be translated.
func Translate(v string) weaver.Node {
	return weaver.Attr("translate", v)
}

// Href sets the href attribute, specifying the URL of a linked resource.
func Href(v string) weaver.Node {
	return weaver.Attr("href", v)
}

// HrefLang sets the hreflang attribute, specifying the human language of the linked resource.
func HrefLang(v string) weaver.Node {
	return weaver.Attr("hreflang", v)
}

// Rel sets the rel attribute, specifying the relationship of the linked document to the current document.
func Rel(v string) weaver.Node {
	return weaver.Attr("rel", v)
}

// Target sets the target attribute, specifying where to display the linked URL.
func Target(v string) weaver.Node {
	return weaver.Attr("target", v)
}

// Download sets the download attribute, instructing the browser to download the URL instead of navigating to it.
func Download(v string) weaver.Node {
	return weaver.Attr("download", v)
}

// Ping sets the ping attribute, specifying a space-separated list of URLs to be notified if the user follows the hyperlink.
func Ping(v string) weaver.Node {
	return weaver.Attr("ping", v)
}

// ReferrerPolicy sets the referrerpolicy attribute, specifying which referrer to use when fetching the resource.
func ReferrerPolicy(v string) weaver.Node {
	return weaver.Attr("referrerpolicy", v)
}

// Media sets the media attribute, specifying the media type that the linked resource applies to.
func Media(v string) weaver.Node {
	return weaver.Attr("media", v)
}

// Src sets the src attribute, specifying the URL of an external resource to embed.
func Src(v string) weaver.Node {
	return weaver.Attr("src", v)
}

// Alt sets the alt attribute, defining an alternative text description for the element.
func Alt(v string) weaver.Node {
	return weaver.Attr("alt", v)
}

// Width sets the width attribute, specifying the intrinsic width of the element in pixels.
func Width(v string) weaver.Node {
	return weaver.Attr("width", v)
}

// Height sets the height attribute, specifying the intrinsic height of the element in pixels.
func Height(v string) weaver.Node {
	return weaver.Attr("height", v)
}

// Loading sets the loading attribute, indicating how the browser should load the image/iframe.
func Loading(v string) weaver.Node {
	return weaver.Attr("loading", v)
}

// Decoding sets the decoding attribute, providing a hint on how the browser should decode the image.
func Decoding(v string) weaver.Node {
	return weaver.Attr("decoding", v)
}

// CrossOrigin sets the crossorigin attribute, indicating whether the fetch of the resource should be done using a CORS request.
func CrossOrigin(v string) weaver.Node {
	return weaver.Attr("crossorigin", v)
}

// SrcSet sets the srcset attribute, defining a list of image candidates for varying display densities or screens.
func SrcSet(v string) weaver.Node {
	return weaver.Attr("srcset", v)
}

// Sizes sets the sizes attribute, defining source sizes for responsive image selection.
func Sizes(v string) weaver.Node {
	return weaver.Attr("sizes", v)
}

// Poster sets the poster attribute, specifying an image to be shown while the video is downloading.
func Poster(v string) weaver.Node {
	return weaver.Attr("poster", v)
}

// Preload sets the preload attribute, providing a hint to the browser about how the media should be loaded.
func Preload(v string) weaver.Node {
	return weaver.Attr("preload", v)
}

// Autoplay sets the autoplay attribute, specifying that the media should automatically start playing.
func Autoplay(enabled bool) weaver.Node {
	return weaver.Flag("autoplay", true)
}

// Controls sets the controls attribute, displaying playback controls for the media element.
func Controls(enabled bool) weaver.Node {
	return weaver.Flag("controls", true)
}

// Loop sets the loop attribute, causing the media element to start over when reaching the end.
func Loop(v string) weaver.Node {
	return weaver.Attr("loop", v)
}

// Muted sets the muted attribute, indicating that the audio output of the media should be silenced.
func Muted(enabled bool) weaver.Node {
	return weaver.Flag("muted", true)
}

// PlaysInline sets the playsinline attribute, specifying that video should play inline within the element's playback area.
func PlaysInline(enabled bool) weaver.Node {
	return weaver.Flag("playsinline", true)
}

// Sandbox sets the sandbox attribute, applying extra restrictions to the content in the iframe.
func Sandbox(v string) weaver.Node {
	return weaver.Attr("sandbox", v)
}

// SrcDoc sets the srcdoc attribute, specifying the HTML content of the page to show in the inline frame.
func SrcDoc(v string) weaver.Node {
	return weaver.Attr("srcdoc", v)
}

// Name sets the name attribute, defining the name of the element.
func Name(v string) weaver.Node {
	return weaver.Attr("name", v)
}

// Value sets the value attribute, defining the initial value of the element.
func Value(v string) weaver.Node {
	return weaver.Attr("value", v)
}

// Type sets the type attribute, specifying the type of element or input control.
func Type(v string) weaver.Node {
	return weaver.Attr("type", v)
}

// Placeholder sets the placeholder attribute, giving a hint to the user of what can be entered in the control.
func Placeholder(v string) weaver.Node {
	return weaver.Attr("placeholder", v)
}

// Disabled sets the disabled attribute, indicating that the form control is disabled.
func Disabled(enabled bool) weaver.Node {
	return weaver.Flag("disabled", true)
}

// ReadOnly sets the readonly attribute, indicating that the form control cannot be edited by the user.
func ReadOnly(enabled bool) weaver.Node {
	return weaver.Flag("readonly", true)
}

// Required sets the required attribute, indicating that the user must specify a value for the input before submitting the form.
func Required(enabled bool) weaver.Node {
	return weaver.Flag("required", true)
}

// Checked sets the checked attribute, indicating that the checkbox or radio control is checked.
func Checked(enabled bool) weaver.Node {
	return weaver.Flag("checked", true)
}

// Selected sets the selected attribute, indicating that the option is currently selected.
func Selected(enabled bool) weaver.Node {
	return weaver.Flag("selected", true)
}

// Multiple sets the multiple attribute, indicating that multiple values can be entered/selected.
func Multiple(enabled bool) weaver.Node {
	return weaver.Flag("multiple", true)
}

// AutoComplete sets the autocomplete attribute, specifying whether the control's value can be automatically completed by the browser.
func AutoComplete(v string) weaver.Node {
	return weaver.Attr("autocomplete", v)
}

// Accept sets the accept attribute, defining the types of files that the server accepts.
func Accept(v string) weaver.Node {
	return weaver.Attr("accept", v)
}

// AcceptCharset sets the accept-charset attribute, specifying the character encodings used for form submission.
func AcceptCharset(v string) weaver.Node {
	return weaver.Attr("accept-charset", v)
}

// Action sets the action attribute, defining the URL that processes the form submission.
func Action(v string) weaver.Node {
	return weaver.Attr("action", v)
}

// Method sets the method attribute, specifying the HTTP method to use when submitting the form.
func Method(v string) weaver.Node {
	return weaver.Attr("method", v)
}

// Enctype sets the enctype attribute, defining the encoding type used when submitting the form.
func Enctype(v string) weaver.Node {
	return weaver.Attr("enctype", v)
}

// NoValidate sets the novalidate attribute, indicating that the form should not be validated when submitted.
func NoValidate(enabled bool) weaver.Node {
	return weaver.Flag("novalidate", true)
}

// Form sets the form attribute, associating the element with a form element.
func Form(v string) weaver.Node {
	return weaver.Attr("form", v)
}

// FormAction sets the formaction attribute, specifying the URL that processes form submission for this element.
func FormAction(v string) weaver.Node {
	return weaver.Attr("formaction", v)
}

// FormEnctype sets the formenctype attribute, specifying the encoding type for form submission for this element.
func FormEnctype(v string) weaver.Node {
	return weaver.Attr("formenctype", v)
}

// FormMethod sets the formmethod attribute, specifying the HTTP method for form submission for this element.
func FormMethod(v string) weaver.Node {
	return weaver.Attr("formmethod", v)
}

// FormNoValidate sets the formnovalidate attribute, indicating that the form shouldn't be validated when submitted by this element.
func FormNoValidate(enabled bool) weaver.Node {
	return weaver.Flag("formnovalidate", true)
}

// FormTarget sets the formtarget attribute, specifying where to display the response after submitting the form for this element.
func FormTarget(v string) weaver.Node {
	return weaver.Attr("formtarget", v)
}

// Capture sets the capture attribute, specifying which camera or microphone to use for media capture.
func Capture(v string) weaver.Node {
	return weaver.Attr("capture", v)
}

// Pattern sets the pattern attribute, specifying a regular expression that the control's value must match.
func Pattern(v string) weaver.Node {
	return weaver.Attr("pattern", v)
}

// MinLength sets the minlength attribute, defining the minimum number of characters allowed.
func MinLength(v string) weaver.Node {
	return weaver.Attr("minlength", v)
}

// MaxLength sets the maxlength attribute, defining the maximum number of characters allowed.
func MaxLength(v string) weaver.Node {
	return weaver.Attr("maxlength", v)
}

// Size sets the size attribute, defining the initial size or width of the control.
func Size(v string) weaver.Node {
	return weaver.Attr("size", v)
}

// Cols sets the cols attribute, specifying the visible width of a text area in average character widths.
func Cols(v string) weaver.Node {
	return weaver.Attr("cols", v)
}

// Rows sets the rows attribute, specifying the visible height of a text area in lines.
func Rows(v string) weaver.Node {
	return weaver.Attr("rows", v)
}

// Wrap sets the wrap attribute, indicating how text in a text area should be wrapped upon submission.
func Wrap(v string) weaver.Node {
	return weaver.Attr("wrap", v)
}

// Min sets the min attribute, defining the minimum value allowed for the input control.
func Min(v string) weaver.Node {
	return weaver.Attr("min", v)
}

// Max sets the max attribute, defining the maximum value allowed for the input control.
func Max(v string) weaver.Node {
	return weaver.Attr("max", v)
}

// Step sets the step attribute, defining the legal number intervals for an input field.
func Step(v string) weaver.Node {
	return weaver.Attr("step", v)
}

// List sets the list attribute, identifying the datalist element containing suggested values.
func List(v string) weaver.Node {
	return weaver.Attr("list", v)
}

// High sets the high attribute, defining the lower bound of the high range for a meter element.
func High(v string) weaver.Node {
	return weaver.Attr("high", v)
}

// Low sets the low attribute, defining the upper bound of the low range for a meter element.
func Low(v string) weaver.Node {
	return weaver.Attr("low", v)
}

// Optimum sets the optimum attribute, defining the optimal numeric value for a meter element.
func Optimum(v string) weaver.Node {
	return weaver.Attr("optimum", v)
}

// For sets the for attribute, associating the element with another element's ID.
func For(v string) weaver.Node {
	return weaver.Attr("for", v)
}

// ColSpan sets the colspan attribute, defining the number of columns a cell should span.
func ColSpan(v string) weaver.Node {
	return weaver.Attr("colspan", v)
}

// RowSpan sets the rowspan attribute, defining the number of rows a cell should span.
func RowSpan(v string) weaver.Node {
	return weaver.Attr("rowspan", v)
}

// Scope sets the scope attribute, defining the cells that the header element relates to.
func Scope(v string) weaver.Node {
	return weaver.Attr("scope", v)
}

// Headers sets the headers attribute, establishing a relationship between table cells and headers.
func Headers(v string) weaver.Node {
	return weaver.Attr("headers", v)
}

// Charset sets the charset attribute, declaring the character encoding of the document.
func Charset(v string) weaver.Node {
	return weaver.Attr("charset", v)
}

// Content sets the content attribute, giving the value associated with the http-equiv or name attribute.
func Content(v string) weaver.Node {
	return weaver.Attr("content", v)
}

// HTTPEquiv sets the http-equiv attribute, defining a pragma directive for processing the document.
func HTTPEquiv(v string) weaver.Node {
	return weaver.Attr("http-equiv", v)
}

// Async sets the async attribute, indicating that the script should be executed asynchronously.
func Async(enabled bool) weaver.Node {
	return weaver.Flag("async", true)
}

// Defer sets the defer attribute, indicating that the script should be executed after the document has been parsed.
func Defer(enabled bool) weaver.Node {
	return weaver.Flag("defer", true)
}

// NoModule sets the nomodule attribute, preventing the script from executing in browsers that support ES modules.
func NoModule(v string) weaver.Node {
	return weaver.Attr("nomodule", v)
}

// Integrity sets the integrity attribute, containing inline metadata used to verify the fetched resource.
func Integrity(v string) weaver.Node {
	return weaver.Attr("integrity", v)
}

// Cite sets the cite attribute, specifying a URL that points to the source of a quotation or edit.
func Cite(v string) weaver.Node {
	return weaver.Attr("cite", v)
}

// DateTime sets the datetime attribute, defining a machine-readable date/time string.
func DateTime(v string) weaver.Node {
	return weaver.Attr("datetime", v)
}

// Open sets the open attribute, indicating that the disclosure widget is currently visible.
func Open(enabled bool) weaver.Node {
	return weaver.Flag("open", true)
}

// Reversed sets the reversed attribute, specifying that list items should be numbered in reverse order.
func Reversed(enabled bool) weaver.Node {
	return weaver.Flag("reversed", true)
}

// Start sets the start attribute, defining the starting value of an ordered list.
func Start(v string) weaver.Node {
	return weaver.Attr("start", v)
}

// IsMap sets the ismap attribute, indicating that the image is a server-side image map.
func IsMap(v string) weaver.Node {
	return weaver.Attr("ismap", v)
}

// UseMap sets the usemap attribute, associating the image with a client-side image map.
func UseMap(v string) weaver.Node {
	return weaver.Attr("usemap", v)
}

// Coords sets the coords attribute, defining the coordinates of an image map region.
func Coords(v string) weaver.Node {
	return weaver.Attr("coords", v)
}

// Shape sets the shape attribute, defining the shape of an image map region.
func Shape(v string) weaver.Node {
	return weaver.Attr("shape", v)
}

// Span sets the span attribute, specifying the number of columns/groups an element spans.
func Span(v string) weaver.Node {
	return weaver.Attr("span", v)
}

// Default sets the default attribute, specifying that a track is to be enabled if user preferences don't indicate otherwise.
func Default(enabled bool) weaver.Node {
	return weaver.Flag("default", true)
}

// Kind sets the kind attribute, specifying the kind of text track.
func Kind(v string) weaver.Node {
	return weaver.Attr("kind", v)
}

// Label sets the label attribute, specifying a user-readable title for the track or option group.
func Label(v string) weaver.Node {
	return weaver.Attr("label", v)
}

// SrcLang sets the srclang attribute, specifying the language of the track text data.
func SrcLang(v string) weaver.Node {
	return weaver.Attr("srclang", v)
}

// Data sets a custom data attribute (`data-*`).
func Data(key, v string) weaver.Node {
	return weaver.Attr("data-"+key, v)
}
