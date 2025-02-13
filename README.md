# XXHTML
## `x` Package Documentation

**Module** 

```
github.com/zulubit/xxhtml/x
```

The `x` package provides a flexible way to build HTML elements programmatically in Go. It enables creating various HTML tags, setting attributes, nesting elements, and rendering HTML as byte slices.

### Example Usage

```go
package main

import (
    "fmt"
    "github.com/zulubit/xxhtml/x"
)

func main() {
    elem := x.Div(x.Class("container"), 
        x.H1(x.C("Hello, World!")),
        x.P(x.C("This is a paragraph.")),
        x.Img(x.Class("logo"), x.Att("src", "img.png"))
    )
    response := elem.Render()
    fmt.Println(string(response)) // Outputs: <div class="container"><h1>Hello, World!</h1><p>This is a paragraph.</p></div>
}
```

### Key Types

- **`Elem`**: Represents an HTML element with attributes, text, and children.

### Key Functions

- **`E(name string, children ...Elem) Elem`**: Creates a new `Elem` with the specified tag name and optional children. If the `name` is empty, it will not render the tag itself but will render its children.
- **`ERAW(value string) Elem`**: Creates an `Elem` with raw HTML content or plain text.
- **`C(content interface{}) Elem`**: Creates an `Elem` with content that is automatically escaped for safe HTML output.
- **`CR(content interface{}) Elem`**: Creates an `Elem` with content that is not escaped, for unescaped HTML output.

### Methods

- **`SelfClose()`**: Marks the `Elem` as self-closing. This method is typically used for HTML elements that do not have closing tags (e.g., `<img>`, `<input>`). 

### Control Structures

- **`IF(condition bool, trueCase Elem) Elem`**: Returns `trueCase` if the condition is true; otherwise, returns an empty `Elem`.
- **`FOR(iterClosure []Elem) []Elem`**: Iterates over a slice of `Elem` and returns all elements in the slice.
- **`TER(condition bool, trueCase Elem, falseCase Elem) Elem`**: Returns `trueCase` if the condition is true; otherwise, returns `falseCase`.
- **`STER(condition bool, trueCase string, falseCase string) string`**: Returns `trueCase` if the condition is true; otherwise, returns `falseCase`.
- **`SIF(condition bool, trueCase string) string`**: Returns `trueCase` if the condition is true; otherwise, returns an empty string.

### Convenience Functions

#### Non-Self-Closing Elements

Convenience functions are provided for various non-self-closing HTML elements, including: `Div`, `Span`, `P`, `A`, `H1`, `H2`, `H3`, `Ul`, `Ol`, `Li`, `Table`, `Tr`, `Td`, `Th`, `Form`, `Button`, `Label`, `Article`, `Aside`, `Header`, `Footer`, `Main`, `Section`, `Nav`, `Figure`, `Figcaption`, `Datalist`, `Details`, `Summary`, `Dialog`, `Map`, `Html`, `Head`, and `Body`.

#### Self-Closing Elements

Convenience functions are provided for self-closing HTML elements, including: `Img`, `Input`, `Embed`, `Area`, `Source`, `Track`, `Param`, `Meta`, `Link`, and `Base`.

### Notes

- **Self-Closing Tags**: Elements that are self-closing in HTML (like `<img>`, `<input>`, etc.) should have the `SelfClose` attribute set to `true`.
- **Non-Self-Closing Tags**: Most other HTML elements will require a closing tag and can include children elements.
- **Empty Tag Names**: If an element has an empty tag name, it will not render the opening and closing tags but will render its children if any are present.

Feel free to explore and use these functions to create HTML structures programmatically in your Go applications.
