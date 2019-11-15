package html

func Text(text string) BodyNode {
	return BodyNode{
		unsafeHTMLString: text,
	}
}

func Span() BodyNode {
	return BodyNode{
		name: "span",
	}
}
