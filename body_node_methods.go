package html

import "strings"

func (bn bodyNode) setAttribute(key string, value string) bodyNode {
	copiedNode := bn
	if copiedNode.attributes == nil {
		copiedNode.attributes = make(map[string]string)
	}
	copiedNode.attributes[key] = value
	copiedNode.attributeOrder = moveOrAppendToEnd(copiedNode.attributeOrder, key)
	return copiedNode
}

func (bn bodyNode) removeAttribute(key string) bodyNode {
	copiedNode := bn
	if copiedNode.attributes == nil {
		copiedNode.attributes = make(map[string]string)
	}
	delete(copiedNode.attributes, key)
	return copiedNode
}

func (bn bodyNode) Class(classes ...string) bodyNode {
	return bn.setAttribute("class", strings.Join(classes, " "))
}

func (bn bodyNode) Href(href string) bodyNode {
	return bn.setAttribute("href", href)
}

func (bn bodyNode) For(href string) bodyNode {
	return bn.setAttribute("for", href)
}

func (bn bodyNode) Placeholder(placeholder string) bodyNode {
	return bn.setAttribute("placeholder", placeholder)
}

func (bn bodyNode) Value(value string) bodyNode {
	return bn.setAttribute("value", value)
}

func (bn bodyNode) Accept(types string) bodyNode {
	return bn.setAttribute("accept", types)
}

func (bn bodyNode) Alt(text string) bodyNode {
	return bn.setAttribute("alt", text)
}

func (bn bodyNode) Name(text string) bodyNode {
	return bn.setAttribute("name", text)
}

func (bn bodyNode) Pattern(text string) bodyNode {
	return bn.setAttribute("pattern", text)
}

func (bn bodyNode) Autofocus() bodyNode {
	return bn.setAttribute("autofocus", "autofocus")
}

func (bn bodyNode) AutoComplete(enabled bool) bodyNode {
	value := map[bool]string{
		false: "on",
		true:  "off",
	}[enabled]
	return bn.setAttribute("autocomplete", value)
}

func (bn bodyNode) Checked(checked bool) bodyNode {
	if checked {
		return bn.setAttribute("checked", "checked")
	}

	return bn.removeAttribute("checked")
}

func (bn bodyNode) Disabled(disabled bool) bodyNode {
	if disabled {
		return bn.setAttribute("disabled", "disabled")
	}

	return bn.removeAttribute("disabled")
}

func (bn bodyNode) Readonly(readonly bool) bodyNode {
	if readonly {
		return bn.setAttribute("readonly", "readonly")
	}

	return bn.removeAttribute("readonly")
}

func (bn bodyNode) Required(required bool) bodyNode {
	if required {
		return bn.setAttribute("required", "required")
	}

	return bn.removeAttribute("required")
}

func (bn bodyNode) Multiple(multiple bool) bodyNode {
	if multiple {
		return bn.setAttribute("multiple", "multiple")
	}

	return bn.removeAttribute("multiple")
}

func (bn bodyNode) Max(value string) bodyNode {
	return bn.setAttribute("max", value)
}

func (bn bodyNode) Min(value string) bodyNode {
	return bn.setAttribute("min", value)
}
