package html

func Script(src string) BodyNode {
	return BodyNode{
		name: "script",
		attributes: map[string]string{
			"src": src,
		},
		attributeOrder: []string{
			"src",
		},
	}
}
