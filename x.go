package x

import (
	"fmt"
	"html"
	"io"
)

// NodeType represents the type of an HTML node.
type NodeType int

const (
	EmptyNode      NodeType = iota // Represents an empty node; default value (0)
	TagNode                        // Represents an HTML tag
	AttributeNode                  // Represents an HTML attribute
	ContentNode                    // Represents text content (escaped)
	RawContentNode                 // Represents unescaped (raw) HTML content
)

// Elem represents an HTML element with attributes, text, and children.
type Elem struct {
	Type       NodeType // Type of the node (TagNode, AttributeNode, etc.)
	Tag        string   // Tag name (for TagNode)
	AttrKey    string   // Attribute key (for AttributeNode)
	AttrVal    string   // Attribute value (for AttributeNode)
	Content    string   // Text content (for ContentNode or RawContentNode)
	Children   []Elem   // Child nodes
	SelfCloses bool     // Indicates if the element is self-closing
}

// Render writes the HTML representation of the element and its children to an io.Writer.
func (e Elem) Render(w io.Writer) error {
	return e.render(w)
}

func (e Elem) render(w io.Writer) error {
	switch e.Type {
	case EmptyNode:
		// Do nothing for empty nodes.
		return nil
	case TagNode:
		// Write opening tag.
		if _, err := w.Write([]byte("<" + e.Tag)); err != nil {
			return err
		}

		// Render attributes: assume that attribute nodes come first.
		firstNonAttrIndex := len(e.Children)
		for i, child := range e.Children {
			if child.Type == AttributeNode {
				if err := child.render(w); err != nil {
					return err
				}
			} else {
				firstNonAttrIndex = i
				break
			}
		}

		// Handle self-closing tags: if the element is marked as self-closing
		// and has no non-attribute children, output as self-closing.
		if e.SelfCloses && firstNonAttrIndex == len(e.Children) {
			if _, err := w.Write([]byte(" />")); err != nil {
				return err
			}
			return nil
		}

		// Write closing angle bracket for opening tag.
		if _, err := w.Write([]byte(">")); err != nil {
			return err
		}

		// Render non-attribute children.
		for i := firstNonAttrIndex; i < len(e.Children); i++ {
			child := e.Children[i]
			// Render any child that is not an attribute.
			if child.Type != AttributeNode {
				if err := child.render(w); err != nil {
					return err
				}
			}
		}

		// Write closing tag.
		if _, err := w.Write([]byte("</" + e.Tag + ">")); err != nil {
			return err
		}
	case AttributeNode:
		attrStr := " " + e.AttrKey
		if e.AttrVal != "" {
			attrStr += fmt.Sprintf(`="%s"`, html.EscapeString(e.AttrVal))
		}
		if _, err := w.Write([]byte(attrStr)); err != nil {
			return err
		}
	case ContentNode:
		// Write escaped content.
		if _, err := w.Write([]byte(html.EscapeString(e.Content))); err != nil {
			return err
		}
	case RawContentNode:
		// Write raw (unescaped) content.
		if _, err := w.Write([]byte(e.Content)); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown node type: %d", e.Type)
	}

	return nil
}

// E initializes a new Elem with the specified tag name and optional children.
func E(tag string, children ...Elem) Elem {
	return Elem{
		Type:     TagNode,
		Tag:      tag,
		Children: children,
	}
}

// Att creates an Elem representing an HTML attribute with a key-value pair.
func Att(key, value string) Elem {
	return Elem{
		Type:    AttributeNode,
		AttrKey: key,
		AttrVal: value,
	}
}

// C creates an Elem with escaped HTML content or plain text.
func C(value interface{}) Elem {
	content := fmt.Sprintf("%v", value)
	return Elem{
		Type:    ContentNode,
		Content: html.EscapeString(content),
	}
}

// CR creates an Elem with unescaped HTML content or plain text.
func CR(value interface{}) Elem {
	content := fmt.Sprintf("%v", value)
	return Elem{
		Type:    RawContentNode,
		Content: content,
	}
}

// SelfClose marks an element as self-closing.
func (e Elem) SelfClose() Elem {
	e.SelfCloses = true
	return e
}

// IF returns trueCase if the condition is true, otherwise returns an empty Elem.
func IF(condition bool, trueCase Elem) Elem {
	if condition {
		return trueCase
	}
	return Elem{Type: EmptyNode}
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
