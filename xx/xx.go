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

// C appends a child element.
//
// Example usage:
//
//	div := xx.E("div").C(xx.E("span").VAL("Hello"))
//	fmt.Println(div.resolve()) // Outputs: <div><span>Hello</span></div>
func (el Elem) C(child Elem) Elem {
	el.children = append(el.children, child)
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
