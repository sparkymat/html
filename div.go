package html

func Div(children ...BodyNode) BodyNode {
	return BodyNode{
		name:     "div",
		children: children,
	}
}
