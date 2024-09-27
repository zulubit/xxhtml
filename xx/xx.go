package xx

// Elem represents an HTML element with attributes, text, and children.
type Elem struct {
	element    string
	attributes string
	children   []Elem
	classes    string
	text       string
}

// Render generates the HTML representation of the element and its children as a byte slice.
//
// Example usage:
//
//	elem := xx.E("div").CLS("container").VAL("Hello, World!")
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
		return tr.text + chiraw
	}

	elp1 := "<" + tr.element
	if tr.attributes != "" {
		// Add attributes to the opening tag if present.
		elp1 += " " + tr.attributes
	}
	elp1 += ">" + tr.text

	// Recursively resolve children elements.
	che := ""
	for _, c := range tr.children {
		che += c.resolve()
	}

	// Construct the closing tag.
	elp2 := "</" + tr.element + ">"

	return elp1 + che + elp2
}

// E initializes a new Elem with the specified tag name.
//
// Example usage:
//
//	div := xx.E("div")
//	fmt.Println(div.resolve()) // Outputs: <div></div>
func E(name string) Elem {
	return Elem{element: name}
}

// ERAW creates an Elem with raw HTML content or plain text.
//
// Example usage:
//
//	raw := xx.ERAW("<h1>Hello, World!</h1>")
//	fmt.Println(raw.resolve()) // Outputs: <h1>Hello, World!</h1>
func ERAW(value string) Elem {
	return Elem{element: "", text: value, attributes: ""}
}

// CLS adds a class to the element.
//
// Example usage:
//
//	div := xx.E("div").CLS("container")
//	fmt.Println(div.resolve()) // Outputs: <div class="container"></div>
func (el Elem) CLS(class string) Elem {
	if el.classes != "" {
		el.classes += " " + class
	} else {
		el.classes = class
	}
	el.attributes += `class="` + el.classes + `"`
	return el
}

// VAL sets or appends the text content of the element.
//
// Example usage:
//
//	span := xx.E("span").VAL("Hello")
//	fmt.Println(span.resolve()) // Outputs: <span>Hello</span>
func (el Elem) VAL(value string) Elem {
	el.text += value
	return el
}

// ATT adds an attribute to the element.
//
// Example usage:
//
//	input := xx.E("input").ATT(`type="text"`).ATT(`placeholder="Enter text"`)
//	fmt.Println(input.resolve()) // Outputs: <input type="text" placeholder="Enter text"></input>
func (el Elem) ATT(value string) Elem {
	el.attributes += value + " "
	return el
}

// C appends one or more child elements.
//
// Example usage:
//
//	div := xx.E("div").C(xx.E("span").VAL("Hello"), xx.E("p").VAL("World"))
//	fmt.Println(div.resolve()) // Outputs: <div><span>Hello</span><p>World</p></div>
func (el Elem) C(children ...Elem) Elem {
	el.children = append(el.children, children...)
	return el
}

// IF returns trueCase if the condition is true, otherwise returns an empty Elem.
//
// Example usage:
//
//	elem := xx.IF(true, xx.E("span").VAL("True"))
//	fmt.Println(elem.resolve()) // Outputs: <span>True</span>
func IF(condition bool, trueCase Elem) Elem {
	if condition {
		return trueCase
	}
	return Elem{}
}

// FOR takes a slice of Elem and returns a parent Elem containing all elements in the slice as its children.
//
// Example usage:
//
//	elems := xx.FOR([]xx.Elem{
//	    xx.E("li").VAL("Item 1"),
//	    xx.E("li").VAL("Item 2"),
//	})
//	fmt.Println(elems.resolve()) // Outputs: <li>Item 1</li><li>Item 2</li>
func FOR(iterClosure []Elem) Elem {
	parent := Elem{}
	for _, child := range iterClosure {
		parent = parent.C(child)
	}
	return parent
}

// TER returns trueCase if the condition is true, otherwise returns falseCase.
//
// Example usage:
//
//	result := xx.TER(true, xx.E("p").VAL("True"), xx.E("p").VAL("False"))
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

// convinience functions

// Div creates a new <div> element.
//
// Example usage:
//   div := xx.Div().CLS("container").VAL("Hello, World!")
//   fmt.Println(div.resolve()) // Outputs: <div class="container">Hello, World!</div>

func Div() Elem {
	return E("div")
}

// Span creates a new <span> element.
//
// Example usage:
//
//	span := xx.Span().VAL("Some text")
//	fmt.Println(span.resolve()) // Outputs: <span>Some text</span>
func Span() Elem {
	return E("span")
}

// P creates a new <p> (paragraph) element.
//
// Example usage:
//
//	p := xx.P().VAL("This is a paragraph.")
//	fmt.Println(p.resolve()) // Outputs: <p>This is a paragraph.</p>
func P() Elem {
	return E("p")
}

