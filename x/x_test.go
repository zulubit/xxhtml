package x

import (
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
			elem:     Img(Att("src", "image.png"), Att("alt", "An image")),
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
					Meta(Att("charset", "UTF-8")),
					Link(Att("rel", "stylesheet"), Att("href", "styles.css")),
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
			result := string(tt.elem.Render())
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestIF(t *testing.T) {
	tests := []struct {
		condition bool
		trueCase  Elem
		expected  string
	}{
		{
			condition: true,
			trueCase:  Div(C("Condition is true")),
			expected:  "<div>Condition is true</div>",
		},
		{
			condition: false,
			trueCase:  Div(C("Condition is true")),
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := string(IF(tt.condition, tt.trueCase).Render())
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestTER(t *testing.T) {
	tests := []struct {
		condition bool
		trueCase  Elem
		falseCase Elem
		expected  string
	}{
		{
			condition: true,
			trueCase:  Div(C("True case")),
			falseCase: Div(C("False case")),
			expected:  "<div>True case</div>",
		},
		{
			condition: false,
			trueCase:  Div(C("True case")),
			falseCase: Div(C("False case")),
			expected:  "<div>False case</div>",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := string(TER(tt.condition, tt.trueCase, tt.falseCase).Render())
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestFOR(t *testing.T) {
	tests := []struct {
		elements []Elem
		expected []string
	}{
		{
			elements: []Elem{
				Div(C("Item 1")),
				Div(C("Item 2")),
			},
			expected: []string{
				"<div>Item 1</div>",
				"<div>Item 2</div>",
			},
		},
	}

	for i, tt := range tests {
		t.Run("", func(t *testing.T) {
			results := FOR(tt.elements)
			for j, result := range results {
				if string(result.Render()) != tt.expected[j] {
					t.Errorf("test %d, expected %q, got %q", i, tt.expected[j], result.Render())
				}
			}
		})
	}
}

func TestSTERSIF(t *testing.T) {
	tests := []struct {
		condition bool
		trueCase  string
		falseCase string
		expected  string
	}{
		{
			condition: true,
			trueCase:  "True",
			falseCase: "False",
			expected:  "True",
		},
		{
			condition: false,
			trueCase:  "True",
			falseCase: "False",
			expected:  "False",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := STER(tt.condition, tt.trueCase, tt.falseCase)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestSIF(t *testing.T) {
	tests := []struct {
		condition bool
		trueCase  string
		expected  string
	}{
		{
			condition: true,
			trueCase:  "True",
			expected:  "True",
		},
		{
			condition: false,
			trueCase:  "True",
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := SIF(tt.condition, tt.trueCase)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
