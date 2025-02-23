# XXHTML
## `x` Package Documentation

**Module**:  
```
github.com/zulubit/xxhtml/x
```

The `x` package provides a flexible and programmatic way to build HTML elements in Go. It supports creating various HTML tags, setting attributes, nesting elements, and rendering HTML as byte slices or streams. The package ensures safe HTML output by escaping content unless explicitly marked as raw.

---

### Example Usage

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/zulubit/xxhtml/x"
)

func main() {
	elem := x.E("div", x.Att("class", "container"),
		x.E("h1", x.C("Hello, World!")),
		x.E("p", x.C("This is a paragraph.")),
		x.E("img", x.Att("class", "logo"), x.Att("src", "img.png")).SelfClose(),
	)

	var buf bytes.Buffer
	err := elem.Render(&buf)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(buf.String())
}
```

**Output:**
```html
<div class="container">
    <h1>Hello, World!</h1>
    <p>This is a paragraph.</p>
    <img class="logo" src="img.png" />
</div>
```

---

### Key Types

- **`Elem`**: Represents an HTML element with attributes, text, and children. It is the core building block for constructing HTML.

---

### Key Functions

- **`E(tag string, children ...Elem) Elem`**: Creates a new `Elem` with the specified tag name and optional children.
- **`Att(key, value string) Elem`**: Creates an `Elem` representing an HTML attribute with a key-value pair.
- **`C(content interface{}) Elem`**: Creates an `Elem` with escaped text content for safe HTML output.
- **`CR(content interface{}) Elem`**: Creates an `Elem` with unescaped (raw) HTML content.
- **`E(...).SelfClose()`**: Marks an element as self-closing (e.g., ``, ``, etc.).

---

### Methods

- **Render(w io.Writer) error**: Writes the HTML representation of the element and its children to an `io.Writer`. This allows flexibility in rendering directly to buffers, files, or HTTP responses.

---

### Control Structures

The package includes several utility functions for conditional rendering and dynamic content generation:

- **`IF(condition bool, trueCase Elem) Elem`**: Returns `trueCase` if the condition is true; otherwise, returns an empty `Elem`.
  ```go
  x.IF(userLoggedIn, x.E("p", x.C("Welcome back!")))
  ```

- **`TER(condition bool, trueCase Elem, falseCase Elem) Elem`**: Returns `trueCase` if the condition is true; otherwise, returns `falseCase`.
  ```go
  x.TER(isAdmin, x.E("button", x.C("Admin Panel")), x.E("button", x.C("User Dashboard")))
  ```

- **`STER(condition bool, trueCase string, falseCase string) string`**: Returns `trueCase` if the condition is true; otherwise, returns `falseCase`.
  ```go
  role := x.STER(isAdmin, "admin", "user")
  ```

- **`SIF(condition bool, trueCase string) string`**: Returns `trueCase` if the condition is true; otherwise, returns an empty string.
  ```go
  class := x.SIF(isActive, "active")
  ```

---

### Node Types

The following node types are supported in the package:

1. **TagNode**: Represents an HTML tag (e.g., ``, ``).
2. **AttributeNode**: Represents an HTML attribute (e.g., `class="container"`).
3. **ContentNode**: Represents escaped text content for safe HTML output.
4. **RawContentNode**: Represents unescaped (raw) HTML content.
5. **EmptyNode**: Represents an empty node that renders no output.

---

### Convenience Functions

#### Non-Self-Closing Elements

Convenience functions are provided for commonly used HTML elements.

Example:

```go
x.Div(x.Class("example"), x.C("example"))
```

is eqluivalent to:

```go
x.E("div", x.Att("class", "example"), x.C("example"))
```

---

### Notes on Rendering

1. **Self-Closing Tags:**  
   Elements that are self-closing (e.g., ``, ``) must have the `.SelfClose()` method called to mark them as such.

2. **Empty Nodes:**  
   Nodes of type `EmptyNode` render no output. These are typically returned by utility functions like `IF(false)`.

3. **Escaped vs. Raw Content:**  
   - Use `C(content)` for escaped content to ensure safe HTML output.
   - Use `CR(content)` for raw (unescaped) HTML content when you trust the input.

4. **Attributes:**  
   Attributes should always be added using the `Att(key, value)` function.

5. **Child Order:**  
   Attribute nodes must precede non-attribute child nodes in the children slice for proper rendering.

---

### Example Scenarios

#### Conditional Rendering with IF and TER

```go
elem := x.E("div",
	x.IF(userLoggedIn,
		x.E("p", x.C("Welcome back!")),
	),
	x.TER(userLoggedIn,
		x.E("button", x.C("Logout")),
		x.E("button", x.C("Login")),
	),
)
```

#### Rendering a Table Dynamically

```go
rows := [][]string{
	{"Name", "Age", "City"},
	{"Alice", "30", "New York"},
	{"Bob", "25", "San Francisco"},
}

table := x.E("table", func() []x.Elem {
	var rowElems []x.Elem
	for _, row := range rows {
		var cellElems []x.Elem
		for _, cell := range row {
			cellElems = append(cellElems, x.E("td", x.C(cell)))
		}
		rowElems = append(rowElems, x.E("tr", cellElems...))
	}
	return rowElems
}()...)
```
