package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// map of HTML tag names to corresponding xx convenience functions
var tagToFunc = map[string]string{
	"div":      "xx.Div()",
	"span":     "xx.Span()",
	"p":        "xx.P()",
	"a":        "xx.A()",
	"img":      "xx.Img()",
	"h1":       "xx.H1()",
	"h2":       "xx.H2()",
	"h3":       "xx.H3()",
	"ul":       "xx.Ul()",
	"ol":       "xx.Ol()",
	"li":       "xx.Li()",
	"table":    "xx.Table()",
	"tr":       "xx.Tr()",
	"td":       "xx.Td()",
	"th":       "xx.Th()",
	"form":     "xx.Form()",
	"input":    "xx.Input()",
	"button":   "xx.Button()",
	"label":    "xx.Label()",
	"textarea": "xx.Textarea()",
}

// ConvertNode converts an HTML node into a custom Go syntax using the xx package.
func ConvertNode(n *html.Node) string {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text == "" {
			return ""
		}
		return fmt.Sprintf("xx.ERAW(`%s`)", text)
	}

	// Check if the tag has a corresponding convenience function
	elemFunc, exists := tagToFunc[n.Data]
	if !exists {
		elemFunc = fmt.Sprintf("xx.E(`%s`)", n.Data) // fallback to generic E() function
	}

	// Initialize the element with the tag name or the convenience function
	elem := elemFunc

	// Handle the element's attributes
	var attrs []string
	for _, attr := range n.Attr {
		if attr.Key == "class" {
			elem += fmt.Sprintf(".CLS(`%s`)", attr.Val)
		} else {
			attrs = append(attrs, fmt.Sprintf(`%s="%s"`, attr.Key, attr.Val))
		}
	}
	for _, attr := range attrs {
		elem += fmt.Sprintf(".ATT(`%s`)", attr)
	}

	// Collect text content and child nodes
	var content string
	var children []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			content += strings.TrimSpace(c.Data) + " "
		} else {
			children = append(children, ConvertNode(c))
		}
	}
	content = strings.TrimSpace(content)
	if content != "" {
		elem += fmt.Sprintf(".VAL(`%s`)", content)
	}

	// Add children to the element with line breaks
	if len(children) > 0 {
		elem += ".C(\n" + strings.Join(children, ",\n") + ",\n)"
	}

	return elem
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your HTML:")
	htmlContent, _ := reader.ReadString('\n')

	// Parse the HTML fragment
	doc, err := html.ParseFragment(strings.NewReader(htmlContent), &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing HTML fragment: %v\n", err)
		return
	}

	// Convert the HTML fragment to custom syntax
	var result strings.Builder
	for _, node := range doc {
		result.WriteString(ConvertNode(node))
	}

	fmt.Println("Generated Go code:")
	fmt.Println(result.String())
}
