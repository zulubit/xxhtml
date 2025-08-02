package x

import (
	"bytes"
	"testing"
)

func TestElem_Render(t *testing.T) {
	tests := []struct {
		name     string
		elem     Elem
		expected string
	}{
		{
			name:     "Div with text",
			elem:     Div(C("Hello, World!")),
			expected: "<div>Hello, World!</div>",
		},
		{
			name:     "Span with Class",
			elem:     Span(Class("highlight"), C("Highlighted text")),
			expected: `<span class="highlight">Highlighted text</span>`,
		},
		{
			name:     "Anchor with href",
			elem:     A(Att("href", "https://example.com"), C("Example")),
			expected: `<a href="https://example.com">Example</a>`,
		},
		{
			name:     "Image with Attributes (self-closing)",
			elem:     Img(Att("src", "image.png"), Att("alt", "An image")).SelfClose(),
			expected: `<img src="image.png" alt="An image" />`,
		},
		{
			name:     "Nested elements",
			elem:     Div(Class("container"), Span(C("Nested span"))),
			expected: `<div class="container"><span>Nested span</span></div>`,
		},
		{
			name: "Large HTML Document",
			elem: Html(
				Head(
					Title(C("Large Document Title")),
					Meta(Att("charset", "UTF-8")).SelfClose(),
					Link(Att("rel", "stylesheet"), Att("href", "styles.css")).SelfClose(),
					Script(Att("src", "script.js")),
				),
				Body(
					Div(Class("header"), H1(C("Main Header"))),
					Div(Class("content"),
						P(C("This is a paragraph in a large HTML document.")),
						Div(Class("nested"),
							Span(C("Some nested content")),
							Ol(
								Li(Class("item1"), C("List item 1")),
								Li(Class("item2"), C("List item 2")),
							),
						),
					),
					Footer(C("Footer content")),
				),
			),
			expected: `<html><head><title>Large Document Title</title><meta charset="UTF-8" /><link rel="stylesheet" href="styles.css" /><script src="script.js"></script></head><body><div class="header"><h1>Main Header</h1></div><div class="content"><p>This is a paragraph in a large HTML document.</p><div class="nested"><span>Some nested content</span><ol><li class="item1">List item 1</li><li class="item2">List item 2</li></ol></div></div><footer>Footer content</footer></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tt.elem.Render(&buf)
			if err != nil {
				t.Fatalf("Render() returned an error: %v", err)
			}
			result := buf.String()
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestIF(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		trueCase  Elem
		expected  string
	}{
		{
			name:      "Condition true",
			condition: true,
			trueCase:  Div(C("Condition is true")),
			expected:  "<div>Condition is true</div>",
		},
		{
			name:      "Condition false",
			condition: false,
			trueCase:  Div(C("Condition is true")),
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := IF(tt.condition, tt.trueCase).Render(&buf)
			if err != nil {
				t.Fatalf("Render() returned an error: %v", err)
			}
			result := buf.String()
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// Similar updates should be applied to the other tests (TestTER, TestFOR, TestSTERSIF, and TestSIF).
// Below are examples of how to handle them.

func TestTER(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		trueCase  Elem
		falseCase Elem
		expected  string
	}{
		{
			name:      "Condition true",
			condition: true,
			trueCase:  Div(C("True case")),
			falseCase: Div(C("False case")),
			expected:  "<div>True case</div>",
		},
		{
			name:      "Condition false",
			condition: false,
			trueCase:  Div(C("True case")),
			falseCase: Div(C("False case")),
			expected:  "<div>False case</div>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := TER(tt.condition, tt.trueCase, tt.falseCase).Render(&buf)
			if err != nil {
				t.Fatalf("Render() returned an error: %v", err)
			}
			result := buf.String()
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
