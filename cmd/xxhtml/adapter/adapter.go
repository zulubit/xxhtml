package adapter

import (
	"fmt"
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
	"html":       "x.Html",
	"!doctype":   "x.DOCTYPE", // Special handling for DOCTYPE
	"head":       "x.Head",
	"body":       "x.Body",
}

// ConvertNode converts an HTML node into a custom Go syntax using the x package.
func ConvertNode(n *html.Node) string {
	if n.Type == html.CommentNode {
		return ""
	}
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text == "" {
			return ""
		}
		return fmt.Sprintf("x.C(`%s`),", text)
	}

	// Handle the DOCTYPE separately
	if n.Type == html.DoctypeNode {
		return "x.DOCTYPE(),\n"
	}

	elem := ""
	// Check if the tag has a corresponding convenience function
	elemFunc, exists := tagToFunc[n.Data]
	if !exists {
		elemFunc = "x.E" // fallback to generic E() function
		elem = fmt.Sprintf(`%s("%s",`, elemFunc, n.Data)
	} else {
		elem = fmt.Sprintf("%s(", elemFunc)
	}

	// Handle the element's attributes
	if len(n.Attr) > 0 {
		var attrParts []string
		for _, attr := range n.Attr {
			if attr.Key == "class" {
				attrParts = append(attrParts, fmt.Sprintf(`x.Class("%s")`, attr.Val))
			} else {
				attrParts = append(attrParts, fmt.Sprintf(`x.Att("%s", "%s")`, attr.Key, attr.Val))
			}
		}
		// Join all attributes into a single string
		elem += strings.Join(attrParts, ", ") + ",\n"
	} else {
		elem += "\n"
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
		elem += "),\n"
	}

	return elem
}

// PrintNode prints the details of an HTML node for debugging.
func PrintNode(n *html.Node, indent int) {
	if n == nil {
		return
	}

	// Indent based on node depth
	fmt.Printf("%sNode Type: %d, Data: %s\n", strings.Repeat("  ", indent), n.Type, n.Data)
	for _, attr := range n.Attr {
		fmt.Printf("%s  Attribute: %s=\"%s\"\n", strings.Repeat("  ", indent+1), attr.Key, attr.Val)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		PrintNode(c, indent+1)
	}
}

// ParseFull parses the entire HTML document and returns the root node.
func ParseFull(htmlContent string) (*html.Node, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML document: %w", err)
	}
	return doc, nil
}

// ParseFragment parses an HTML fragment within the body element and returns the nodes.
func ParseFragment(htmlContent string) ([]*html.Node, error) {
	doc, err := html.ParseFragment(strings.NewReader(htmlContent), &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML fragment: %w", err)
	}
	return doc, nil
}
