package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// map of HTML tag names to corresponding x convenience functions
var tagToFunc = map[string]string{
	"div":        "x.Div",
	"span":       "x.Span",
	"p":          "x.P",
	"a":          "x.A",
	"img":        "x.Img",
	"h1":         "x.H1",
	"h2":         "x.H2",
	"h3":         "x.H3",
	"ul":         "x.Ul",
	"ol":         "x.Ol",
	"li":         "x.Li",
	"table":      "x.Table",
	"tr":         "x.Tr",
	"td":         "x.Td",
	"th":         "x.Th",
	"form":       "x.Form",
	"input":      "x.Input",
	"button":     "x.Button",
	"label":      "x.Label",
	"article":    "x.Article",
	"aside":      "x.Aside",
	"header":     "x.Header",
	"footer":     "x.Footer",
	"main":       "x.Main",
	"section":    "x.Section",
	"nav":        "x.Nav",
	"figure":     "x.Figure",
	"figcaption": "x.Figcaption",
	"datalist":   "x.Datalist",
	"option":     "x.Option",
	"details":    "x.Details",
	"summary":    "x.Summary",
	"dialog":     "x.Dialog",
	"embed":      "x.Embed",
	"map":        "x.Map",
	"area":       "x.Area",
	"source":     "x.Source",
	"track":      "x.Track",
	"param":      "x.Param",
	"script":     "x.Script",
	"style":      "x.Style",
	"meta":       "x.Meta",
	"link":       "x.Link",
	"title":      "x.Title",
	"base":       "x.Base",
}

// ConvertNode converts an HTML node into a custom Go syntax using the x package.

func ConvertNode(n *html.Node) string {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text == "" {
			return ""
		}
		return fmt.Sprintf("x.C(`%s`),", text)
	}

	elem := ""
	// Check if the tag has a corresponding convenience function
	elemFunc, exists := tagToFunc[n.Data]
	if !exists {
		elemFunc = "x.E" // fallback to generic E() function
		elem = fmt.Sprintf(`%s("%s",x.X{`, elemFunc, n.Data)
	} else {
		elem = fmt.Sprintf("%s(x.X{", elemFunc)
	}

	// Handle the element's attributes
	if len(n.Attr) > 0 {
		var attrParts []string
		for _, attr := range n.Attr {
			if attr.Key == "class" {
				attrParts = append(attrParts, fmt.Sprintf("Class: `%s`", attr.Val))
			} else if attr.Key == "id" {
				attrParts = append(attrParts, fmt.Sprintf("Id: `%s`", attr.Val))
			} else {
				attrParts = append(attrParts, fmt.Sprintf("Att: `%s=\"%s\"`", attr.Key, attr.Val))
			}
		}
		elem += strings.Join(attrParts, ", ") + "},\n"
	} else {
		elem += "},\n"
	}

	// Collect child nodes
	var children []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		child := ConvertNode(c)
		if child != "" { // Avoid appending empty strings
			children = append(children, child)
		}
	}

	// Add children to the element
	if len(children) > 0 {
		elem += strings.Join(children, "\n") + "\n),"
	} else {
		elem += ")"
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
	s := result.String()
	fmt.Println("\033[31mGenerated Go code:\033[0m")
	fmt.Println(s[:len(s)-1])
}
