package html

import "github.com/sparkymat/html/input"

func Label(text string) bodyNode { return nodeWithText("label", text) }

func Input(inputType input.Type) bodyNode {
	return bodyNode{
		name: "input",
		attributes: map[string]string{
			"type": inputType.String(),
		},
	}
}
