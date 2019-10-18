package html

type bodyNode struct {
	name             string
	attributes       map[string]string
	attributeOrder   []string
	children         []bodyNode
	unsafeHTMLString string
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

func nodeWithChildren(tag string, children []bodyNode) bodyNode {
	return bodyNode{
		name:     tag,
		children: children,
	}
}

func nodeWithText(tag string, text string) bodyNode {
	return bodyNode{
		name: tag,
		children: []bodyNode{
			Text(text),
		},
	}
}

func (bn bodyNode) getUnsafeHTMLString() string      { return bn.unsafeHTMLString }
func (bn bodyNode) getName() string                  { return bn.name }
func (bn bodyNode) getChildrenCount() int            { return len(bn.children) }
func (bn bodyNode) getChildAt(index int) node        { return bn.children[index] }
func (bn bodyNode) getAttributes() map[string]string { return bn.attributes }
func (bn bodyNode) getAttributeOrder() []string      { return bn.attributeOrder }
