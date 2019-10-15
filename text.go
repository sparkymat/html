package html

func Text(text string) bodyChildNode {
	return bodyChildNode{
		unsafeHTMLString: text,
	}
}
