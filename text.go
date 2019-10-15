package html

func Text(text string) bodyNode {
	return bodyNode{
		unsafeHTMLString: text,
	}
}
