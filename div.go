package html

func Div(children ...bodyChildNode) bodyChildNode {
	return bodyChildNode{
		name:     "div",
		children: children,
	}
}
