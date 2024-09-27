# HTML Builder Library

A simple Go library to create HTML elements programmatically.

## Installation

    go get github.com/your-username/your-library-name

## Usage

```
    package main

    import (
        "net/http"
        "strconv"
        "github.com/your-username/your-library-name"
    )

    func main() {
        router := http.NewServeMux()

        count := []int{1, 2, 3}
        tm := htmlbuilder.XX("div").Cls("test").Cls("test2").Val("hello world").Add(
            htmlbuilder.XX("h1").Cls("asdf").Val("hello world").Add(
                htmlbuilder.XX("span").Val(htmlbuilder.Xter(false, "yes", "no")),
                htmlbuilder.XFOR(func(count []int) []htmlbuilder.Elem {
                    el := []htmlbuilder.Elem{}
                    for _, c := range count {
                        el = append(el, htmlbuilder.XX("span").Val(strconv.Itoa(c)))
                    }
                    return el
                }(count)),
            ),
            htmlbuilder.XIF(false, htmlbuilder.XX("h1").Val("hello world"), htmlbuilder.XX("")),
            htmlbuilder.XX("").Raw().Val(`<svg>asdf</svg>`),
        )

        router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            w.Write(tm.CreateResponse())
        })

        http.ListenAndServe(":8083", router)
    }
```

## Documentation

- func XX(name string) Elem
- func XIF(condition bool, trueCase Elem, falseCase Elem) Elem
- func XFOR(iterClosure []Elem) Elem
- func Xter(condition bool, trueCase string, falseCase string) string

## Elem Methods

- func (tr Elem) Att(s string) Elem
- func (tr Elem) Cls(s string) Elem
- func (tr Elem) Val(s string) Elem
- func (tr Elem) Add(c ...Elem) Elem
- func (tr Elem) Raw() Elem
- func (tr Elem) CreateResponse() []byte
- func (tr Elem) resolve() string

