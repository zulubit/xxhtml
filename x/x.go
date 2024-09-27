package x

import "html"

// Elem represents an HTML element with attributes, text, and children.
type Elem struct {
	element  string
	children []Elem
	tag      *X
	value    string
}

// X represents the properties of an element like value, class, id, and additional attributes.
type X struct {
	Class string
	Id    string
	Att   string
}

// Render generates the HTML representation of the element and its children as a byte slice.
//
// Example usage:
//
//	elem := xx.E("div", xx.X{class: "container", value: "Hello, World!"})
//	response := elem.Render()
//	fmt.Println(string(response)) // Outputs: <div class="container">Hello, World!</div>
func (tr Elem) Render() []byte {
	return []byte(tr.resolve())
}

// resolve constructs the HTML string for the element and recursively for its children.
func (tr Elem) resolve() string {
	if tr.element == "" {
		// If the element is raw text, return it directly along with any children.
		chiraw := ""
		for _, c := range tr.children {
			chiraw += c.resolve()
		}
		return tr.value + chiraw
	}

	// Start with the opening tag.
	elp1 := "<" + tr.element

	// Add attributes from `X` struct if present.
	if tr.tag != nil {
		if tr.tag.Class != "" {
			elp1 += " class=\"" + tr.tag.Class + "\""
		}
		if tr.tag.Id != "" {
			elp1 += " id=\"" + tr.tag.Id + "\""
		}
		if tr.tag.Att != "" {
			elp1 += " " + tr.tag.Att
		}
	}

	// Close the opening tag.
	elp1 += ">"

	// Add element value.
	elp1 += tr.value

	// Recursively resolve children elements.
	che := ""
	for _, c := range tr.children {
		che += c.resolve()
	}

	// Construct the closing tag.
	elp2 := "</" + tr.element + ">"

	// Return the complete HTML string.
	return elp1 + che + elp2
}

// E initializes a new Elem with the specified tag name, attributes from X, and optional children.
//
// Example usage:
//
//	div := xx.E("div", xx.X{class: "container", value: "Hello, World!"}, xx.Span(xx.X{value: "Child 1"}))
//	fmt.Println(div.resolve()) // Outputs: <div class="container">Hello, World!<span>Child 1</span></div>
func E(name string, tag X, children ...Elem) Elem {
	// Initialize the element with the specified tag name
	return Elem{element: name, tag: &tag, children: children}

}

// ERAW creates an Elem with raw HTML content or plain text.
//
// Example usage:
//
//	raw := xx.ERAW("<h1>Hello, World!</h1>")
//	fmt.Println(raw.resolve()) // Outputs: <h1>Hello, World!</h1>
func ERAW(value string) Elem {
	el := E("", X{})
	el.value = value
	return el
}

// C creates an Elem with raw HTML content or plain text.
//
// Example usage:
//
//	raw := xx.C("<h1>Hello, World!</h1>")
//	fmt.Println(raw.resolve()) // Outputs: <h1>Hello, World!</h1>
func C(value string) Elem {
	value = html.EscapeString(value)
	return ERAW(value)
}

// IF returns trueCase if the condition is true, otherwise returns an empty Elem.
//
// Example usage:
//
//	elem := xx.IF(true, xx.E("span", xx.X{value: "True"}))
//	fmt.Println(elem.resolve()) // Outputs: <span>True</span>
func IF(condition bool, trueCase Elem) Elem {
	if condition {
		return trueCase
	}
	return Elem{}
}

// FOR takes a slice of Elem and returns all elements in the slice.
//
// Example usage:
//
//	elems := xx.FOR([]xx.Elem{
//	    xx.E("li", xx.X{value: "Item 1"}),
//	    xx.E("li", xx.X{value: "Item 2"}),
//	})
//	fmt.Println(elems[0].resolve()) // Outputs: <li>Item 1</li>
//	fmt.Println(elems[1].resolve()) // Outputs: <li>Item 2</li>
func FOR(iterClosure []Elem) []Elem {
	return iterClosure
}

