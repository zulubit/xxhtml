package x

// Convenience functions

// Div creates a new <div> element with optional children.
func Div(attributes string, children ...Elem) Elem {
	return E("div", attributes, children...)
}

// Span creates a new <span> element with optional children.
func Span(attributes string, children ...Elem) Elem {
	return E("span", attributes, children...)
}

// P creates a new <p> (paragraph) element with optional children.
func P(attributes string, children ...Elem) Elem {
	return E("p", attributes, children...)
}

// A creates a new <a> (anchor) element with optional children.
func A(attributes string, children ...Elem) Elem {
	return E("a", attributes, children...)
}

// Img creates a new <img> (image) element. This is a self-closing tag.
func Img(attributes string) Elem {
	elem := E("img", attributes)
	elem.selfClosing = true
	return elem
}

// H1 creates a new <h1> (heading) element with optional children.
func H1(attributes string, children ...Elem) Elem {
	return E("h1", attributes, children...)
}

// H2 creates a new <h2> (heading) element with optional children.
func H2(attributes string, children ...Elem) Elem {
	return E("h2", attributes, children...)
}

// H3 creates a new <h3> (heading) element with optional children.
func H3(attributes string, children ...Elem) Elem {
	return E("h3", attributes, children...)
}

// Ul creates a new <ul> (unordered list) element with optional children.
func Ul(attributes string, children ...Elem) Elem {
	return E("ul", attributes, children...)
}

// Ol creates a new <ol> (ordered list) element with optional children.
func Ol(attributes string, children ...Elem) Elem {
	return E("ol", attributes, children...)
}

// Li creates a new <li> (list item) element with optional children.
func Li(attributes string, children ...Elem) Elem {
	return E("li", attributes, children...)
}

// Table creates a new <table> element with optional children.
func Table(attributes string, children ...Elem) Elem {
	return E("table", attributes, children...)
}

// Tr creates a new <tr> (table row) element with optional children.
func Tr(attributes string, children ...Elem) Elem {
	return E("tr", attributes, children...)
}

// Td creates a new <td> (table cell) element with optional children.
func Td(attributes string, children ...Elem) Elem {
	return E("td", attributes, children...)
}

// Th creates a new <th> (table header cell) element with optional children.
func Th(attributes string, children ...Elem) Elem {
	return E("th", attributes, children...)
}

// Form creates a new <form> element with optional children.
func Form(attributes string, children ...Elem) Elem {
	return E("form", attributes, children...)
}

// Input creates a new <input> element. This is a self-closing tag.
func Input(attributes string) Elem {
	elem := E("input", attributes)
	elem.selfClosing = true
	return elem
}

// Button creates a new <button> element with optional children.
func Button(attributes string, children ...Elem) Elem {
	return E("button", attributes, children...)
}

// Label creates a new <label> element with optional children.
func Label(attributes string, children ...Elem) Elem {
	return E("label", attributes, children...)
}

// Article creates a new <article> element with optional children.
func Article(attributes string, children ...Elem) Elem {
	return E("article", attributes, children...)
}

// Aside creates a new <aside> element with optional children.
func Aside(attributes string, children ...Elem) Elem {
	return E("aside", attributes, children...)
}

// Header creates a new <header> element with optional children.
func Header(attributes string, children ...Elem) Elem {
	return E("header", attributes, children...)
}

// Footer creates a new <footer> element with optional children.
func Footer(attributes string, children ...Elem) Elem {
	return E("footer", attributes, children...)
}

// Main creates a new <main> element with optional children.
func Main(attributes string, children ...Elem) Elem {
	return E("main", attributes, children...)
}

// Section creates a new <section> element with optional children.
func Section(attributes string, children ...Elem) Elem {
	return E("section", attributes, children...)
}

// Nav creates a new <nav> element with optional children.
func Nav(attributes string, children ...Elem) Elem {
	return E("nav", attributes, children...)
}

// Figure creates a new <figure> element with optional children.
func Figure(attributes string, children ...Elem) Elem {
	return E("figure", attributes, children...)
}

// Figcaption creates a new <figcaption> element with optional children.
func Figcaption(attributes string, children ...Elem) Elem {
	return E("figcaption", attributes, children...)
}

// Datalist creates a new <datalist> element with optional children.
func Datalist(attributes string, children ...Elem) Elem {
	return E("datalist", attributes, children...)
}

// Option creates a new <option> element. This is a self-closing tag.
func Option(attributes string) Elem {
	elem := E("option", attributes)
	elem.selfClosing = true
	return elem
}

// Details creates a new <details> element with optional children.
func Details(attributes string, children ...Elem) Elem {
	return E("details", attributes, children...)
}

// Summary creates a new <summary> element with optional children.
func Summary(attributes string, children ...Elem) Elem {
	return E("summary", attributes, children...)
}

// Dialog creates a new <dialog> element with optional children.
func Dialog(attributes string, children ...Elem) Elem {
	return E("dialog", attributes, children...)
}

// Embed creates a new <embed> element. This is a self-closing tag.
func Embed(attributes string) Elem {
	elem := E("embed", attributes)
	elem.selfClosing = true
	return elem
}

// Map creates a new <map> element with optional children.
func Map(attributes string, children ...Elem) Elem {
	return E("map", attributes, children...)
}

// Area creates a new <area> element. This is a self-closing tag.
func Area(attributes string) Elem {
	elem := E("area", attributes)
	elem.selfClosing = true
	return elem
}

// Source creates a new <source> element. This is a self-closing tag.
func Source(attributes string) Elem {
	elem := E("source", attributes)
	elem.selfClosing = true
	return elem
}

// Track creates a new <track> element. This is a self-closing tag.
func Track(attributes string) Elem {
	elem := E("track", attributes)
	elem.selfClosing = true
	return elem
}

// Param creates a new <param> element. This is a self-closing tag.
func Param(attributes string) Elem {
	elem := E("param", attributes)
	elem.selfClosing = true
	return elem
}

// Script creates a new <script> element with optional children.
func Script(attributes string, children ...Elem) Elem {
	return E("script", attributes, children...)
}

// Style creates a new <style> element with optional children.
func Style(attributes string, children ...Elem) Elem {
	return E("style", attributes, children...)
}

// Meta creates a new <meta> element. This is a self-closing tag.
func Meta(attributes string) Elem {
	elem := E("meta", attributes)
	elem.selfClosing = true
	return elem
}

// Link creates a new <link> element. This is a self-closing tag.
func Link(attributes string) Elem {
	elem := E("link", attributes)
	elem.selfClosing = true
	return elem
}

// Title creates a new <title> element with optional children.
func Title(attributes string, children ...Elem) Elem {
	return E("title", attributes, children...)
}

// Base creates a new <base> element. This is a self-closing tag.
func Base(attributes string) Elem {
	elem := E("base", attributes)
	elem.selfClosing = true
	return elem
}

// DOCTYPE generates the raw <!DOCTYPE html> declaration.
func DOCTYPE() Elem {
	return ERAW("<!DOCTYPE html>")
}

// Html creates a new <html> element with optional children.
func Html(attributes string, children ...Elem) Elem {
	return E("html", attributes, children...)
}

// Head creates a new <head> element with optional children.
func Head(attributes string, children ...Elem) Elem {
	return E("head", attributes, children...)
}

// Body creates a new <body> element with optional children.
func Body(attributes string, children ...Elem) Elem {
	return E("body", attributes, children...)
}
