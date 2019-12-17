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

func Select() BodyNode {
	return BodyNode{
		name: "select",
	}
}

func Option(label string, value string) BodyNode {
	return BodyNode{
		name: "option",
		attributes: map[string]string{
			"value": value,
		},
		attributeOrder: []string{"value"},
		children: []BodyNode{
			Text(label),
		},
	}
}

func SelectWithOptions(name string, multiple bool, options map[string]string, currentValue string) BodyNode {
	optionNodes := []BodyNode{}

	for value, label := range options {
		optionNodes = append(optionNodes, Option(label, value).Selected(currentValue == value))
	}

	return Select().Name(name).Multiple(multiple).Children(optionNodes...)
}