// TER returns trueCase if the condition is true, otherwise returns falseCase.
//
// Example usage:
//
//	result := xx.TER(true, xx.E("p", xx.X{value: "True"}), xx.E("p", xx.X{value: "False"}))
//	fmt.Println(result.resolve()) // Outputs: <p>True</p>
func TER(condition bool, trueCase Elem, falseCase Elem) Elem {
	if condition {
		return trueCase
	}
	return falseCase
}

// STER returns trueCase if the boolean condition is true, otherwise returns falseCase.
//
// Example usage:
//
//	result := xx.STER(true, "True", "False")
//	fmt.Println(result) // Outputs: True
func STER(condition bool, trueCase string, falseCase string) string {
	if condition {
		return trueCase
	}
	return falseCase
}

// SIF returns trueCase if the boolean condition is true, otherwise returns an empty string.
//
// Example usage:
//
//	result := xx.SIF(true, "True")
//	fmt.Println(result) // Outputs: True
func SIF(condition bool, trueCase string) string {
	if condition {
		return trueCase
	}
	return ""
}

// Convenience functions

// Div creates a new <div> element with optional children.
//
// Example usage:
//
//	div := xx.Div(xx.X{class: "container", value: "Hello, World!"}, xx.Span(xx.X{value: "Child 1"}))
//	fmt.Println(div.resolve()) // Outputs: <div class="container">Hello, World!<span>Child 1</span></div>
func Div(x X, children ...Elem) Elem {
	return E("div", x, children...)
}

// Span creates a new <span> element with optional children.
//
// Example usage:
//
//	span := xx.Span(xx.X{value: "Some text"}, xx.Span(xx.X{value: "Child Span"}))
//	fmt.Println(span.resolve()) // Outputs: <span>Some text<span>Child Span</span></span>
func Span(x X, children ...Elem) Elem {
	return E("span", x, children...)
}

// P creates a new <p> (paragraph) element with optional children.
//
// Example usage:
//
//	p := xx.P(xx.X{value: "This is a paragraph."})
//	fmt.Println(p.resolve()) // Outputs: <p>This is a paragraph.</p>
func P(x X, children ...Elem) Elem {
	return E("p", x, children...)
}

// A creates a new <a> (anchor) element with optional children.
//
// Example usage:
//
//	a := xx.A(xx.X{att: `href="https://example.com"`, value: "Click here"}, xx.Span(xx.X{value: "Icon"}))
//	fmt.Println(a.resolve()) // Outputs: <a href="https://example.com">Click here<span>Icon</span></a>
func A(x X, children ...Elem) Elem {
	return E("a", x, children...)
}

// Img creates a new <img> (image) element.
//
// Example usage:
//
//	img := xx.Img(xx.X{att: `src="image.png" alt="Image description"`})
//	fmt.Println(img.resolve()) // Outputs: <img src="image.png" alt="Image description"></img>
func Img(x X) Elem {
	return E("img", x)
}

// H1 creates a new <h1> (heading) element with optional children.
//
// Example usage:
//
//	h1 := xx.H1(xx.X{value: "Heading 1"})
//	fmt.Println(h1.resolve()) // Outputs: <h1>Heading 1</h1>
func H1(x X, children ...Elem) Elem {
	return E("h1", x, children...)
}

// H2 creates a new <h2> (heading) element with optional children.
//
// Example usage:
//
//	h2 := xx.H2(xx.X{value: "Heading 2"})
//	fmt.Println(h2.resolve()) // Outputs: <h2>Heading 2</h2>
func H2(x X, children ...Elem) Elem {
	return E("h2", x, children...)
}

// H3 creates a new <h3> (heading) element with optional children.
//
// Example usage:
//
//	h3 := xx.H3(xx.X{value: "Heading 3"})
//	fmt.Println(h3.resolve()) // Outputs: <h3>Heading 3</h3>
func H3(x X, children ...Elem) Elem {
	return E("h3", x, children...)
}

// Ul creates a new <ul> (unordered list) element with optional children.
//
// Example usage:
//
//	ul := xx.Ul(xx.X{}, xx.Li(xx.X{value: "Item 1"}), xx.Li(xx.X{value: "Item 2"}))
//	fmt.Println(ul.resolve()) // Outputs: <ul><li>Item 1</li><li>Item 2</li></ul>
func Ul(x X, children ...Elem) Elem {
	return E("ul", x, children...)
}

