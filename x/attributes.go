package x

// Class creates an Elem representing a CSS class.
func Class(classes string) Elem {
	return Elem{
		Type:    AttributeNode,
		AttrKey: "class",
		AttrVal: classes,
	}
}
