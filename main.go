package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// ConvertNode converts an HTML node into a custom Go syntax using the xx package.
func ConvertNode(n *html.Node) string {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text == "" {
			return ""
		}
		return fmt.Sprintf("xx.ERAW(`%s`)", text)
	}

	// Initialize the element with the tag name
	elem := fmt.Sprintf("xx.E(`%s`)", n.Data)

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
	for _, child := range children {
		elem += fmt.Sprintf(".C(\n%s\n)", child)
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