// Ol creates a new <ol> (ordered list) element with optional children.
//
// Example usage:
//
//	ol := xx.Ol(xx.X{}, xx.Li(xx.X{value: "First"}), xx.Li(xx.X{value: "Second"}))
//	fmt.Println(ol.resolve()) // Outputs: <ol><li>First</li><li>Second</li></ol>
func Ol(x X, children ...Elem) Elem {
	return E("ol", x, children...)
}

// Li creates a new <li> (list item) element.
//
// Example usage:
//
//	li := xx.Li(xx.X{value: "List item"})
//	fmt.Println(li.resolve()) // Outputs: <li>List item</li>
func Li(x X) Elem {
	return E("li", x)
}

// Table creates a new <table> element with optional children.
//
// Example usage:
//
//	table := xx.Table(xx.X{}, xx.Tr(xx.X{}, xx.Td(xx.X{value: "Cell 1"}), xx.Td(xx.X{value: "Cell 2"})))
//	fmt.Println(table.resolve()) // Outputs: <table><tr><td>Cell 1</td><td>Cell 2</td></tr></table>
func Table(x X, children ...Elem) Elem {
	return E("table", x, children...)
}

// Tr creates a new <tr> (table row) element with optional children.
//
// Example usage:
//
//	tr := xx.Tr(xx.X{}, xx.Td(xx.X{value: "Cell 1"}), xx.Td(xx.X{value: "Cell 2"}))
//	fmt.Println(tr.resolve()) // Outputs: <tr><td>Cell 1</td><td>Cell 2</td></tr>
func Tr(x X, children ...Elem) Elem {
	return E("tr", x, children...)
}

// Td creates a new <td> (table cell) element.
//
// Example usage:
//
//	td := xx.Td(xx.X{value: "Cell content"})
//	fmt.Println(td.resolve()) // Outputs: <td>Cell content</td>
func Td(x X) Elem {
	return E("td", x)
}

// Th creates a new <th> (table header cell) element.
//
// Example usage:
//
//	th := xx.Th(xx.X{value: "Header"})
//	fmt.Println(th.resolve()) // Outputs: <th>Header</th>
func Th(x X) Elem {
	return E("th", x)
}

// Form creates a new <form> element with optional children.
//
// Example usage:
//
//	form := xx.Form(xx.X{att: `action="/submit" method="post"`}, xx.Input(xx.X{att: `type="text"`}))
//	fmt.Println(form.resolve()) // Outputs: <form action="/submit" method="post"><input type="text"></form>
func Form(x X, children ...Elem) Elem {
	return E("form", x, children...)
}

// Input creates a new <input> element.
//
// Example usage:
//
//	input := xx.Input(xx.X{att: `type="text" placeholder="Enter text"`})
//	fmt.Println(input.resolve()) // Outputs: <input type="text" placeholder="Enter text"></input>
func Input(x X) Elem {
	return E("input", x)
}

// Button creates a new <button> element with optional children.
//
// Example usage:
//
//	button := xx.Button(xx.X{value: "Click me"})
//	fmt.Println(button.resolve()) // Outputs: <button>Click me</button>
func Button(x X, children ...Elem) Elem {
	return E("button", x, children...)
}

// Label creates a new <label> element with optional children.
//
// Example usage:
//
//	label := xx.Label(xx.X{att: `for="inputID"`, value: "Label text"})
//	fmt.Println(label.resolve()) // Outputs: <label for="inputID">Label text</label>
func Label(x X, children ...Elem) Elem {
	return E("label", x, children...)
}

// Article creates a new <article> element with optional children.
//
// Example usage:
//
//	article := xx.Article(xx.X{value: "Article content"}, xx.Section(xx.X{value: "Section"}))
//	fmt.Println(article.resolve()) // Outputs: <article>Article content<section>Section</section></article>
func Article(x X, children ...Elem) Elem {
	return E("article", x, children...)
}

// Aside creates a new <aside> element with optional children.
//
// Example usage:
//
//	aside := xx.Aside(xx.X{value: "Aside content"})
//	fmt.Println(aside.resolve()) // Outputs: <aside>Aside content</aside>
func Aside(x X, children ...Elem) Elem {
	return E("aside", x, children...)
}

