package html

type BodyNode struct {
	name             string
	attributes       map[string]string
	attributeOrder   []string
	children         []BodyNode
	unsafeHTMLString string
}

func (bn BodyNode) Attributes(attrs map[string]string) BodyNode {
	copiedNode := bn
	copiedNode.attributes = attrs
	return copiedNode
}

func (bn BodyNode) Children(children ...BodyNode) BodyNode {
	copiedNode := bn
	copiedNode.children = children
	return copiedNode
}

func (bn BodyNode) String() string {
	return nodeAsString(bn)
}

func nodeWithChildren(tag string, children []BodyNode) BodyNode {
	return BodyNode{
		name:     tag,
		children: children,
	}
}

func nodeWithText(tag string, text string) BodyNode {
	return BodyNode{
		name: tag,
		children: []BodyNode{
			Text(text),
		},
	}
}

func (bn BodyNode) getUnsafeHTMLString() string      { return bn.unsafeHTMLString }
func (bn BodyNode) getName() string                  { return bn.name }
func (bn BodyNode) getChildrenCount() int            { return len(bn.children) }
func (bn BodyNode) getChildAt(index int) node        { return bn.children[index] }
func (bn BodyNode) getAttributes() map[string]string { return bn.attributes }
func (bn BodyNode) getAttributeOrder() []string      { return bn.attributeOrder }