// A creates a new <a> (anchor) element.
//
// Example usage:
//
//	a := xx.A().ATT(`href="https://example.com"`).VAL("Click here")
//	fmt.Println(a.resolve()) // Outputs: <a href="https://example.com">Click here</a>
func A() Elem {
	return E("a")
}

// Img creates a new <img> (image) element.
//
// Example usage:
//
//	img := xx.Img().ATT(`src="image.png"`).ATT(`alt="Image description"`)
//	fmt.Println(img.resolve()) // Outputs: <img src="image.png" alt="Image description"></img>
func Img() Elem {
	return E("img")
}

// H1 creates a new <h1> (heading) element.
//
// Example usage:
//
//	h1 := xx.H1().VAL("Heading 1")
//	fmt.Println(h1.resolve()) // Outputs: <h1>Heading 1</h1>
func H1() Elem {
	return E("h1")
}

// H2 creates a new <h2> (heading) element.
//
// Example usage:
//
//	h2 := xx.H2().VAL("Heading 2")
//	fmt.Println(h2.resolve()) // Outputs: <h2>Heading 2</h2>
func H2() Elem {
	return E("h2")
}

// H3 creates a new <h3> (heading) element.
//
// Example usage:
//
//	h3 := xx.H3().VAL("Heading 3")
//	fmt.Println(h3.resolve()) // Outputs: <h3>Heading 3</h3>
func H3() Elem {
	return E("h3")
}

// Ul creates a new <ul> (unordered list) element.
//
// Example usage:
//
//	ul := xx.Ul().C(xx.Li().VAL("Item 1"), xx.Li().VAL("Item 2"))
//	fmt.Println(ul.resolve()) // Outputs: <ul><li>Item 1</li><li>Item 2</li></ul>
func Ul() Elem {
	return E("ul")
}

// Ol creates a new <ol> (ordered list) element.
//
// Example usage:
//
//	ol := xx.Ol().C(xx.Li().VAL("First"), xx.Li().VAL("Second"))
//	fmt.Println(ol.resolve()) // Outputs: <ol><li>First</li><li>Second</li></ol>
func Ol() Elem {
	return E("ol")
}

// Li creates a new <li> (list item) element.
//
// Example usage:
//
//	li := xx.Li().VAL("List item")
//	fmt.Println(li.resolve()) // Outputs: <li>List item</li>
func Li() Elem {
	return E("li")
}

// Table creates a new <table> element.
//
// Example usage:
//
//	table := xx.Table()
//	fmt.Println(table.resolve()) // Outputs: <table></table>
func Table() Elem {
	return E("table")
}

// Tr creates a new <tr> (table row) element.
//
// Example usage:
//
//	tr := xx.Tr().C(xx.Td().VAL("Cell 1"), xx.Td().VAL("Cell 2"))
//	fmt.Println(tr.resolve()) // Outputs: <tr><td>Cell 1</td><td>Cell 2</td></tr>
func Tr() Elem {
	return E("tr")
}

// Td creates a new <td> (table cell) element.
//
// Example usage:
//
//	td := xx.Td().VAL("Cell content")
//	fmt.Println(td.resolve()) // Outputs: <td>Cell content</td>
func Td() Elem {
	return E("td")
}

// Th creates a new <th> (table header cell) element.
//
// Example usage:
//
//	th := xx.Th().VAL("Header")
//	fmt.Println(th.resolve()) // Outputs: <th>Header</th>
func Th() Elem {
	return E("th")
}

// Form creates a new <form> element.
//
// Example usage:
//
//	form := xx.Form().ATT(`action="/submit"`).ATT(`method="post"`)
//	fmt.Println(form.resolve()) // Outputs: <form action="/submit" method="post"></form>
func Form() Elem {
	return E("form")
}

// Input creates a new <input> element.
//
// Example usage:
//
//	input := xx.Input().ATT(`type="text"`).ATT(`placeholder="Enter text"`)
//	fmt.Println(input.resolve()) // Outputs: <input type="text" placeholder="Enter text"></input>
func Input() Elem {
	return E("input")
}

// Button creates a new <button> element.
//
// Example usage:
//
//	button := xx.Button().VAL("Click me")
//	fmt.Println(button.resolve()) // Outputs: <button>Click me</button>
func Button() Elem {
	return E("button")
}

// Label creates a new <label> element.
//
// Example usage:
//
//	label := xx.Label().ATT(`for="inputID"`).VAL("Label text")
//	fmt.Println(label.resolve()) // Outputs: <label for="inputID">Label text</label>
func Label() Elem {
	return E("label")
}