// Header creates a new <header> element with optional children.
//
// Example usage:
//
//	header := xx.Header(xx.X{value: "Header content"}, xx.Nav(xx.X{value: "Navigation"}))
//	fmt.Println(header.resolve()) // Outputs: <header>Header content<nav>Navigation</nav></header>
func Header(x X, children ...Elem) Elem {
	return E("header", x, children...)
}

// Footer creates a new <footer> element with optional children.
//
// Example usage:
//
//	footer := xx.Footer(xx.X{value: "Footer content"})
//	fmt.Println(footer.resolve()) // Outputs: <footer>Footer content</footer>
func Footer(x X, children ...Elem) Elem {
	return E("footer", x, children...)
}

// Main creates a new <main> element with optional children.
//
// Example usage:
//
//	main := xx.Main(xx.X{value: "Main content"})
//	fmt.Println(main.resolve()) // Outputs: <main>Main content</main>
func Main(x X, children ...Elem) Elem {
	return E("main", x, children...)
}

// Section creates a new <section> element with optional children.
//
// Example usage:
//
//	section := xx.Section(xx.X{value: "Section content"})
//	fmt.Println(section.resolve()) // Outputs: <section>Section content</section>
func Section(x X, children ...Elem) Elem {
	return E("section", x, children...)
}

// Nav creates a new <nav> element with optional children.
//
// Example usage:
//
//	nav := xx.Nav(xx.X{value: "Navigation links"})
//	fmt.Println(nav.resolve()) // Outputs: <nav>Navigation links</nav>
func Nav(x X, children ...Elem) Elem {
	return E("nav", x, children...)
}

// Figure creates a new <figure> element with optional children.
//
// Example usage:
//
//	figure := xx.Figure(xx.X{value: "Figure content"}, xx.Figcaption(xx.X{value: "Caption"}))
//	fmt.Println(figure.resolve()) // Outputs: <figure>Figure content<figcaption>Caption</figcaption></figure>
func Figure(x X, children ...Elem) Elem {
	return E("figure", x, children...)
}

// Figcaption creates a new <figcaption> element with optional children.
//
// Example usage:
//
//	figcaption := xx.Figcaption(xx.X{value: "Caption text"})
//	fmt.Println(figcaption.resolve()) // Outputs: <figcaption>Caption text</figcaption>
func Figcaption(x X, children ...Elem) Elem {
	return E("figcaption", x, children...)
}

// Datalist creates a new <datalist> element with optional children.
//
// Example usage:
//
//	datalist := xx.Datalist(xx.Option(xx.X{value: "Option 1"}), xx.Option(xx.X{value: "Option 2"}))
//	fmt.Println(datalist.resolve()) // Outputs: <datalist><option>Option 1</option><option>Option 2</option></datalist>
func Datalist(x X, children ...Elem) Elem {
	return E("datalist", x, children...)
}

// Option creates a new <option> element with optional children.
//
// Example usage:
//
//	option := xx.Option(xx.X{value: "Option text"})
//	fmt.Println(option.resolve()) // Outputs: <option>Option text</option>
func Option(x X) Elem {
	return E("option", x)
}

// Details creates a new <details> element with optional children.
//
// Example usage:
//
//	details := xx.Details(xx.X{value: "Details content"}, xx.Summary(xx.X{value: "Summary"}))
//	fmt.Println(details.resolve()) // Outputs: <details>Details content<summary>Summary</summary></details>
func Details(x X, children ...Elem) Elem {
	return E("details", x, children...)
}

// Summary creates a new <summary> element with optional children.
//
// Example usage:
//
//	summary := xx.Summary(xx.X{value: "Summary text"})
//	fmt.Println(summary.resolve()) // Outputs: <summary>Summary text</summary>
func Summary(x X, children ...Elem) Elem {
	return E("summary", x, children...)
}

// Dialog creates a new <dialog> element with optional children.
//
// Example usage:
//
//	dialog := xx.Dialog(xx.X{value: "Dialog content"})
//	fmt.Println(dialog.resolve()) // Outputs: <dialog>Dialog content</dialog>
func Dialog(x X, children ...Elem) Elem {
	return E("dialog", x, children...)
}

