package html

import "strings"

type bodyChildNode struct {
	name             string
	attributes       map[string]string
	attributeOrder   []string
	children         []bodyChildNode
	unsafeHTMLString string
}

func (bn bodyChildNode) Class(classes ...string) bodyChildNode {
	copiedNode := bn
	if copiedNode.attributes == nil {
		copiedNode.attributes = make(map[string]string)
	}

	copiedNode.attributes["class"] = strings.Join(classes, " ")

	return copiedNode
}

func (bn bodyChildNode) Attributes(attrs map[string]string) bodyChildNode {
	copiedNode := bn
	copiedNode.attributes = attrs
	return copiedNode
}

func (bn bodyChildNode) Children(children ...bodyChildNode) bodyChildNode {
	copiedNode := bn
	copiedNode.children = children
	return copiedNode
}

func (bn bodyChildNode) String() string {
	return nodeAsString(bn)
}

func (bn bodyChildNode) getUnsafeHTMLString() string      { return bn.unsafeHTMLString }
func (bn bodyChildNode) getName() string                  { return bn.name }
func (bn bodyChildNode) getChildrenCount() int            { return len(bn.children) }
func (bn bodyChildNode) getChildAt(index int) node        { return bn.children[index] }
func (bn bodyChildNode) getAttributes() map[string]string { return bn.attributes }
func (bn bodyChildNode) getAttributeOrder() []string      { return bn.attributeOrder }
