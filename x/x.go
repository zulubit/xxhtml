package x

import (
	"bytes"
	"fmt"
	"html"
)

// Elem represents an HTML element with attributes, text, and children.
// eltype: 1 = tag, 2 = class, 3 = attribute, 4 = content
type Elem struct {
	children *[]Elem
	key      *string
	value    *string
	eltype   int
	sc       bool // Indicates if the element is self-closing
}

type El []byte

// Render generates the HTML representation of the element and its children as a byte slice.
func (e Elem) Render() []byte {
	return e.resolve()
}

// resolve constructs the HTML string for the element and recursively for its children.
func (e Elem) resolve() El {
	var buf bytes.Buffer

	switch e.eltype {
	case 1: // Tag element
		if e.value == nil || *e.value == "" {
			// If the tag name is empty, render only children if they exist.
			if e.children != nil {
				for _, c := range *e.children {
					buf.Write(c.resolve())
				}
			}
			return buf.Bytes()
		}

		buf.WriteString("<")
		buf.WriteString(*e.value)

		if e.children != nil {
			for _, c := range *e.children {
				if c.eltype == 2 || c.eltype == 3 { // Class or Attribute
					buf.WriteString(" ")
					buf.Write(c.resolve())
				}
			}
		}

		if e.sc { // Self-closing tag
			buf.WriteString(" />")
			return buf.Bytes()
		}

		buf.WriteString(">")

		if e.children != nil {
			for _, c := range *e.children {
				if c.eltype == 4 { // Content
					buf.Write(c.resolve())
				} else if c.eltype != 2 && c.eltype != 3 { // Other children elements
					buf.Write(c.resolve())
				}
			}
		}

		buf.WriteString("</" + *e.value + ">")

	case 2: // Class element
		buf.WriteString(`class="`)
		buf.WriteString(*e.value)
		buf.WriteString(`"`)

	case 3: // Attribute element
		buf.WriteString(*e.key)
		buf.WriteString(`="`)
		buf.WriteString(*e.value)
		buf.WriteString(`"`)

	case 4: // Content element
		buf.WriteString(*e.value)

	default:
		// Handle any unexpected types if necessary
	}

	return buf.Bytes()
}

// E initializes a new Elem with the specified tag name, attributes, and optional children.
func E(name string, children ...Elem) Elem {
	var nam *string
	nam = &name
	var childrenPtr *[]Elem
	if len(children) > 0 {
		childrenPtr = &children
	}
	return Elem{value: nam, children: childrenPtr, eltype: 1}
}

// Class creates an Elem representing a CSS class.
func Class(classes string) Elem {
	var class *string
	class = &classes
	return Elem{eltype: 2, value: class}
}

// Att creates an Elem representing an HTML attribute with a key-value pair.
func Att(key string, value string) Elem {
	var att *string
	var k *string
	att = &value
	k = &key
	return Elem{eltype: 3, value: att, key: k}
}

// SelfClose marks an element as self-closing.
func (e Elem) SelfClose() Elem {
	if e.eltype == 1 {
		e.sc = true
	}
	return e
}

// ERAW creates an Elem with raw HTML content or plain text.
func ERAW(value string) Elem {
	val := value
	return Elem{value: &val, eltype: 4}
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
