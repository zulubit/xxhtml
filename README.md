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
    elem := x.Div(x.X{Class: "container"}, 
        x.H1(x.X{}, x.C("Hello, World!")),
        x.P(x.X{}, x.C("This is a paragraph.")),
    )
    response := elem.Render()
    fmt.Println(string(response)) // Outputs: <div class="container"><h1>Hello, World!</h1><p>This is a paragraph.</p></div>
}
```


### Key Types

- **`Elem`**: Represents an HTML element with attributes, text, and children.
- **`X`**: Represents attributes like class, id, and other HTML attributes.

### Key Functions

- **`E(name string, tag X, children ...Elem) Elem`**: Creates a new `Elem` with the specified tag name, attributes, and optional children.
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

- **`Div(x X, children ...Elem) Elem`**: Creates a `<div>` element.
- **`Span(x X, children ...Elem) Elem`**: Creates a `<span>` element.
- **`P(x X, children ...Elem) Elem`**: Creates a `<p>` element.
- **`A(x X, children ...Elem) Elem`**: Creates an `<a>` element.
- **`H1(x X, children ...Elem) Elem`**: Creates an `<h1>` element.
- **`H2(x X, children ...Elem) Elem`**: Creates an `<h2>` element.
- **`H3(x X, children ...Elem) Elem`**: Creates an `<h3>` element.
- **`Ul(x X, children ...Elem) Elem`**: Creates a `<ul>` element.
- **`Ol(x X, children ...Elem) Elem`**: Creates an `<ol>` element.
- **`Li(x X) Elem`**: Creates a `<li>` element.
- **`Table(x X, children ...Elem) Elem`**: Creates a `<table>` element.
- **`Tr(x X, children ...Elem) Elem`**: Creates a `<tr>` element.
- **`Td(x X) Elem`**: Creates a `<td>` element.
- **`Th(x X) Elem`**: Creates a `<th>` element.
- **`Form(x X, children ...Elem) Elem`**: Creates a `<form>` element.
- **`Button(x X, children ...Elem) Elem`**: Creates a `<button>` element.
- **`Label(x X, children ...Elem) Elem`**: Creates a `<label>` element.
- **`Article(x X, children ...Elem) Elem`**: Creates an `<article>` element.
- **`Aside(x X, children ...Elem) Elem`**: Creates an `<aside>` element.
- **`Header(x X, children ...Elem) Elem`**: Creates a `<header>` element.
- **`Footer(x X, children ...Elem) Elem`**: Creates a `<footer>` element.
- **`Main(x X, children ...Elem) Elem`**: Creates a `<main>` element.
- **`Section(x X, children ...Elem) Elem`**: Creates a `<section>` element.
- **`Nav(x X, children ...Elem) Elem`**: Creates a `<nav>` element.
- **`Figure(x X, children ...Elem) Elem`**: Creates a `<figure>` element.
- **`Figcaption(x X, children ...Elem) Elem`**: Creates a `<figcaption>` element.
- **`Datalist(x X, children ...Elem) Elem`**: Creates a `<datalist>` element.
- **`Option(x X) Elem`**: Creates an `<option>` element.
- **`Details(x X, children ...Elem) Elem`**: Creates a `<details>` element.
- **`Summary(x X, children ...Elem) Elem`**: Creates a `<summary>` element.
- **`Dialog(x X, children ...Elem) Elem`**: Creates a `<dialog>` element.
- **`Map(x X, children ...Elem) Elem`**: Creates a `<map>` element.

#### Self-Closing Elements

- **`Img(x X) Elem`**: Creates a self-closing `<img>` element.
- **`Input(x X) Elem`**: Creates a self-closing `<input>` element.
- **`Embed(x X) Elem`**: Creates a self-closing `<embed>` element.
- **`Area(x X) Elem`**: Creates a self-closing `<area>` element.
- **`Source(x X) Elem`**: Creates a self-closing `<source>` element.
- **`Track(x X) Elem`**: Creates a self-closing `<track>` element.
- **`Param(x X) Elem`**: Creates a self-closing `<param>` element.
- **`Meta(x X) Elem`**: Creates a self-closing `<meta>` element.
- **`Link(x X) Elem`**: Creates a self-closing `<link>` element.
- **`Base(x X) Elem`**: Creates a self-closing `<base>` element.
