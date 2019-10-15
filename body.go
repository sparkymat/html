package html

import "fmt"

type htmlBodyNode struct {
	children []bodyNode
}

// Body returns a <body> node
func Body(children ...bodyNode) htmlBodyNode {
	return htmlBodyNode{children: children}
}

func (b htmlBodyNode) String() string {
	nodeString := ""
	for _, childNode := range b.children {
		nodeString = fmt.Sprintf("%s%s", nodeString, childNode.String())
	}
	return fmt.Sprintf(`<body>
%s</body>
`, nodeString)
}
