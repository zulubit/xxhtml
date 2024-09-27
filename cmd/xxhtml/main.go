package main

import (
	"bufio"
	"flag"
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
		elem = fmt.Sprintf(`%s("%s",&x.X{`, elemFunc, n.Data)
	} else {
		elem = fmt.Sprintf("%s(x.X{", elemFunc)
	}

	// Handle the element's attributes
	if len(n.Attr) > 0 {
		var attrParts []string
		for _, attr := range n.Attr {
			attrParts = append(attrParts, fmt.Sprintf(`%s="%s"`, attr.Key, attr.Val))
		}
		// Join all attributes into a single string
		allAttrs := strings.Join(attrParts, " ")
		elem += fmt.Sprintf("Att: `%s`},\n", allAttrs)
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

// parseFull parses the entire HTML document and returns the root node.
func parseFull(htmlContent string) (*html.Node, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML document: %w", err)
	}
	return doc, nil
}

// parseFragment parses an HTML fragment within the body element and returns the nodes.
func parseFragment(htmlContent string) ([]*html.Node, error) {
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

func initCli() {
	// Define the --full flag
	fullFlag := flag.Bool("full", false, "Parse the entire HTML document")
	flag.Parse()

	// Read HTML input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your HTML:")
	htmlContent, _ := reader.ReadString('\n')

	// Determine the parse mode based on the flag
	var nodes []string

	if *fullFlag {
		doc, parseErr := parseFull(htmlContent)
		if parseErr != nil {
			fmt.Fprintf(os.Stderr, "%v\n", parseErr)
			return
		}
		// fmt.Println("Node structure for full document:")
		// PrintNode(doc, 0)
		// Convert the entire document, treat it as a root node
		nodes = []string{ConvertNode(doc)}
	} else {
		doc, parseErr := parseFragment(htmlContent)
		if parseErr != nil {
			fmt.Fprintf(os.Stderr, "%v\n", parseErr)
			return
		}
		// Convert each fragment node
		for _, node := range doc {
			nodes = append(nodes, ConvertNode(node))
		}
	}

	// Convert the HTML content to custom syntax
	var result strings.Builder
	for _, node := range nodes {
		result.WriteString(node)
	}
	s := result.String()
	fmt.Println("\033[31mGenerated Go code:\033[0m")
	// Remove trailing comma for proper syntax
	if len(s) > 0 && s[len(s)-1] == ',' {
		s = s[:len(s)-1]
	}
	fmt.Println(s)
}

func main() {
	initCli()
}
