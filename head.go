package html

import "fmt"

type HTMLHeadNode struct {
	children []HeadNode
}

// Head returns a <head> node
func Head(children ...HeadNode) HTMLHeadNode {
	return HTMLHeadNode{children: children}
}

func (h HTMLHeadNode) String() string {
	nodeString := ""
	for _, childNode := range h.children {
		nodeString = fmt.Sprintf("%s%s", nodeString, childNode.String())
	}
	return fmt.Sprintf(`<head>
%s</head>
`, nodeString)
}
