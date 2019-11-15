package html

import "strings"

func (bn BodyNode) setAttribute(key string, value string) BodyNode {
	copiedNode := bn
	if copiedNode.attributes == nil {
		copiedNode.attributes = make(map[string]string)
	}
	copiedNode.attributes[key] = value
	copiedNode.attributeOrder = moveOrAppendToEnd(copiedNode.attributeOrder, key)
	return copiedNode
}

func (bn BodyNode) removeAttribute(key string) BodyNode {
	copiedNode := bn
	if copiedNode.attributes == nil {
		copiedNode.attributes = make(map[string]string)
	}
	delete(copiedNode.attributes, key)
	return copiedNode
}

func (bn BodyNode) Class(classes ...string) BodyNode {
	return bn.setAttribute("class", strings.Join(classes, " "))
}

func (bn BodyNode) Href(href string) BodyNode {
	return bn.setAttribute("href", href)
}

func (bn BodyNode) For(href string) BodyNode {
	return bn.setAttribute("for", href)
}

func (bn BodyNode) Placeholder(placeholder string) BodyNode {
	return bn.setAttribute("placeholder", placeholder)
}

func (bn BodyNode) Value(value string) BodyNode {
	return bn.setAttribute("value", value)
}

func (bn BodyNode) Accept(types string) BodyNode {
	return bn.setAttribute("accept", types)
}

func (bn BodyNode) Alt(text string) BodyNode {
	return bn.setAttribute("alt", text)
}

func (bn BodyNode) Name(text string) BodyNode {
	return bn.setAttribute("name", text)
}

func (bn BodyNode) Pattern(text string) BodyNode {
	return bn.setAttribute("pattern", text)
}

func (bn BodyNode) Autofocus() BodyNode {
	return bn.setAttribute("autofocus", "autofocus")
}

func (bn BodyNode) AutoComplete(enabled bool) BodyNode {
	value := map[bool]string{
		false: "on",
		true:  "off",
	}[enabled]
	return bn.setAttribute("autocomplete", value)
}

func (bn BodyNode) Checked(checked bool) BodyNode {
	if checked {
		return bn.setAttribute("checked", "checked")
	}

	return bn.removeAttribute("checked")
}

func (bn BodyNode) Disabled(disabled bool) BodyNode {
	if disabled {
		return bn.setAttribute("disabled", "disabled")
	}

	return bn.removeAttribute("disabled")
}

func (bn BodyNode) Readonly(readonly bool) BodyNode {
	if readonly {
		return bn.setAttribute("readonly", "readonly")
	}

	return bn.removeAttribute("readonly")
}

func (bn BodyNode) Required(required bool) BodyNode {
	if required {
		return bn.setAttribute("required", "required")
	}

	return bn.removeAttribute("required")
}

func (bn BodyNode) Multiple(multiple bool) BodyNode {
	if multiple {
		return bn.setAttribute("multiple", "multiple")
	}

	return bn.removeAttribute("multiple")
}

func (bn BodyNode) Max(value string) BodyNode {
	return bn.setAttribute("max", value)
}

func (bn BodyNode) Min(value string) BodyNode {
	return bn.setAttribute("min", value)
}

func (bn BodyNode) Src(value string) BodyNode {
	return bn.setAttribute("src", value)
}

func (bn BodyNode) Width(value string) BodyNode {
	return bn.setAttribute("width", value)
}

func (bn BodyNode) Height(value string) BodyNode {
	return bn.setAttribute("height", value)
}

func (bn BodyNode) Poster(value string) BodyNode {
	return bn.setAttribute("poster", value)
}

func (bn BodyNode) Id(value string) BodyNode {
	return bn.setAttribute("id", value)
}
