package html

func Text(text string) BodyNode {
	return BodyNode{
		unsafeHTMLString: text,
	}
}
