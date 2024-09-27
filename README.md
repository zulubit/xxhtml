# HTML Builder Library

A simple Go library to create HTML elements programmatically.

## Installation

```bash
go get github.com/6oof/xxhtml
```

## Example

```go
package main

import (
    "net/http"
    "strconv"
    "github.com/6oof/xxhtml"
)

func main() {
    router := http.NewServeMux()

    count := []int{1, 2, 3}
    tm := xx.E("div", `class="test test2"`, "hello world",
        xx.E("h1", `class="asdf"`, "hello world",
            xx.E("span", "", xx.STER(false, "yes", "no")),
            xx.FOR(func(count []int) []xx.Elem {
                el := []xx.Elem{}
                for _, c := range count {
                    el = append(el, xx.E("span", "", strconv.Itoa(c)))
                }
                return el
            }(count)),
        ),
        xx.IF(false, xx.E("h1", "", "hello world")),
        xx.ERAW(`<svg>asdf</svg>`),
    )

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write(tm.Render())
    })

    http.ListenAndServe(":8083", router)
}
```

## Documentation

### Element Creation

- `func E(name string, attributes string, value string, children ...Elem) Elem`  
  Initializes a new `Elem` with the specified tag name, attributes, text content, and children.

- `func ERAW(value string) Elem`  
  Creates a raw HTML element, often used for inserting raw HTML content or plain text.

### Conditional and Loop Functions

- `func IF(condition bool, trueCase Elem) Elem`  
  Returns `trueCase` if the condition is true, otherwise returns an empty `Elem`.

- `func FOR(iterClosure []Elem) Elem`  
  Takes a slice of `Elem` and returns a parent `Elem` containing all elements in the slice as its children.

- `func TER(condition bool, trueCase Elem, falseCase Elem) Elem`  
  Returns `trueCase` if the condition is true, otherwise returns `falseCase`.

- `func STER(condition bool, trueCase string, falseCase string) string`  
  Returns `trueCase` if the boolean condition is true, otherwise returns `falseCase`.

- `func SIF(condition bool, trueCase string) string`  
  Returns `trueCase` if the boolean condition is true, otherwise returns an empty string.

