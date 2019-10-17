package html

func Script(src string) bodyNode {
	return bodyNode{
		name: "script",
		attributes: map[string]string{
			"src": src,
		},
	}
}
