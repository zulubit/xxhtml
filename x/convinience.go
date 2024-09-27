package x

// Convenience functions

// Div creates a new <div> element with optional children.
func Div(children ...Elem) Elem {
	return E("div", children...)
}

// Span creates a new <span> element with optional children.
func Span(children ...Elem) Elem {
	return E("span", children...)
}

// P creates a new <p> (paragraph) element with optional children.
func P(children ...Elem) Elem {
	return E("p", children...)
}

// A creates a new <a> (anchor) element with optional children.
func A(children ...Elem) Elem {
	return E("a", children...)
}

// Img creates a new <img> (image) element. This is a self-closing tag.
func Img(children ...Elem) Elem {
	return E("img", children...).SelfClose()
}

// H1 creates a new <h1> (heading) element with optional children.
func H1(children ...Elem) Elem {
	return E("h1", children...)
}

// H2 creates a new <h2> (heading) element with optional children.
func H2(children ...Elem) Elem {
	return E("h2", children...)
}

// H3 creates a new <h3> (heading) element with optional children.
func H3(children ...Elem) Elem {
	return E("h3", children...)
}

// Ul creates a new <ul> (unordered list) element with optional children.
func Ul(children ...Elem) Elem {
	return E("ul", children...)
}

// Ol creates a new <ol> (ordered list) element with optional children.
func Ol(children ...Elem) Elem {
	return E("ol", children...)
}

// Li creates a new <li> (list item) element with optional children.
func Li(children ...Elem) Elem {
	return E("li", children...)
}

// Table creates a new <table> element with optional children.
func Table(children ...Elem) Elem {
	return E("table", children...)
}

// Tr creates a new <tr> (table row) element with optional children.
func Tr(children ...Elem) Elem {
	return E("tr", children...)
}

// Td creates a new <td> (table cell) element with optional children.
func Td(children ...Elem) Elem {
	return E("td", children...)
}

// Th creates a new <th> (table header cell) element with optional children.
func Th(children ...Elem) Elem {
	return E("th", children...)
}

// Form creates a new <form> element with optional children.
func Form(children ...Elem) Elem {
	return E("form", children...)
}

// Input creates a new <input> element. This is a self-closing tag.
func Input(children ...Elem) Elem {
	return E("input", children...).SelfClose()
}

// Button creates a new <button> element with optional children.
func Button(children ...Elem) Elem {
	return E("button", children...)
}

// Label creates a new <label> element with optional children.
func Label(children ...Elem) Elem {
	return E("label", children...)
}

// Article creates a new <article> element with optional children.
func Article(children ...Elem) Elem {
	return E("article", children...)
}

// Aside creates a new <aside> element with optional children.
func Aside(children ...Elem) Elem {
	return E("aside", children...)
}

// Header creates a new <header> element with optional children.
func Header(children ...Elem) Elem {
	return E("header", children...)
}

// Footer creates a new <footer> element with optional children.
func Footer(children ...Elem) Elem {
	return E("footer", children...)
}

// Main creates a new <main> element with optional children.
func Main(children ...Elem) Elem {
	return E("main", children...)
}

// Section creates a new <section> element with optional children.
func Section(children ...Elem) Elem {
	return E("section", children...)
}

// Nav creates a new <nav> element with optional children.
func Nav(children ...Elem) Elem {
	return E("nav", children...)
}

// Figure creates a new <figure> element with optional children.
func Figure(children ...Elem) Elem {
	return E("figure", children...)
}

// Figcaption creates a new <figcaption> element with optional children.
func Figcaption(children ...Elem) Elem {
	return E("figcaption", children...)
}

// Datalist creates a new <datalist> element with optional children.
func Datalist(children ...Elem) Elem {
	return E("datalist", children...)
}

// Option creates a new <option> element.
func Option(children ...Elem) Elem {
	return E("option", children...).SelfClose()
}

// Details creates a new <details> element with optional children.
func Details(children ...Elem) Elem {
	return E("details", children...)
}

// Summary creates a new <summary> element with optional children.
func Summary(children ...Elem) Elem {
	return E("summary", children...)
}

// Dialog creates a new <dialog> element with optional children.
func Dialog(children ...Elem) Elem {
	return E("dialog", children...)
}

// Embed creates a new <embed> element. This is a self-closing tag.
func Embed(children ...Elem) Elem {
	return E("embed", children...).SelfClose()
}

// Map creates a new <map> element with optional children.
func Map(children ...Elem) Elem {
	return E("map", children...)
}

// Area creates a new <area> element. This is a self-closing tag.
func Area(children ...Elem) Elem {
	return E("area", children...).SelfClose()
}

// Source creates a new <source> element. This is a self-closing tag.
func Source(children ...Elem) Elem {
	return E("source", children...).SelfClose()
}

// Track creates a new <track> element. This is a self-closing tag.
func Track(children ...Elem) Elem {
	return E("track", children...).SelfClose()
}

// Param creates a new <param> element. This is a self-closing tag.
func Param(children ...Elem) Elem {
	return E("param", children...).SelfClose()
}

// Script creates a new <script> element with optional children.
func Script(children ...Elem) Elem {
	return E("script", children...)
}

// Style creates a new <style> element with optional children.
func Style(children ...Elem) Elem {
	return E("style", children...)
}

// Meta creates a new <meta> element. This is a self-closing tag.
func Meta(children ...Elem) Elem {
	return E("meta", children...).SelfClose()
}

// Link creates a new <link> element. This is a self-closing tag.
func Link(children ...Elem) Elem {
	return E("link", children...).SelfClose()
}

// Title creates a new <title> element with optional children.
func Title(children ...Elem) Elem {
	return E("title", children...)
}

// Base creates a new <base> element. This is a self-closing tag.
func Base(children ...Elem) Elem {
	return E("base", children...).SelfClose()
}

// DOCTYPE generates the raw <!DOCTYPE html> declaration.
func DOCTYPE() Elem {
	return ERAW("<!DOCTYPE html>")
}

// Html creates a new <html> element with optional children.
func Html(children ...Elem) Elem {
	return E("html", children...)
}

// Head creates a new <head> element with optional children.
func Head(children ...Elem) Elem {
	return E("head", children...)
}

// Body creates a new <body> element with optional children.
func Body(children ...Elem) Elem {
	return E("body", children...)
}
