package html

func H1(children ...bodyNode) bodyNode {
	return bodyNode{
		name:     "h1",
		children: children,
	}
}

func P(body string) bodyNode {
	return bodyNode{
		name: "p",
		children: []bodyNode{
			bodyNode{
				unsafeHTMLString: body,
			},
		},
	}
}
