package html

type link struct {
	bodyChildNode
}

func A() link {
	return link{
		bodyChildNode{
			name: "a",
		},
	}
}

func (l link) Href(url string) link {
	copiedLink := l
	if copiedLink.bodyChildNode.attributes == nil {
		copiedLink.bodyChildNode.attributes = make(map[string]string)
	}
	copiedLink.bodyChildNode.attributes["href"] = url

	return copiedLink
}
