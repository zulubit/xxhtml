package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func ConvertNode(n *html.Node) string {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text == "" {
			// If the text is empty after trimming, return an empty string
			return ""
		}
		return fmt.Sprintf("`%s`", strings.TrimSpace(n.Data))
	}

	// Handle the element and its attributes
	var attrs string
	for _, attr := range n.Attr {
		attrs += fmt.Sprintf(`%s="%s" `, attr.Key, attr.Val)
	}
	attrs = strings.TrimSpace(attrs)

	var children []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// Handle both text nodes and element nodes
		child := strings.TrimSpace(ConvertNode(c))
		if child != "" {
			children = append(children, child)
		}
	}

	// Construct the Go code for this element
	element := fmt.Sprintf(`xx.E("%s", %s, %s)`, n.Data, fmt.Sprintf("`%s`", attrs), strings.Join(children, ", "))
	return element
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
	var result string
	for _, node := range doc {
		result += ConvertNode(node)
	}

	fmt.Println("Generated Go code:")
	fmt.Println(result)
}
