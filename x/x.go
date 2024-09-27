/*
Package x provides a simple API for creating and rendering HTML elements in Go.

Example usage:

	import (
		"fmt"
		"x" // Replace with the actual import path
	)

	func main() {
		elem := x.E("div", `class="container" id="main"`, x.C("Hello, World!"))
		response := elem.Render()
		fmt.Println(string(response)) // Outputs: <div class="container" id="main">Hello, World!</div>
	}
*/
package x

import (
	"bytes"
	"fmt"
	"html"
)

// Elem represents an HTML element with attributes, text, and children.
type Elem struct {
	element     string
	children    *[]Elem
	attributes  *string
	value       *string
	selfClosing bool
}

// Render generates the HTML representation of the element and its children as a byte slice.
func (e Elem) Render() []byte {
	return e.resolve()
}

// resolve constructs the HTML string for the element and recursively for its children.
func (e Elem) resolve() []byte {
	var buf bytes.Buffer

	attributes := e.attributes
	children := e.children
	value := e.value

	if e.element == "" {
		if value != nil {
			buf.WriteString(*value)
		}
		if children != nil {
			for _, c := range *children {
				buf.Write(c.resolve())
			}
		}
		return buf.Bytes()
	}

	buf.WriteString("<")
	buf.WriteString(e.element)

	if attributes != nil {
		buf.WriteString(" ")
		buf.WriteString(*attributes)
	}

	if e.selfClosing {
		buf.WriteString(" />")
		return buf.Bytes()
	}

	buf.WriteString(">")

	if value != nil {
		buf.WriteString(*value)
	}

	if children != nil {
		for _, c := range *children {
			buf.Write(c.resolve())
		}
	}

	buf.WriteString("</")
	buf.WriteString(e.element)
	buf.WriteString(">")

	return buf.Bytes()
}

// SC marks the element as self-closing.
func (e *Elem) SC() {
	e.selfClosing = true
}

// E initializes a new Elem with the specified tag name, attributes, and optional children.
func E(name string, attributes string, children ...Elem) Elem {
	var attr *string
	if attributes != "" {
		attr = &attributes
	}
	var childrenPtr *[]Elem
	if len(children) > 0 {
		childrenPtr = &children
	}
	return Elem{element: name, attributes: attr, children: childrenPtr}
}

// ERAW creates an Elem with raw HTML content or plain text.
func ERAW(value string) Elem {
	val := value
	return Elem{value: &val}
}

// C creates an Elem with escaped HTML content or plain text.
func C(value interface{}) Elem {
	switch v := value.(type) {
	case string:
		return ERAW(html.EscapeString(v))
	default:
		return ERAW(html.EscapeString(fmt.Sprintf("%v", value)))
	}
}

// CR creates an Elem with unescaped HTML content or plain text.
func CR(value interface{}) Elem {
	switch v := value.(type) {
	case string:
		return ERAW(v)
	default:
		return ERAW(fmt.Sprintf("%v", value))
	}
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
