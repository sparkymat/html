package html

type headChildNode struct {
	node
}

// Title returns a <title> tag with the specified text
func Title(text string) headChildNode {
	return headChildNode{
		standardNode{
			name: "title",
			children: []node{
				Text(text),
			},
		},
	}
}

func Base(href string, target string) headChildNode {
	return headChildNode{
		standardNode{
			name: "base",
			attributes: map[string]string{
				"href":   href,
				"target": target,
			},
			attributeOrder: []string{"href", "target"},
		},
	}
}

func Style(rules string) headChildNode {
	return headChildNode{
		standardNode{
			name: "style",
			children: []node{
				Text(rules),
			},
		},
	}
}

func Link(href string) headChildNode {
	return headChildNode{
		standardNode{
			name: "link",
			attributes: map[string]string{
				"rel":  "stylesheet",
				"href": href,
			},
			attributeOrder: []string{"rel", "href"},
		},
	}
}

func ScriptInline(body string) headChildNode {
	return headChildNode{
		standardNode{
			name: "script",
			children: []node{
				Text(body),
			},
		},
	}
}

func ScriptExternal(href string) headChildNode {
	return headChildNode{
		standardNode{
			name: "script",
			attributes: map[string]string{
				"src": href,
			},
		},
	}
}
