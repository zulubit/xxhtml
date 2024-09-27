// . Import recommended
package xxhtml

// Elem represents an HTML element with attributes, classes, text, and children.
type Elem struct {
	element     string
	attributes  string
	classes     string
	classstring string
	children    []Elem
	text        string
	rawFlag     bool
}

// CreateResponse generates the HTML response as a byte slice.
//
// Example usage:
//
//	elem := htmlbuilder.XX("div").Val("Hello, World!")
//	response := elem.CreateResponse()
//	fmt.Println(string(response)) // Outputs: <div>Hello, World!</div>
func (tr Elem) CreateResponse() []byte {
	return []byte(tr.resolve())
}

// resolve constructs the HTML string for the element and its children.
func (tr Elem) resolve() string {
	if tr.rawFlag {
		chiraw := ""
		for _, c := range tr.children {
			chiraw += c.resolve()
		}
		return tr.text + chiraw
	}

	elp1 := "<" + tr.element + tr.classstring + tr.attributes + ">" + tr.text
	che := ""
	for _, c := range tr.children {
		che += c.resolve()
	}
	elp2 := "</" + tr.element + ">"

	return elp1 + che + elp2
}

// XX initializes a new Elem with the specified tag name.
//
// Example usage:
//
//	div := htmlbuilder.XX("div").Val("Hello, World!")
//	fmt.Println(div.resolve()) // Outputs: <div>Hello, World!</div>
func XX(name string) Elem {
	return Elem{element: name}
}

// XIF returns trueCase if condition is true, otherwise returns falseCase.
//
// Example usage:
//
//	elem := htmlbuilder.XIF(true, htmlbuilder.XX("span").Val("True"), htmlbuilder.XX("span").Val("False"))
//	fmt.Println(elem.resolve()) // Outputs: <span>True</span>
func XIF(condition bool, trueCase Elem, falseCase Elem) Elem {
	if condition {
		return trueCase
	}
	if falseCase.attributes == "" && !falseCase.rawFlag {
		return XX("").Raw()
	}
	return falseCase
}

// XFOR returns an Elem containing all elements in iterClosure.
//
// Example usage:
//
//	elems := htmlbuilder.XFOR([]htmlbuilder.Elem{
//	    htmlbuilder.XX("li").Val("Item 1"),
//	    htmlbuilder.XX("li").Val("Item 2"),
//	})
//	fmt.Println(elems.resolve()) // Outputs: <li>Item 1</li><li>Item 2</li>
func XFOR(iterClosure []Elem) Elem {
	return XX("").Raw().Add(iterClosure...)
}

// Xter returns trueCase if condition is true, otherwise returns falseCase.
//
// Example usage:
//
//	result := htmlbuilder.Xter(true, "True", "False")
//	fmt.Println(result) // Outputs: True
func Xter(condition bool, trueCase string, falseCase string) string {
	if condition {
		return trueCase
	}
	return falseCase
}

// Att adds an attribute to the Elem.
//
// Example usage:
//
//	div := htmlbuilder.XX("div").Att(`id="main"`).Val("Content")
//	fmt.Println(div.resolve()) // Outputs: <div id="main">Content</div>
func (tr Elem) Att(s string) Elem {
	if tr.attributes == "" {
		tr.attributes = s
	} else {
		tr.attributes += " " + s
	}
	return tr
}

// Cls adds a class to the Elem.
//
// Example usage:
//
//	div := htmlbuilder.XX("div").Cls("container").Val("Content")
//	fmt.Println(div.resolve()) // Outputs: <div class="container">Content</div>
func (tr Elem) Cls(s string) Elem {
	if tr.classes == "" {
		tr.classes = s
	} else {
		tr.classes += " " + s
	}
	tr.classstring = ` class="` + tr.classes + `"`
	return tr
}

// Val sets the inner text of the Elem.
//
// Example usage:
//
//	div := htmlbuilder.XX("div").Val("Hello, World!")
//	fmt.Println(div.resolve()) // Outputs: <div>Hello, World!</div>
func (tr Elem) Val(s string) Elem {
	tr.text = s
	return tr
}

// Add appends children to the Elem.
//
// Example usage:
//
//	div := htmlbuilder.XX("div").Add(
//	    htmlbuilder.XX("span").Val("Child 1"),
//	    htmlbuilder.XX("span").Val("Child 2"),
//	)
//	fmt.Println(div.resolve()) // Outputs: <div><span>Child 1</span><span>Child 2</span></div>
func (tr Elem) Add(c ...Elem) Elem {
	tr.children = append(tr.children, c...)
	return tr
}

// Raw marks the Elem as raw HTML content.
//
// Example usage:
//
//	div := htmlbuilder.XX("div").Raw().Val(`<span>Raw Content</span>`)
//	fmt.Println(div.resolve()) // Outputs: <span>Raw Content</span>
func (tr Elem) Raw() Elem {
	tr.rawFlag = true
	return tr
}
