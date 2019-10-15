package html

type link struct {
	bodyNode
}

func A() link {
	return link{
		bodyNode{
			name: "a",
		},
	}
}

func (l link) Href(url string) link {
	copiedLink := l
	if copiedLink.bodyNode.attributes == nil {
		copiedLink.bodyNode.attributes = make(map[string]string)
	}
	copiedLink.bodyNode.attributes["href"] = url

	return copiedLink
}
