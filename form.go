package html

import (
	"github.com/sparkymat/html/input"
	"github.com/sparkymat/html/method"
	"github.com/sparkymat/html/selectoption"
)

func Label(text string) BodyNode { return nodeWithText("label", text) }

func LabelWithChildren(children ...BodyNode) BodyNode {
	return BodyNode{
		name:     "label",
		children: children,
	}
}

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

func Option(option selectoption.Type) BodyNode {
	return BodyNode{
		name: "option",
		attributes: map[string]string{
			"value": option.Value,
		},
		attributeOrder: []string{"value"},
		children: []BodyNode{
			Text(option.Label),
		},
	}
}

func SelectWithOptions(name string, multiple bool, options []selectoption.Type, currentValue string) BodyNode {
	optionNodes := []BodyNode{}

	for _, option := range options {
		optionNodes = append(optionNodes, Option(option).Selected(currentValue == option.Value))
	}

	return Select().Name(name).Multiple(multiple).Children(optionNodes...)
}
