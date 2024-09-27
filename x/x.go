package x

import (
	"html"
	"strings"
)

// Elem represents an HTML element with attributes, text, and children.
type Elem struct {
	element     string
	children    []Elem
	tag         *X
	value       string
	selfClosing bool
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
//	elem := xx.E("div", x.X{class: "container", value: "Hello, World!"})
//	response := elem.Render()
//	fmt.Println(string(response)) // Outputs: <div class="container">Hello, World!</div>
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

	if tr.tag != nil {
		if tr.tag.Class != "" {
			sb.WriteString(" class=\"")
			sb.WriteString(tr.tag.Class)
			sb.WriteString("\"")
		}
		if tr.tag.Id != "" {
			sb.WriteString(" id=\"")
			sb.WriteString(tr.tag.Id)
			sb.WriteString("\"")
		}
		if tr.tag.Att != "" {
			sb.WriteString(" ")
			sb.WriteString(tr.tag.Att)
		}
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

// E initializes a new Elem with the specified tag name, attributes from X, and optional children.
func E(name string, tag X, children ...Elem) Elem {
	// Initialize the element with the specified tag name
	return Elem{element: name, tag: &tag, children: children}
}

// ERAW creates an Elem with raw HTML content or plain text.
func ERAW(value string) Elem {
	el := E("", X{})
	el.value = value
	return el
}

// C creates an Elem with raw HTML content or plain text.
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

// Convenience functions

// Div creates a new <div> element with optional children.
func Div(x X, children ...Elem) Elem {
	return E("div", x, children...)
}

// Span creates a new <span> element with optional children.
func Span(x X, children ...Elem) Elem {
	return E("span", x, children...)
}

// P creates a new <p> (paragraph) element with optional children.
func P(x X, children ...Elem) Elem {
	return E("p", x, children...)
}

// A creates a new <a> (anchor) element with optional children.
func A(x X, children ...Elem) Elem {
	return E("a", x, children...)
}

// Img creates a new <img> (image) element.
func Img(x X) Elem {
	elem := E("img", x)
	elem.selfClosing = true
	return elem
}

// H1 creates a new <h1> (heading) element with optional children.
func H1(x X, children ...Elem) Elem {
	return E("h1", x, children...)
}

// H2 creates a new <h2> (heading) element with optional children.
func H2(x X, children ...Elem) Elem {
	return E("h2", x, children...)
}

// H3 creates a new <h3> (heading) element with optional children.
func H3(x X, children ...Elem) Elem {
	return E("h3", x, children...)
}

// Ul creates a new <ul> (unordered list) element with optional children.
func Ul(x X, children ...Elem) Elem {
	return E("ul", x, children...)
}

// Ol creates a new <ol> (ordered list) element with optional children.
func Ol(x X, children ...Elem) Elem {
	return E("ol", x, children...)
}

// Li creates a new <li> (list item) element.
func Li(x X, children ...Elem) Elem {
	return E("li", x, children...)
}

// Table creates a new <table> element with optional children.
func Table(x X, children ...Elem) Elem {
	return E("table", x, children...)
}

// Tr creates a new <tr> (table row) element with optional children.
func Tr(x X, children ...Elem) Elem {
	return E("tr", x, children...)
}

// Td creates a new <td> (table cell) element.
func Td(x X) Elem {
	return E("td", x)
}

// Th creates a new <th> (table header cell) element.
func Th(x X) Elem {
	return E("th", x)
}

// Form creates a new <form> element with optional children.
func Form(x X, children ...Elem) Elem {
	return E("form", x, children...)
}

// Input creates a new <input> element.
func Input(x X) Elem {
	elem := E("input", x)
	elem.selfClosing = true
	return elem
}

// Button creates a new <button> element with optional children.
func Button(x X, children ...Elem) Elem {
	return E("button", x, children...)
}

// Label creates a new <label> element with optional children.
func Label(x X, children ...Elem) Elem {
	return E("label", x, children...)
}

// Article creates a new <article> element with optional children.
func Article(x X, children ...Elem) Elem {
	return E("article", x, children...)
}

// Aside creates a new <aside> element with optional children.
func Aside(x X, children ...Elem) Elem {
	return E("aside", x, children...)
}

// Header creates a new <header> element with optional children.
func Header(x X, children ...Elem) Elem {
	return E("header", x, children...)
}

// Footer creates a new <footer> element with optional children.
func Footer(x X, children ...Elem) Elem {
	return E("footer", x, children...)
}

// Main creates a new <main> element with optional children.
func Main(x X, children ...Elem) Elem {
	return E("main", x, children...)
}

// Section creates a new <section> element with optional children.
func Section(x X, children ...Elem) Elem {
	return E("section", x, children...)
}

// Nav creates a new <nav> element with optional children.
func Nav(x X, children ...Elem) Elem {
	return E("nav", x, children...)
}

// Figure creates a new <figure> element with optional children.
func Figure(x X, children ...Elem) Elem {
	return E("figure", x, children...)
}

// Figcaption creates a new <figcaption> element with optional children.
func Figcaption(x X, children ...Elem) Elem {
	return E("figcaption", x, children...)
}

// Datalist creates a new <datalist> element with optional children.
func Datalist(x X, children ...Elem) Elem {
	return E("datalist", x, children...)
}

// Option creates a new <option> element with optional children.
func Option(x X) Elem {
	return E("option", x)
}

// Details creates a new <details> element with optional children.
func Details(x X, children ...Elem) Elem {
	return E("details", x, children...)
}

// Summary creates a new <summary> element with optional children.
func Summary(x X, children ...Elem) Elem {
	return E("summary", x, children...)
}

// Dialog creates a new <dialog> element with optional children.
func Dialog(x X, children ...Elem) Elem {
	return E("dialog", x, children...)
}

// Embed creates a new <embed> element with optional attributes.
func Embed(x X) Elem {
	elem := E("embed", x)
	elem.selfClosing = true
	return elem
}

// Map creates a new <map> element with optional children.
func Map(x X, children ...Elem) Elem {
	return E("map", x, children...)
}

// Area creates a new <area> element with optional attributes.
func Area(x X) Elem {
	elem := E("area", x)
	elem.selfClosing = true
	return elem
}

// Source creates a new <source> element with optional attributes.
func Source(x X) Elem {
	elem := E("source", x)
	elem.selfClosing = true
	return elem
}

// Track creates a new <track> element with optional attributes.
func Track(x X) Elem {
	elem := E("track", x)
	elem.selfClosing = true
	return elem
}

// Param creates a new <param> element with optional attributes.
func Param(x X) Elem {
	elem := E("param", x)
	elem.selfClosing = true
	return elem
}

// Script creates a new <script> element with optional attributes.
func Script(x X) Elem {
	return E("script", x)
}

// Style creates a new <style> element with optional children.
func Style(x X, children ...Elem) Elem {
	return E("style", x, children...)
}

// Meta creates a new <meta> element with optional attributes.
func Meta(x X) Elem {
	elem := E("meta", x)
	elem.selfClosing = true
	return elem
}

// Link creates a new <link> element with optional attributes.
func Link(x X) Elem {
	elem := E("link", x)
	elem.selfClosing = true
	return elem
}

// Title creates a new <title> element with optional children.
func Title(x X, children ...Elem) Elem {
	return E("title", x, children...)
}

// Base creates a new <base> element with optional attributes.
func Base(x X) Elem {
	elem := E("base", x)
	elem.selfClosing = true
	return elem
}

// DOCTYPE generates the raw <!DOCTYPE html> declaration.
func DOCTYPE() Elem {
	return ERAW("<!DOCTYPE html>")
}

// Html creates a new <html> element with optional children.
func Html(x X, children ...Elem) Elem {
	return E("html", x, children...)
}

// Head creates a new <body> element with optional children.
func Head(x X, children ...Elem) Elem {
	return E("head", x, children...)
}

// Body creates a new <body> element with optional children.
func Body(x X, children ...Elem) Elem {
	return E("body", x, children...)
}
