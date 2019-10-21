package html

import "github.com/sparkymat/html/input"

func Label(text string) BodyNode { return nodeWithText("label", text) }

func Input(inputType input.Type) BodyNode {
	return BodyNode{
		name: "input",
		attributes: map[string]string{
			"type": inputType.String(),
		},
	}
}
