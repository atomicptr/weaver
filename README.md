# Weaver - Simple and efficient HTML DSL for Go!

Weaver is a lightweight HTML domain specific language (DSL) for Go. It allows you to build clean, maintable HTML
templates without template engines, context switching or runtime parsing overhead. Weaver renders fast and writes
directly into any `io.Writer`.

## Installation

```bash
$ go get atomicptr.dev/weaver
```

## Example Usage

```go
package main

import (
	"log"
	"net/http"

	"atomicptr.dev/weaver/html/attr"
	"atomicptr.dev/weaver/html/elem"
)

func main() {
	log.Printf("listening on 127.0.0.1:8080...")

	err := http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := elem.Html(
			attr.Data("theme", "dark"),
			attr.Lang("en"),

			elem.Head(
				elem.Title("Weaver - Simple and efficient DSL for Go!"),
				elem.Link(
					attr.Rel("stylesheet"),
					attr.Type("text/css"),
					attr.Href("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css"),
				),
			),

			elem.Body(
				elem.Main(
					attr.Class("container"),

					elem.Div(
						attr.Style("margin-bottom: 2rem;"),

						elem.H1(elem.Text("Weaver")),
						elem.H2(elem.Text("Simple and efficient DSL for Go!")),
						elem.A(
							attr.Href("https://atomicptr.dev/weaver"),
							attr.Target("_blank"),

							elem.Text("Learn More!"),
						),
					),

					elem.Div(
						elem.Img(attr.Src("https://go.dev/images/gophers/ladder.svg")),
					),
				),
			),
		)

		err := page.RenderHtml(w) // this can be any io.Writer
		if err != nil {
			log.Fatal(err)
		}
	}))
	if err != nil {
		log.Fatal(err)
	}
}
```

Run the example via:

```bash
$ go run examples/main.go
```

## License

MIT
