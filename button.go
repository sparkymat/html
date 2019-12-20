package html

import "github.com/sparkymat/html/button"

func Button(buttonType button.Type) BodyNode {
	return BodyNode{
		name: "button",
		attributes: map[string]string{
			"type": buttonType.String(),
		},
		attributeOrder: []string{"type"},
	}
}
