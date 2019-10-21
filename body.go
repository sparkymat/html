package html

import "fmt"

type HTMLBodyNode struct {
	children []BodyNode
}

// Body returns a <body> node
func Body(children ...BodyNode) HTMLBodyNode {
	return HTMLBodyNode{children: children}
}

func (b HTMLBodyNode) String() string {
	nodeString := ""
	for _, childNode := range b.children {
		nodeString = fmt.Sprintf("%s%s", nodeString, childNode.String())
	}
	return fmt.Sprintf(`<body>
%s</body>
`, nodeString)
}
