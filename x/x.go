package x

import (
	"html"
	"strings"
)

// Elem represents an HTML element with attributes, text, and children.
type Elem struct {
	element     string
	children    []Elem
	attributes  string
	value       string
	selfClosing bool
}

// Render generates the HTML representation of the element and its children as a byte slice.
//
// Example usage:
//
//	elem := E("div", "class=\"container\" id=\"main\"", C("Hello, World!"))
//	response := elem.Render()
//	fmt.Println(string(response)) // Outputs: <div class="container" id="main">Hello, World!</div>
func (tr Elem) Render() []byte {
	return []byte(tr.resolve())
}

// resolve constructs the HTML string for the element and recursively for its children.
func (tr Elem) resolve() string {
	var sb strings.Builder

	if tr.element == "" {
		sb.WriteString(tr.value)
		for _, c := range tr.children {
			sb.WriteString(c.resolve())
		}
		return sb.String()
	}

	sb.WriteString("<")
	sb.WriteString(tr.element)

	if tr.attributes != "" {
		sb.WriteString(" ")
		sb.WriteString(tr.attributes)
	}

	if tr.selfClosing {
		sb.WriteString(" />")
		return sb.String()
	}

	sb.WriteString(">")
	sb.WriteString(tr.value)
	for _, c := range tr.children {
		sb.WriteString(c.resolve())
	}
	sb.WriteString("</")
	sb.WriteString(tr.element)
	sb.WriteString(">")

	return sb.String()
}

// E initializes a new Elem with the specified tag name, attributes, and optional children.
func E(name string, attributes string, children ...Elem) Elem {
	// Initialize the element with the specified tag name and attributes
	return Elem{element: name, attributes: attributes, children: children}
}

// ERAW creates an Elem with raw HTML content or plain text.
func ERAW(value string) Elem {
	return Elem{value: value}
}

// C creates an Elem with escaped HTML content or plain text.
func C(value string) Elem {
	value = html.EscapeString(value)
	return ERAW(value)
}

// IF returns trueCase if the condition is true, otherwise returns an empty Elem.
func IF(condition bool, trueCase Elem) Elem {
	if condition {
		return trueCase
	}
	return Elem{}
}

// FOR takes a slice of Elem and returns all elements in the slice.
func FOR(iterClosure []Elem) []Elem {
	return iterClosure
}

// TER returns trueCase if the condition is true, otherwise returns falseCase.
func TER(condition bool, trueCase Elem, falseCase Elem) Elem {
	if condition {
		return trueCase
	}
	return falseCase
}

// STER returns trueCase if the boolean condition is true, otherwise returns falseCase.
func STER(condition bool, trueCase string, falseCase string) string {
	if condition {
		return trueCase
	}
	return falseCase
}

// SIF returns trueCase if the boolean condition is true, otherwise returns an empty string.
func SIF(condition bool, trueCase string) string {
	if condition {
		return trueCase
	}
	return ""
}
