# HTML Builder Library

A simple Go library for creating HTML elements programmatically.

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
    "github.com/6oof/xxhtml/xx"
)

func main() {
    router := http.NewServeMux()

    count := []int{1, 2, 3}
    tm := xx.E("div").CLS("test test2").VAL("hello world").
        C(xx.E("h1").CLS("asdf").VAL("hello world").
            C(xx.E("span").VAL(xx.STER(false, "yes", "no")))).
        C(xx.FOR(func(count []int) []xx.Elem {
            var elems []xx.Elem
            for _, c := range count {
                elems = append(elems, xx.E("span").VAL(strconv.Itoa(c)))
            }
            return elems
        }(count))).
        C(xx.IF(false, xx.E("h1").VAL("hello world"))).
        C(xx.ERAW("<svg>asdf</svg>"))

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write(tm.Render())
    })

    http.ListenAndServe(":8083", router)
}
```

## Documentation

### Element Creation

- `func E(name string) Elem`  
  Initializes a new `Elem` with the specified tag name.

- `func ERAW(value string) Elem`  
  Creates a raw HTML element, often used for inserting raw HTML content or plain text.

- `func (el Elem) CLS(class string) Elem`  
  Adds a class to the element.

- `func (el Elem) VAL(value string) Elem`  
  Sets or appends the text content of the element.

- `func (el Elem) ATT(value string) Elem`  
  Adds an attribute to the element.

- `func (el Elem) C(child Elem) Elem`  
  Appends a child element.

### Conditional and Loop Functions

- `func IF(condition bool, trueCase Elem) Elem`  
  Returns `trueCase` if the condition is true; otherwise returns an empty `Elem`.

- `func FOR(iterClosure []Elem) Elem`  
  Takes a slice of `Elem` and returns a parent `Elem` containing all elements in the slice as its children.

- `func TER(condition bool, trueCase Elem, falseCase Elem) Elem`  
  Returns `trueCase` if the condition is true; otherwise returns `falseCase`.

- `func STER(condition bool, trueCase string, falseCase string) string`  
  Returns `trueCase` if the boolean condition is true; otherwise returns `falseCase`.

- `func SIF(condition bool, trueCase string) string`  
  Returns `trueCase` if the boolean condition is true; otherwise returns an empty string.

