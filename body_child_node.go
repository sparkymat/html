package html

type bodyChildNode struct {
	name             string
	attributes       map[string]string
	attributeOrder   []string
	children         []bodyChildNode
	unsafeHTMLString string
}

func (bn bodyChildNode) Attributes(attrs map[string]string) bodyChildNode {
	copiedNode := bn
	copiedNode.attributes = attrs
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
