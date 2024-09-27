package xx

import (
	"testing"
)

// TestE tests the E function for creating HTML elements.
func TestE(t *testing.T) {
	elem := E("div", `class="container"`, "Hello, World!")
	expected := `<div class="container">Hello, World!</div>`
	if string(elem.Render()) != expected {
		t.Errorf("E() = %s; want %s", elem.Render(), expected)
	}
}

// TestERAW tests the ERAW function for creating raw HTML elements.
func TestERAW(t *testing.T) {
	elem := ERAW("<h1>Hello, World!</h1>")
	expected := `<h1>Hello, World!</h1>`
	if string(elem.Render()) != expected {
		t.Errorf("ERAW() = %s; want %s", elem.Render(), expected)
	}
}

// TestIF tests the IF function for conditional rendering of HTML elements.
func TestIF(t *testing.T) {
	elem := IF(true, E("span", "", "True"))
	expected := `<span>True</span>`
	if string(elem.Render()) != expected {
		t.Errorf("IF(true) = %s; want %s", elem.Render(), expected)
	}

	elem = IF(false, E("span", "", "True"))
	expected = ``
	if string(elem.Render()) != expected {
		t.Errorf("IF(false) = %s; want %s", elem.Render(), expected)
	}
}

// TestFOR tests the FOR function for creating a list of HTML elements.
func TestFOR(t *testing.T) {
	elems := FOR([]Elem{
		E("li", "", "Item 1"),
		E("li", "", "Item 2"),
	})
	expected := `<li>Item 1</li><li>Item 2</li>`
	if string(elems.Render()) != expected {
		t.Errorf("FOR() = %s; want %s", elems.Render(), expected)
	}
}

// TestTER tests the TER function for conditional string rendering.
func TestTER(t *testing.T) {
	elem := TER(true, E("p", "", "True"), E("p", "", "False"))
	expected := `<p>True</p>`
	if string(elem.Render()) != expected {
		t.Errorf("TER(true) = %s; want %s", elem.Render(), expected)
	}

	elem = TER(false, E("p", "", "True"), E("p", "", "False"))
	expected = `<p>False</p>`
	if string(elem.Render()) != expected {
		t.Errorf("TER(false) = %s; want %s", elem.Render(), expected)
	}
}

// TestSTER tests the STER function for conditional string rendering.
func TestSTER(t *testing.T) {
	result := STER(true, "True", "False")
	expected := "True"
	if result != expected {
		t.Errorf("STER(true) = %s; want %s", result, expected)
	}

	result = STER(false, "True", "False")
	expected = "False"
	if result != expected {
		t.Errorf("STER(false) = %s; want %s", result, expected)
	}
}

// TestSIF tests the SIF function for conditional string rendering.
func TestSIF(t *testing.T) {
	result := SIF(true, "True")
	expected := "True"
	if result != expected {
		t.Errorf("SIF(true) = %s; want %s", result, expected)
	}

	result = SIF(false, "True")
	expected = ""
	if result != expected {
		t.Errorf("SIF(false) = %s; want %s", result, expected)
	}
}

// TestChildrenRendering tests the rendering of HTML elements with child elements.
func TestChildrenRendering(t *testing.T) {
	// Create an element with children
	parentElem := E("div", `class="parent"`, "",
		E("span", `class="child"`, "Child 1"),
		E("span", `class="child"`, "Child 2"),
	)

	// Expected HTML output
	expected := `<div class="parent"><span class="child">Child 1</span><span class="child">Child 2</span></div>`

	// Check if the rendered output matches the expected output
	if string(parentElem.Render()) != expected {
		t.Errorf("ChildrenRendering() = %s; want %s", parentElem.Render(), expected)
	}
}
