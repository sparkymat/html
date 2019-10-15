package html

func H1(children ...bodyChildNode) bodyChildNode {
	return bodyChildNode{
		name:     "h1",
		children: children,
	}
}

func P(body string) bodyChildNode {
	return bodyChildNode{
		name: "p",
		children: []bodyChildNode{
			bodyChildNode{
				unsafeHTMLString: body,
			},
		},
	}
}
