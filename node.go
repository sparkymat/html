package html

import "fmt"

type node interface {
	String() string
}

type textNode string

func (tn textNode) String() string {
	return string(tn)
}

type standardNode struct {
	name           string
	attributes     map[string]string
	attributeOrder []string
	children       []node
}

func (n standardNode) String() string {
	var closeTag string

	openTag := fmt.Sprintf("<%s", n.name)
	hasCloseTag := true
	multiline := true

	if len(n.children) == 0 {
		hasCloseTag = false
		multiline = false
	}
	if len(n.children) == 1 {
		if _, isText := n.children[0].(textNode); isText {
			multiline = false
		}
	}

	if len(n.attributes) > 0 {
		if len(n.attributeOrder) > 0 {
			for _, key := range n.attributeOrder {
				value := n.attributes[key]
				openTag = fmt.Sprintf("%s %s=\"%s\"", openTag, key, value)
			}
		} else {
			for key, value := range n.attributes {
				openTag = fmt.Sprintf("%s %s=\"%s\"", openTag, key, value)
			}
		}
	}

	if multiline {
		closeTag = fmt.Sprintf("\n")
	}

	if hasCloseTag {
		openTag = fmt.Sprintf("%s>", openTag)
		closeTag = fmt.Sprintf("%s</%s>", closeTag, n.name)
	} else {
		closeTag = fmt.Sprintf("%s />", closeTag)
	}
	var body string
	for _, child := range n.children {
		body = fmt.Sprintf("%s%s", body, child.String())
	}
	return fmt.Sprintf("%s%s%s\n", openTag, body, closeTag)
}

func (n standardNode) Attributes(attrs map[string]string) standardNode {
	copiedNode := n
	copiedNode.attributes = attrs
	return copiedNode
}
