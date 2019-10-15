package html

func Div(children ...bodyNode) bodyNode {
	return bodyNode{
		name:     "div",
		children: children,
	}
}
