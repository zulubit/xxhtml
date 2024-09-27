package xx

// Elem represents an HTML element with attributes, text, and children.
type Elem struct {
	element    string
	attributes string
	children   []Elem
	text       string
}

// Render generates the HTML representation of the element and its children as a byte slice.
//
// Example usage:
//
//	elem := xx.E("div", `class="container"`, "Hello, World!")
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
	elp1 := ""
	if tr.attributes != "" {
		// Construct the opening tag with attributes and text content.
		elp1 = "<" + tr.element + " " + tr.attributes + ">" + tr.text
	} else {

		elp1 = "<" + tr.element + ">" + tr.text
	}

	// Recursively resolve children elements.
	che := ""
	for _, c := range tr.children {
		che += c.resolve()
	}

	// Construct the closing tag.
	elp2 := "</" + tr.element + ">"

	return elp1 + che + elp2
}

// E initializes a new Elem with the specified tag name, attributes, text content, and children elements.
//
// Example usage:
//
//	div := xx.E("div", `class="container"`, "Hello, World!")
//	fmt.Println(div.resolve()) // Outputs: <div class="container">Hello, World!</div>
func E(name string, attributes string, value string, children ...Elem) Elem {
	return Elem{element: name, text: value, attributes: attributes, children: children}
}

// ERAW creates a raw HTML element, often used for inserting raw HTML content or plain text.
//
// Example usage:
//
//	raw := xx.ERAW("<h1>Hello, World!</h1>")
//	fmt.Println(raw.resolve()) // Outputs: <h1>Hello, World!</h1>
func ERAW(value string) Elem {
	return Elem{element: "", text: value, attributes: ""}
}

// IF returns trueCase if the condition is true, otherwise returns an empty Elem.
//
// Example usage:
//
//	elem := xx.IF(true, xx.E("span", "", "True"))
//	fmt.Println(elem.resolve()) // Outputs: <span>True</span>
func IF(condition bool, trueCase Elem) Elem {
	if condition {
		return trueCase
	}
	return E("", "", "")
}

// FOR takes a slice of Elem and returns a parent Elem containing all elements in the slice as its children.
//
// Example usage:
//
//	elems := xx.FOR([]xx.Elem{
//	    xx.E("li", "", "Item 1"),
//	    xx.E("li", "", "Item 2"),
//	})
//	fmt.Println(elems.resolve()) // Outputs: <li>Item 1</li><li>Item 2</li>
func FOR(iterClosure []Elem) Elem {
	return E("", "", "", iterClosure...)
}

// TER returns trueCase if the condition is true, otherwise returns falseCase.
//
// Example usage:
//
//	result := xx.TER(true, xx.E("p", "", "True"), xx.E("p", "", "False"))
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
