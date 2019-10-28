package html

import (
	"github.com/sparkymat/html/input"
	"github.com/sparkymat/html/method"
)

func Label(text string) BodyNode { return nodeWithText("label", text) }

func Form(action string, m method.Type) BodyNode {
	return BodyNode{
		name: "form",
		attributes: map[string]string{
			"action": action,
			"method": m.String(),
		},
		attributeOrder: []string{"action", "method"},
	}
}

func Input(inputType input.Type) BodyNode {
	return BodyNode{
		name: "input",
		attributes: map[string]string{
			"type": inputType.String(),
		},
		attributeOrder: []string{"type"},
	}
}
