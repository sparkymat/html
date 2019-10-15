package html

import "fmt"

type bodyNode struct {
	children []bodyChildNode
}

// Body returns a <body> node
func Body(children ...bodyChildNode) bodyNode {
	return bodyNode{children: children}
}

func (b bodyNode) String() string {
	nodeString := ""
	for _, childNode := range b.children {
		nodeString = fmt.Sprintf("%s%s", nodeString, childNode.String())
	}
	return fmt.Sprintf(`<body>
%s</body>
`, nodeString)
}
