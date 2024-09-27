# XXHTML
## `x` Package Documentation

**Module:** [github.com/6oof/xxhtml/x](https://github.com/6oof/xxhtml/x)

The `x` package provides a flexible way to build HTML elements programmatically in Go. It enables creating various HTML tags, setting attributes, nesting elements, and rendering HTML as byte slices.

### Example Usage

```go
package main

import (
    "fmt"
    "github.com/6oof/xxhtml/x"
)

func main() {
    elem := x.Div(`class="container"`, 
        x.H1("", x.C("Hello, World!")),
        x.P("", x.C("This is a paragraph.")),
    )
    response := elem.Render()
    fmt.Println(string(response)) // Outputs: <div class="container"><h1>Hello, World!</h1><p>This is a paragraph.</p></div>
}
```


### Key Types

- **`Elem`**: Represents an HTML element with attributes, text, and children.
- **`X`**: Represents attributes like class, id, and other HTML attributes.

### Key Functions

- **`E(name string, attributes string, children ...Elem) Elem`**: Creates a new `Elem` with the specified tag name, attributes, and optional children.
- **`ERAW(value string) Elem`**: Creates an `Elem` with raw HTML content or plain text.
- **`C(content string) Elem`**: Creates an `Elem` with content that is automatically escaped for safe HTML output.

### Control Structures

- **`IF(condition bool, trueCase Elem) Elem`**: Returns `trueCase` if the condition is true; otherwise, returns an empty `Elem`.
- **`FOR(iterClosure []Elem) []Elem`**: Iterates over a slice of `Elem` and returns all elements in the slice.
- **`TER(condition bool, trueCase Elem, falseCase Elem) Elem`**: Returns `trueCase` if the condition is true; otherwise, returns `falseCase`.
- **`STER(condition bool, trueCase string, falseCase string) string`**: Returns `trueCase` if the condition is true; otherwise, returns `falseCase`.
- **`SIF(condition bool, trueCase string) string`**: Returns `trueCase` if the condition is true; otherwise, returns an empty string.

### Convenience Functions

#### Non-Self-Closing Elements

- **`Div(attributes string, children ...Elem) Elem`**: Creates a `<div>` element.
- **`Span(attributes string, children ...Elem) Elem`**: Creates a `<span>` element.
- **`P(attributes string, children ...Elem) Elem`**: Creates a `<p>` element.
- **`A(attributes string, children ...Elem) Elem`**: Creates an `<a>` element.
- **`H1(attributes string, children ...Elem) Elem`**: Creates an `<h1>` element.
- **`H2(attributes string, children ...Elem) Elem`**: Creates an `<h2>` element.
- **`H3(attributes string, children ...Elem) Elem`**: Creates an `<h3>` element.
- **`Ul(attributes string, children ...Elem) Elem`**: Creates a `<ul>` element.
- **`Ol(attributes string, children ...Elem) Elem`**: Creates an `<ol>` element.
- **`Li(attributes string) Elem`**: Creates a `<li>` element.
- **`Table(attributes string, children ...Elem) Elem`**: Creates a `<table>` element.
- **`Tr(attributes string, children ...Elem) Elem`**: Creates a `<tr>` element.
- **`Td(attributes string) Elem`**: Creates a `<td>` element.
- **`Th(attributes string) Elem`**: Creates a `<th>` element.
- **`Form(attributes string, children ...Elem) Elem`**: Creates a `<form>` element.
- **`Button(attributes string, children ...Elem) Elem`**: Creates a `<button>` element.
- **`Label(attributes string, children ...Elem) Elem`**: Creates a `<label>` element.
- **`Article(attributes string, children ...Elem) Elem`**: Creates an `<article>` element.
- **`Aside(attributes string, children ...Elem) Elem`**: Creates an `<aside>` element.
- **`Header(attributes string, children ...Elem) Elem`**: Creates a `<header>` element.
- **`Footer(attributes string, children ...Elem) Elem`**: Creates a `<footer>` element.
- **`Main(attributes string, children ...Elem) Elem`**: Creates a `<main>` element.
- **`Section(attributes string, children ...Elem) Elem`**: Creates a `<section>` element.
- **`Nav(attributes string, children ...Elem) Elem`**: Creates a `<nav>` element.
- **`Figure(attributes string, children ...Elem) Elem`**: Creates a `<figure>` element.
- **`Figcaption(attributes string, children ...Elem) Elem`**: Creates a `<figcaption>` element.
- **`Datalist(attributes string, children ...Elem) Elem`**: Creates a `<datalist>` element.
- **`Option(attributes string) Elem`**: Creates an `<option>` element.
- **`Details(attributes string, children ...Elem) Elem`**: Creates a `<details>` element.
- **`Summary(attributes string, children ...Elem) Elem`**: Creates a `<summary>` element.
- **`Dialog(attributes string, children ...Elem) Elem`**: Creates a `<dialog>` element.
- **`Map(attributes string, children ...Elem) Elem`**: Creates a `<map>` element.

#### Self-Closing Elements

- **`Img(attributes string) Elem`**: Creates a self-closing `<img>` element.
- **`Input(attributes string) Elem`**: Creates a self-closing `<input>` element.
- **`Embed(attributes string) Elem`**: Creates a self-closing `<embed>` element.
- **`Area(attributes string) Elem`**: Creates a self-closing `<area>` element.
- **`Source(attributes string) Elem`**: Creates a self-closing `<source>` element.
- **`Track(attributes string) Elem`**: Creates a self-closing `<track>` element.
- **`Param(attributes string) Elem`**: Creates a self-closing `<param>` element.
- **`Meta(attributes string) Elem`**: Creates a self-closing `<meta>` element.
- **`Link(attributes string) Elem`**: Creates a self-closing `<link>` element.
- **`Base(attributes string) Elem`**: Creates a self-closing `<base>` element.
