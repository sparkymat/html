package html

import "strings"

type bodyNode struct {
	name             string
	attributes       map[string]string
	attributeOrder   []string
	children         []bodyNode
	unsafeHTMLString string
}

func (bn bodyNode) Class(classes ...string) bodyNode {
	copiedNode := bn
	if copiedNode.attributes == nil {
		copiedNode.attributes = make(map[string]string)
	}

	copiedNode.attributes["class"] = strings.Join(classes, " ")

	return copiedNode
}

func (bn bodyNode) Attributes(attrs map[string]string) bodyNode {
	copiedNode := bn
	copiedNode.attributes = attrs
	return copiedNode
}

func (bn bodyNode) Children(children ...bodyNode) bodyNode {
	copiedNode := bn
	copiedNode.children = children
	return copiedNode
}

func (bn bodyNode) String() string {
	return nodeAsString(bn)
}

func (bn bodyNode) getUnsafeHTMLString() string      { return bn.unsafeHTMLString }
func (bn bodyNode) getName() string                  { return bn.name }
func (bn bodyNode) getChildrenCount() int            { return len(bn.children) }
func (bn bodyNode) getChildAt(index int) node        { return bn.children[index] }
func (bn bodyNode) getAttributes() map[string]string { return bn.attributes }
func (bn bodyNode) getAttributeOrder() []string      { return bn.attributeOrder }