// Embed creates a new <embed> element with optional attributes.
//
// Example usage:
//
//	embed := xx.Embed(xx.X{att: `src="video.mp4" type="video/mp4"`})
//	fmt.Println(embed.resolve()) // Outputs: <embed src="video.mp4" type="video/mp4">
func Embed(x X) Elem {
	return E("embed", x)
}

// Map creates a new <map> element with optional children.
//
// Example usage:
//
//	mapElem := xx.Map(xx.X{value: "Map content"}, xx.Area(xx.X{att: `shape="rect" coords="34,44,270,350" href="https://example.com"`, value: "Area"}))
//	fmt.Println(mapElem.resolve()) // Outputs: <map>Map content<area shape="rect" coords="34,44,270,350" href="https://example.com">Area</area></map>
func Map(x X, children ...Elem) Elem {
	return E("map", x, children...)
}

// Area creates a new <area> element with optional attributes.
//
// Example usage:
//
//	area := xx.Area(xx.X{att: `shape="rect" coords="34,44,270,350" href="https://example.com"`, value: "Area"})
//	fmt.Println(area.resolve()) // Outputs: <area shape="rect" coords="34,44,270,350" href="https://example.com">Area</area>
func Area(x X) Elem {
	return E("area", x)
}

// Source creates a new <source> element with optional attributes.
//
// Example usage:
//
//	source := xx.Source(xx.X{att: `src="audio.mp3" type="audio/mp3"`})
//	fmt.Println(source.resolve()) // Outputs: <source src="audio.mp3" type="audio/mp3">
func Source(x X) Elem {
	return E("source", x)
}

// Track creates a new <track> element with optional attributes.
//
// Example usage:
//
//	track := xx.Track(xx.X{att: `src="captions.vtt" kind="subtitles" srclang="en" label="English"`})
//	fmt.Println(track.resolve()) // Outputs: <track src="captions.vtt" kind="subtitles" srclang="en" label="English">
func Track(x X) Elem {
	return E("track", x)
}

// Param creates a new <param> element with optional attributes.
//
// Example usage:
//
//	param := xx.Param(xx.X{att: `name="autoplay" value="true"`})
//	fmt.Println(param.resolve()) // Outputs: <param name="autoplay" value="true">
func Param(x X) Elem {
	return E("param", x)
}

// Script creates a new <script> element with optional attributes.
//
// Example usage:
//
//	script := xx.Script(xx.X{att: `src="script.js"`})
//	fmt.Println(script.resolve()) // Outputs: <script src="script.js"></script>
func Script(x X) Elem {
	return E("script", x)
}

// Style creates a new <style> element with optional children.
//
// Example usage:
//
//	style := xx.Style(xx.X{}, xx.ERAW("body { background: #f00; }"))
//	fmt.Println(style.resolve()) // Outputs: <style>body { background: #f00; }</style>
func Style(x X, children ...Elem) Elem {
	return E("style", x, children...)
}

// Meta creates a new <meta> element with optional attributes.
//
// Example usage:
//
//	meta := xx.Meta(xx.X{att: `name="viewport" content="width=device-width, initial-scale=1"`})
//	fmt.Println(meta.resolve()) // Outputs: <meta name="viewport" content="width=device-width, initial-scale=1">
func Meta(x X) Elem {
	return E("meta", x)
}

// Link creates a new <link> element with optional attributes.
//
// Example usage:
//
//	link := xx.Link(xx.X{att: `rel="stylesheet" href="styles.css"`})
//	fmt.Println(link.resolve()) // Outputs: <link rel="stylesheet" href="styles.css">
func Link(x X) Elem {
	return E("link", x)
}

// Title creates a new <title> element with optional children.
//
// Example usage:
//
//	title := xx.Title(xx.X{value: "Page Title"})
//	fmt.Println(title.resolve()) // Outputs: <title>Page Title</title>
func Title(x X, children ...Elem) Elem {
	return E("title", x, children...)
}

// Base creates a new <base> element with optional attributes.
//
// Example usage:
//
//	base := xx.Base(xx.X{att: `href="https://example.com/"`})
//	fmt.Println(base.resolve()) // Outputs: <base href="https://example.com/">
func Base(x X) Elem {
	return E("base", x)
}
