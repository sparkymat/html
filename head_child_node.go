package html

type headChildNode struct {
	node
}

// Title returns a <title> tag with the specified text
func Title(text string) headChildNode {
	return headChildNode{
		standardNode{
			name: "title",
			children: []node{
				Text(text),
			},
		},
	}
}
