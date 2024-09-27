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
	"fmt"
	"html"
	"strings"
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
	return []byte(e.resolve())
}

// resolve constructs the HTML string for the element and recursively for its children.
func (e Elem) resolve() string {
	var sb strings.Builder

	if e.element == "" {
		if e.value != nil {
			sb.WriteString(*e.value)
		}
		if e.children != nil {
			for _, c := range *e.children {
				sb.WriteString(c.resolve())
			}
		}
		return sb.String()
	}

	sb.WriteString("<")
	sb.WriteString(e.element)

	if e.attributes != nil {
		sb.WriteString(" ")
		sb.WriteString(*e.attributes)
	}

	if e.selfClosing {
		sb.WriteString(" />")
		return sb.String()
	}

	sb.WriteString(">")
	if e.value != nil {
		sb.WriteString(*e.value)
	}
	if e.children != nil {
		for _, c := range *e.children {
			sb.WriteString(c.resolve())
		}
	}
	sb.WriteString("</")
	sb.WriteString(e.element)
	sb.WriteString(">")

	return sb.String()
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
	escaped := html.EscapeString(fmt.Sprintf("%v", value))
	return ERAW(escaped)
}

// CR creates an Elem with unescaped HTML content or plain text.
func CR(value interface{}) Elem {
	raw := fmt.Sprintf("%v", value)
	return ERAW(raw)
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
