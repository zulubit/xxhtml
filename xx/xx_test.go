package xx

import (
	"testing"
)

// TestE tests the E function for creating HTML elements.
func TestE(t *testing.T) {
	elem := E("div").CLS("container").VAL("Hello, World!")
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
	elem := IF(true, E("span").VAL("True"))
	expected := `<span>True</span>`
	if string(elem.Render()) != expected {
		t.Errorf("IF(true) = %s; want %s", elem.Render(), expected)
	}

	elem = IF(false, E("span").VAL("True"))
	expected = ``
	if string(elem.Render()) != expected {
		t.Errorf("IF(false) = %s; want %s", elem.Render(), expected)
	}
}

// TestFOR tests the FOR function for creating a list of HTML elements.
func TestFOR(t *testing.T) {
	elems := FOR([]Elem{
		E("li").VAL("Item 1"),
		E("li").VAL("Item 2"),
	})
	expected := `<li>Item 1</li><li>Item 2</li>`
	if string(elems.Render()) != expected {
		t.Errorf("FOR() = %s; want %s", elems.Render(), expected)
	}
}

// TestTER tests the TER function for conditional rendering of HTML elements.
func TestTER(t *testing.T) {
	elem := TER(true, E("p").VAL("True"), E("p").VAL("False"))
	expected := `<p>True</p>`
	if string(elem.Render()) != expected {
		t.Errorf("TER(true) = %s; want %s", elem.Render(), expected)
	}

	elem = TER(false, E("p").VAL("True"), E("p").VAL("False"))
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
	parentElem := E("div").CLS("parent").
		C(E("span").CLS("child").VAL("Child 1"), E("span").CLS("child").VAL("Child 2"))

	// Expected HTML output
	expected := `<div class="parent"><span class="child">Child 1</span><span class="child">Child 2</span></div>`

	// Check if the rendered output matches the expected output
	if string(parentElem.Render()) != expected {
		t.Errorf("ChildrenRendering() = %s; want %s", parentElem.Render(), expected)
	}
}

// TestDiv tests the Div function.
func TestDiv(t *testing.T) {
	expected := "<div></div>"
	result := Div().resolve()
	if result != expected {
		t.Errorf("Div() = %s; want %s", result, expected)
	}
}

// TestSpan tests the Span function.
func TestSpan(t *testing.T) {
	expected := "<span></span>"
	result := Span().resolve()
	if result != expected {
		t.Errorf("Span() = %s; want %s", result, expected)
	}
}

// TestP tests the P function.
func TestP(t *testing.T) {
	expected := "<p></p>"
	result := P().resolve()
	if result != expected {
		t.Errorf("P() = %s; want %s", result, expected)
	}
}

// TestA tests the A function.
func TestA(t *testing.T) {
	expected := `<a href="https://example.com" >Click here</a>`
	result := A().ATT(`href="https://example.com"`).VAL("Click here").resolve()
	if result != expected {
		t.Errorf("A() = %s; want %s", result, expected)
	}
}

// TestImg tests the Img function.
func TestImg(t *testing.T) {
	expected := `<img src="image.png" alt="Image description" ></img>`
	result := Img().ATT(`src="image.png"`).ATT(`alt="Image description"`).resolve()
	if result != expected {
		t.Errorf("Img() = %s; want %s", result, expected)
	}
}

// TestH1 tests the H1 function.
func TestH1(t *testing.T) {
	expected := "<h1>Heading 1</h1>"
	result := H1().VAL("Heading 1").resolve()
	if result != expected {
		t.Errorf("H1() = %s; want %s", result, expected)
	}
}

// TestH2 tests the H2 function.
func TestH2(t *testing.T) {
	expected := "<h2>Heading 2</h2>"
	result := H2().VAL("Heading 2").resolve()
	if result != expected {
		t.Errorf("H2() = %s; want %s", result, expected)
	}
}

// TestH3 tests the H3 function.
func TestH3(t *testing.T) {
	expected := "<h3>Heading 3</h3>"
	result := H3().VAL("Heading 3").resolve()
	if result != expected {
		t.Errorf("H3() = %s; want %s", result, expected)
	}
}

// TestUl tests the Ul function.
func TestUl(t *testing.T) {
	expected := "<ul><li>Item 1</li><li>Item 2</li></ul>"
	result := Ul().C(Li().VAL("Item 1"), Li().VAL("Item 2")).resolve()
	if result != expected {
		t.Errorf("Ul() = %s; want %s", result, expected)
	}
}

// TestOl tests the Ol function.
func TestOl(t *testing.T) {
	expected := "<ol><li>First</li><li>Second</li></ol>"
	result := Ol().C(Li().VAL("First"), Li().VAL("Second")).resolve()
	if result != expected {
		t.Errorf("Ol() = %s; want %s", result, expected)
	}
}

// TestLi tests the Li function.
func TestLi(t *testing.T) {
	expected := "<li>List item</li>"
	result := Li().VAL("List item").resolve()
	if result != expected {
		t.Errorf("Li() = %s; want %s", result, expected)
	}
}

// TestTable tests the Table function.
func TestTable(t *testing.T) {
	expected := "<table></table>"
	result := Table().resolve()
	if result != expected {
		t.Errorf("Table() = %s; want %s", result, expected)
	}
}

// TestTr tests the Tr function.
func TestTr(t *testing.T) {
	expected := "<tr><td>Cell 1</td><td>Cell 2</td></tr>"
	result := Tr().C(Td().VAL("Cell 1"), Td().VAL("Cell 2")).resolve()
	if result != expected {
		t.Errorf("Tr() = %s; want %s", result, expected)
	}
}

// TestTd tests the Td function.
func TestTd(t *testing.T) {
	expected := "<td>Cell content</td>"
	result := Td().VAL("Cell content").resolve()
	if result != expected {
		t.Errorf("Td() = %s; want %s", result, expected)
	}
}

// TestTh tests the Th function.
func TestTh(t *testing.T) {
	expected := "<th>Header</th>"
	result := Th().VAL("Header").resolve()
	if result != expected {
		t.Errorf("Th() = %s; want %s", result, expected)
	}
}

// TestForm tests the Form function.
func TestForm(t *testing.T) {
	expected := `<form action="/submit" method="post" ></form>`
	result := Form().ATT(`action="/submit"`).ATT(`method="post"`).resolve()
	if result != expected {
		t.Errorf("Form() = %s; want %s", result, expected)
	}
}

// TestInput tests the Input function.
func TestInput(t *testing.T) {
	expected := `<input type="text" placeholder="Enter text" ></input>`
	result := Input().ATT(`type="text"`).ATT(`placeholder="Enter text"`).resolve()
	if result != expected {
		t.Errorf("Input() = %s; want %s", result, expected)
	}
}

// TestButton tests the Button function.
func TestButton(t *testing.T) {
	expected := "<button>Click me</button>"
	result := Button().VAL("Click me").resolve()
	if result != expected {
		t.Errorf("Button() = %s; want %s", result, expected)
	}
}

// TestLabel tests the Label function.
func TestLabel(t *testing.T) {
	expected := `<label for="inputID" >Label text</label>`
	result := Label().ATT(`for="inputID"`).VAL("Label text").resolve()
	if result != expected {
		t.Errorf("Label() = %s; want %s", result, expected)
	}
}
